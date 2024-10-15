package userApi

import "github.com/gofiber/fiber/v3"

type UserInfoResponse struct {
}

func (userApi *UserApi) UserInfo(c fiber.Ctx) error {
	// 1.校验Token
	return nil
}
