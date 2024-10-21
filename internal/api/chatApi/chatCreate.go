package chatApi

import (
	"AI_Server/init/conf"
	"AI_Server/init/data"
	"AI_Server/internal/data/mysql/chat"
	"AI_Server/internal/models"
	"AI_Server/utils/res"
	"bufio"
	"bytes"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"net/http"
)

type ChatCreateRequest struct {
	SessionID   uint    `json:"sessionID" form:"sessionID" binding:"required"`
	Content     string  `json:"content" form:"content" binding:"required"`
	ImageIDList []uint  `json:"imageIDList" form:"imageIDList"`
	AiModel     AiModel `json:"aiModel" form:"aiModel"`
}

type OpenAiRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream"`
}
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAiResponse struct {
	Choices []struct {
		Delta struct {
			Content string `json:"content"`
		} `json:"delta"`
		FinishReason interface{} `json:"finish_reason"`
		Index        int         `json:"index"`
		Logprobs     interface{} `json:"logprobs"`
	} `json:"choices"`
	Object            string      `json:"object"`
	Usage             interface{} `json:"usage"`
	Created           int         `json:"created"`
	SystemFingerprint interface{} `json:"system_fingerprint"`
	Model             string      `json:"model"`
	Id                string      `json:"id"`
}

func (*ChatApi) ChatCreate(c *fiber.Ctx) error {
	// 设置响应头
	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")
	c.Set("Transfer-Encoding", "chunked")

	var req *ChatCreateRequest
	err := c.BodyParser(&req)
	if err != nil {
		return res.FailWithMsgAndReason(c, "请求参数错误", err.Error())
	}
	log.Info().Any("req", req).Msg("请求信息")

	// 1. 查找session存不存在
	var session *models.Session
	err = data.DB.Preload("User").Preload("Role").
		Take(&session, req.SessionID).Error
	if err != nil {
		return res.FailWithMsg(c, "会话不存在")
	}
	if session.Role.ID == 0 {
		return res.FailWithMsg(c, "角色被删除")
	}
	// 构建请求
	aiRequest, err := NewOpenAiRequest(req, session.Role)
	if err != nil {
		return res.FailWithMsg(c, "创建请求失败")
	}
	// 创建对话
	_, err = chat.CreateChat(req.Content, req.SessionID, session.Role.ID, session.User.ID)
	if err != nil {
		return res.FailWithMsg(c, "创建对话失败")
	}

	var msgChan chan string
	msgChan, err = SendRequest(aiRequest)
	if err != nil {
		return err
	}
	c.Context().SetBodyStreamWriter(func(w *bufio.Writer) {
		for m := range msgChan {
			fmt.Fprintf(w, "data: Message: %s\n\n", m)
			if err != nil {
				log.Info().Err(err).Msg("发送消息失败")
				return
			}
			if err = w.Flush(); err != nil {
				log.Info().Err(err).Msg("Flush 失败")
				return
			}
		}
	})
	return nil
}

func NewOpenAiRequest(req *ChatCreateRequest, aiRole models.AiRole) (*http.Request, error) {
	aiModels := conf.GlobalConfig.AI.Models
	log.Info().Any("req", aiModels[req.AiModel]).Msg("req信息")
	log.Info().Any("AiRole", aiRole).Msg("AiRole信息")
	//clients := &http.Client{}
	requestBody := &OpenAiRequest{
		Model: aiModels[req.AiModel].Name,
		Messages: []Message{
			{
				Role:    "system",
				Content: aiRole.Prompt,
			},
			{
				Role:    "user",
				Content: req.Content,
			},
		},
		Stream: true,
	}
	marshal, err2 := sonic.Marshal(requestBody)
	if err2 != nil {
		panic(err2)
	}

	request, err := http.NewRequest("POST", aiModels[req.AiModel].ProxyURL, bytes.NewReader(marshal))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", "Bearer "+aiModels[req.AiModel].APIKey)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-DashScope-SSE", "enable")

	return request, nil
}

// SendRequest 发送请求
func SendRequest(req *http.Request) (msgChan chan string, err error) {
	client := &http.Client{}
	msgChan = make(chan string)

	resp, err := client.Do(req)
	if err != nil {
		log.Info().Err(err).Msg("调用大模型失败")
		return nil, err
	}
	log.Info().Msg("调用大模型成功")
	go func() {
		defer resp.Body.Close()
		scanner := bufio.NewScanner(resp.Body)
		scanner.Split(bufio.ScanLines) // 按照行读

		for scanner.Scan() {
			msg := scanner.Text()
			if msg == "" {
				continue
			}
			if msg == "data: [DONE]" {
				// 读完了
				break
			}
			var m OpenAiResponse
			err = sonic.Unmarshal([]byte(msg[5:]), &m)
			if err != nil {
				fmt.Println(msg, err)
				continue
			}
			msgChan <- m.Choices[0].Delta.Content
		}
		close(msgChan)
	}()
	return msgChan, nil
}
