package routers

import (
	"AI_Server/internal/api/chatApi"
	"github.com/gofiber/fiber/v3"
)

func ChatRouterRouter(r fiber.Router) {
	app := chatApi.ChatApi{}
	r.Get("/chat/create/", app.ChatCreate)
}
