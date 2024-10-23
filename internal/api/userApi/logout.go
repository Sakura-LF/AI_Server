package userApi

import (
	"AI_Server/common/jwt"
	"AI_Server/internal/data/redis/user"
	"AI_Server/utils/res"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"time"
)

func (userApi *UserApi) Logout(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	claims, err := jwt.ParseToken(token)
	if err != nil {
		return res.FailWithMsg(c, "Token 解析失败")
	}

	log.Info().Time("exp", claims.ExpiresAt.Time).Msg("过期时间")

	//log.Info().Any("now", duration).Msg("过期时间")

	if err := user.SetLogoutToken(context.Background(), token, claims.ExpiresAt.Time.Sub(time.Now())); err != nil {
		return res.OkWithMsgAndData(c, "注销失败", err.Error())
	}

	// 获取 Token
	return res.OkWithMsg(c, "注销成功")
}
