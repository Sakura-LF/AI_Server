package routers

import (
	"AI_Server/internal/api/userApi"
	"github.com/gofiber/fiber/v3"
)

func UserRouter(r fiber.Router) {
	app := userApi.UserApi{}
	r.Post("/user/register", app.Register)
	r.Get("/user/captcha", app.Captcha)
}
