package routers

import (
	"AI_Server/init/conf"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func InitRouters() {
	app := fiber.New(fiber.Config{
		JSONEncoder:       sonic.Marshal,
		JSONDecoder:       sonic.Unmarshal,
		AppName:           "AiServer",
		EnablePrintRoutes: true,
		Prefork:           false,
	})

	apiRouter := app.Group("/api")

	SettingRouter(apiRouter)
	UserRouter(apiRouter)
	AiRoleRouter(apiRouter)
	SessionRouter(apiRouter)

	log.Fatal().Err(app.Listen(conf.GlobalConfig.Server.Http.Addr))
}
