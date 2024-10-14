package message

import (
	"AI_Server/internal/modeles"
	"AI_Server/utils/rand"
	"errors"
)

var (
	ErrInvalidMessageType = errors.New("invalid message type")
)

func SendCode(messageType modeles.RegisterSource, to string) error {
	// 根据类型选择不同的发送方式
	randomCode := rand.GetRandomCode(6)
	switch messageType {
	case modeles.EmailRegister:
		return SendVerifyCode(to, randomCode)
	case modeles.TelRegister:
		panic("not implemented")
		//return SendVerifyCode(to, randomCode)
	default:
		return ErrInvalidMessageType
	}
}
