package aiRole

import (
	"AI_Server/common/jwt"
	"AI_Server/internal/data/mysql/aiRole"
	"AI_Server/internal/data/mysql/session"
	"AI_Server/utils/res"
	"github.com/gofiber/fiber/v2"
)

type AiRoleRequest struct {
	AiRoleID uint `json:"aiRoleID"`
}

func (role *AiRoleApi) RecentSession(c *fiber.Ctx) error {
	var req AiRoleRequest
	err := c.BodyParser(&req)
	if err != nil {
		return res.FailWithMsgAndReason(c, "请求参数错误", err.Error())
	}
	claims := c.Locals("claims").(jwt.PayLoad)

	// 根据aiRoleID查询角色
	findAiRole, err := aiRole.FinAiRole(req.AiRoleID)
	if err != nil {
		return res.FailWithMsg(c, "角色不存在")
	}
	// 查询角色的会话信息
	findSession, err := session.FinSessionByRoleIDAndUserID(findAiRole.ID, claims.UserId)
	if err != nil {
		return res.OkWithMsgAndData(c, "当前用户无此角色会话", findSession)
	}

	return res.OkWithMsgAndData(c, "会话查询成功", findSession)
}
