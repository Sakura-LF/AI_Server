package routers

import (
	"AI_Server/internal/api/aiRole"
	"AI_Server/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func AiRoleRouter(r fiber.Router) {
	app := aiRole.AiRoleApi{}
	r.Post("/ai/role/create", middleware.AuthToken(), app.RoleCreate)
}
