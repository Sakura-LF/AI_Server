package middleware

import (
	"AI_Server/common/jwt"
	"AI_Server/internal/modeles"
	"AI_Server/utils/res"
	"github.com/gofiber/fiber/v3"
	"github.com/rs/zerolog/log"
)

func AuthToken() fiber.Handler {
	return func(c fiber.Ctx) error {
		token := c.Get("Authorization")
		log.Info().Msg(token)
		if token == "" {
			return res.FailWithMsg(c, "未携带 Token 请先登录")
		}
		claims, err := jwt.ParseToken(token)
		if err != nil {
			return res.FailWithMsg(c, "Token 解析失败")
		}
		c.Locals("claims", claims.PayLoad)
		log.Info().Msg("Token 认证通过")
		return c.Next()
	}
}

func AuthAdmin() fiber.Handler {
	return func(c fiber.Ctx) error {
		log.Info().Msg("Test")
		token := c.Get("Authorization")
		if token == "" {
			return res.FailWithMsg(c, "未携带 Token 请先登录")
		}
		claims, err := jwt.ParseToken(token)
		if err != nil {
			return res.FailWithMsg(c, "Token 解析失败")
		}
		if claims.PayLoad.Role != modeles.UserRoleNormal {
			return res.FailWithMsg(c, "无权限")
		}

		return c.Next()
	}
}
