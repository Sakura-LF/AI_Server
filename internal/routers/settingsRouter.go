package routers

import (
	"AI_Server/internal/api/settingsApi"
	"github.com/gofiber/fiber/v3"
)

func SettingRouter(r fiber.Router) {
	app := settingsApi.SettingsApi{}
	r.Get("/settings/:name", app.SettingInfoView)
	r.Put("/settings/:name", app.SettingUpdateView)
}
