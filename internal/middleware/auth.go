package middleware

import (
	"AI_Server/common/jwt"
	"AI_Server/internal/data/redis/user"
	"AI_Server/internal/models"
	"AI_Server/utils/res"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func AuthToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		//log.Info().Msg(token)
		if token == "" {
			return res.FailWithMsg(c, "未携带 Token 请先登录")
		}
		claims, err := jwt.ParseToken(token)
		if err != nil {
			return res.FailWithMsg(c, "Token 解析失败")
		}
		// 判断 token 是否在redis中
		logoutToken, err := user.GetLogoutToken(context.Background(), token)
		if err != nil {
			return res.FailWithMsg(c, err.Error())
		} else if logoutToken {
			return res.FailWithMsg(c, "Token 已过期")
		}

		c.Locals("claims", claims)
		log.Info().Uint("UserID", claims.UserId).Any("Role", claims.Role).Msg("Token 认证通过")
		return c.Next()
	}
}

func AuthAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Info().Msg("Test")
		token := c.Get("Authorization")
		if token == "" {
			return res.FailWithMsg(c, "未携带 Token 请先登录")
		}
		claims, err := jwt.ParseToken(token)
		if err != nil {
			return res.FailWithMsg(c, "Token 解析失败")
		}
		if claims.PayLoad.Role != models.UserRoleNormal {
			return res.FailWithMsg(c, "无权限")
		}

		return c.Next()
	}
}
