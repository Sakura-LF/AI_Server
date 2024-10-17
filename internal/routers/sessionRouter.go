package routers

import (
	"AI_Server/internal/api/session"
	"AI_Server/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func SessionRouter(r fiber.Router) {
	app := session.SessionApi{}
	r.Post("/session/create", middleware.AuthToken(), app.SessionCreate)
	r.Patch("/session/update", middleware.AuthToken(), app.SessionUpdate)
	r.Delete("/session/delete", middleware.AuthToken(), app.SessionDelete)
}
