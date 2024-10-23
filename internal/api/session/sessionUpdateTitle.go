package session

import (
	"AI_Server/common/jwt"
	"AI_Server/internal/data/mysql/session"
	"AI_Server/utils/res"
	"github.com/gofiber/fiber/v2"
)

type SessionUpdateTitleRequest struct {
	SessionID uint   `json:"sessionID"`
	Title     string `json:"title"`
}

// SessionUpdate 修改session title
func (*SessionApi) SessionUpdate(c *fiber.Ctx) error {
	var req SessionUpdateTitleRequest
	err := c.BodyParser(&req)
	if err != nil {
		return res.FailWithMsgAndReason(c, "请求参数错误", err.Error())
	}
	claims := c.Locals("claims")
	userClaims, _ := claims.(*jwt.CustomClaims)

	findSession, err := session.FinSessionByUserID(userClaims.UserId, req.SessionID)
	if err != nil {
		return res.FailWithMsgAndReason(c, "会话不存在", err.Error())
	}

	// 更新title
	title, err := session.UpdateSessionTitle(findSession, req.Title)
	if err != nil {
		return res.FailWithMsgAndReason(c, "更新会话失败", err.Error())
	}
	return res.OkWithMsgAndData(c, "更新会话成功", title)
}
