package api

import "github.com/gofiber/fiber/v3"

type SettingsApi struct {
}

func (receiver SettingsApi) SettingInfoView(c fiber.Ctx) error {
	c.WriteString("hello world")
	return nil
}
