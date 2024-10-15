package routers

import (
	"AI_Server/internal/api/userApi"
	"AI_Server/internal/middleware"
	"github.com/gofiber/fiber/v3"
)

func UserRouter(r fiber.Router) {
	app := userApi.UserApi{}
	r.Post("/user/login", app.Login)
	r.Get("/user/captcha", app.Captcha)
	r.Get("/user/info", app.UserInfo, middleware.AuthToken())
}
