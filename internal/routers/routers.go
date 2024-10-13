package routers

import (
	"AI_Server/init/conf"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v3"
	"github.com/rs/zerolog/log"
)

func InitRouters() {
	app := fiber.New(fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
		AppName:     "AiServer",
	})

	apiRouter := app.Group("/api")

	SettingRouter(apiRouter)
	UserRouter(apiRouter)

	log.Fatal().Err(app.Listen(conf.GlobalConfig.Server.Http.Addr, fiber.ListenConfig{
		EnablePrefork:     false,
		EnablePrintRoutes: true,
	}))
}
