package routers

import (
	"AI_Server/internal/api"
	"github.com/gofiber/fiber/v3"
)

func SettingRouter(r fiber.Router) {
	app := api.SettingsApi{}
	r.Get("/settings", app.SettingInfoView)
}
