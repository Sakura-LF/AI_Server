package routers

import (
	"AI_Server/internal/api/aiRole"
	"github.com/gofiber/fiber/v3"
)

func AiRoleRouter(r fiber.Router) {
	app := aiRole.AiRoleApi{}
	r.Post("/ai/role/create", app.RoleCreate)
}
