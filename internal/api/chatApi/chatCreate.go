package chatApi

import (
	"AI_Server/init/conf"
	"AI_Server/init/data"
	"AI_Server/internal/data/mysql/chat"
	"AI_Server/internal/data/mysql/user"
	"AI_Server/internal/models"
	"AI_Server/utils/res"
	"bufio"
	"bytes"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
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
	createdChat, err := chat.CreateChat(req.Content, req.SessionID, session.Role.ID, session.User.ID)
	if err != nil {
		return res.FailWithMsg(c, "创建对话失败")
	}

	var msgChan chan string
	// 开启数据库事务,扣积分,发送请求
	err = data.DB.Transaction(func(tx *gorm.DB) error {
		err := user.DeductUserPoints(tx, &session.User, conf.GlobalConfig.AI.ChatScope)
		if err != nil {
			return err
		}

		msgChan, err = SendRequest(aiRequest)
		if err != nil {
			return err
		}
		return nil
	})

	c.Context().SetBodyStreamWriter(func(w *bufio.Writer) {
		var aiContent string
		for m := range msgChan {
			fmt.Fprintf(w, "data: Message: %s\n\n", m)
			if err != nil {
				log.Info().Err(err).Msg("发送消息失败")
				return
			}
			aiContent += m
			if err = w.Flush(); err != nil {
				log.Info().Err(err).Msg("Flush 失败")
				return
			}
		}
		if err = chat.UpdateChat(aiContent, createdChat); err != nil {
			return
		}
	})
	return nil
}

func NewOpenAiRequest(req *ChatCreateRequest, aiRole models.AiRole) (*http.Request, error) {
	aiModels := conf.GlobalConfig.AI.Models
	log.Info().Any("req", aiModels[req.AiModel]).Msg("req信息")
	log.Info().Any("AiRole", aiRole).Msg("AiRole信息")
	// 寻找这个会话有关的所有的聊天记录
	chats, err := chat.FindChats(req.SessionID)
	if err != nil {
		return nil, err
	}
	// 测试聊天记录的顺序
	//for _, value := range chats {
	//	log.Info().Str("UserMsg", value.UserContent).Str("AiMsg", value.AiContent).Msg("chat")
	//}

	// 构建OpenAiRequest
	requestBody := &OpenAiRequest{
		Model:    aiModels[req.AiModel].Name,
		Messages: make([]Message, 0, 15),
		Stream:   true,
	}
	// 指定模型角色
	requestBody.Messages = append(requestBody.Messages, Message{
		Role:    "system",
		Content: aiRole.Prompt,
	})
	// 在消息中添加历史消息
	for _, value := range chats {
		requestBody.Messages = append(requestBody.Messages, Message{
			Role:    "user",
			Content: value.UserContent,
		})
		requestBody.Messages = append(requestBody.Messages, Message{
			Role:    "assistant",
			Content: value.AiContent,
		})
	}
	// 在最后附上用户的消息
	requestBody.Messages = append(requestBody.Messages, Message{
		Role:    "user",
		Content: req.Content,
	})

	// 序列化请求体
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
