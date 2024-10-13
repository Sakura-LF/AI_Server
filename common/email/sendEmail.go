package email

import (
	"AI_Server/init/conf"
	"fmt"
	"github.com/rs/zerolog/log"
	"gopkg.in/gomail.v2"
)

var htmlContent = `
		<!DOCTYPE html>
		<html lang="zh-CN">
		<head>
			<meta charset="UTF-8">
		</head>
		<body style="font-family: Arial, sans-serif; background-color: #f4f4f9;">
			<div style="max-width: 600px; margin: auto; padding: 20px; background: white; border-radius: 8px;">
				<h1 style="color: #333;">欢迎使用 Ai_Server</h1>
				<p>您的验证码为：<strong>%s</strong></p>
				<p>此验证码有效时间为 5 分钟，请妥善保管。</p>
				<footer style="margin-top: 20px; color: #999; text-align: center;">本邮件由系统自动发送，请勿直接回复。</footer>
			</div>
		</body>
		</html>
	`

func NewEmailClient() (gomail.SendCloser, error) {
	// 读取配置
	email := conf.GlobalConfig.Email
	// 构建 client
	dialer := gomail.NewDialer(email.Domain, email.Port, email.UserName, email.Password)
	dial, err := dialer.Dial()

	if err != nil {
		log.Error().Msg("邮件客户端连接失败")
		return nil, err
	}
	return dial, nil
}

func SendEmail(subject string, to string, body string) error {
	email := conf.GlobalConfig.Email
	client, err := NewEmailClient()
	if err != nil {
		return err
	}
	// 创建邮件消息
	message := gomail.NewMessage()
	message.SetHeader("From", email.UserName)
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)

	// 设置邮件正文为 HTML 格式

	// 将验证码插入 HTML 模板
	htmlContent = fmt.Sprintf(htmlContent, body)
	message.SetBody("text/html", htmlContent)

	err = client.Send(email.UserName, []string{to}, message)
	if err != nil {
		return err
	}
	return nil
}

// SendVerifyCode 发送验证码
func SendVerifyCode(to string, code string) error {
	subject := "[AI_Server] 验证码"
	body := code
	return SendEmail(subject, to, body)
}
