package routers

import (
	"AI_Server/internal/api/aiRole"
	"AI_Server/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func AiRoleRouter(r fiber.Router) {
	app := aiRole.AiRoleApi{}
	r.Post("/ai/role/create", middleware.AuthToken(), app.RoleCreate)
	r.Get("/ai/role/list", app.AiRoleList)
	r.Post("/ai/role/work_order", middleware.AuthToken(), app.AiRoleCreateWorker)

	r.Get("/ai/role/recent_session", middleware.AuthToken(), app.RecentSession)
}
