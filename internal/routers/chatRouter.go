package routers

import (
	"AI_Server/internal/api/chatApi"
	"AI_Server/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func ChatRouter(r fiber.Router) {
	app := chatApi.ChatApi{}
	r.Get("/chat/create", middleware.AuthToken(), app.ChatCreate)
}
