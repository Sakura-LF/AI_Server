package session

import (
	"AI_Server/common/jwt"
	"AI_Server/internal/data/mysql/session"
	"AI_Server/utils/res"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type SessionDeletRequest struct {
	SessionIDs []uint `json:"sessionIDs"`
}

func (*SessionApi) SessionDelete(c *fiber.Ctx) error {
	var req SessionDeletRequest
	err := c.BodyParser(&req)
	if err != nil {
		return res.FailWithMsgAndReason(c, "请求参数错误", err.Error())
	}
	log.Info().Any("Request", req).Msg("请求参数")
	_ = c.Locals("claims").(*jwt.CustomClaims)

	// 1.查找是否存在这些session
	sessions, err := session.FindSessions(req.SessionIDs)
	if err != nil {
		return res.FailWithMsgAndReason(c, "查找 session 失败", err.Error())
	}
	// 2.删除
	err = session.DeleteSessions(sessions)
	if err != nil {
		return res.FailWithMsgAndReason(c, "session未找到或已删除", err.Error())
	}
	return res.OkWithMsgAndData(c, "删除 session 成功", nil)
}
