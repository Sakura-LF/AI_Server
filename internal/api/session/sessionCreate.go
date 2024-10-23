package session

import (
	"AI_Server/common/jwt"
	"AI_Server/internal/data/mysql/aiRole"
	"AI_Server/internal/data/mysql/session"
	"AI_Server/utils/res"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"time"
)

type CreateSessionRequest struct {
	// 角色ID
	RoleId uint `json:"roleId"`
	// 会话名称
	SessionName string `json:"sessionName"`
}

// SessionCreate 基于用户 ID 创建会话
func (*SessionApi) SessionCreate(c *fiber.Ctx) error {
	var req CreateSessionRequest
	err := c.BodyParser(&req)
	if err != nil {
		return res.FailWithMsgAndReason(c, "请求参数错误", err.Error())
	}
	claims := c.Locals("claims")
	userClaims, _ := claims.(*jwt.CustomClaims)

	if req.SessionName == "" {
		// 根据时间戳生成会话名称
		format := time.Now().Format("2006-01-02 15:04:05")
		req.SessionName = "会话" + format
	}
	log.Info().Any("req", req).Msg("请求信息")

	// 如果没有选择角色ID
	if req.RoleId == 0 {
		findRole, err := aiRole.FinAiRoleIsSystem()
		if err != nil {
			return res.FailWithMsgAndReason(c, "无系统角色,请自行创建角色", err.Error())
		}
		log.Info().Any("aiRole", findRole).Msg("系统角色")
		req.RoleId = findRole.ID
	} else {
		// 查询角色是否存在
		findRole, err := aiRole.FinAiRole(req.RoleId)
		if err != nil {
			return res.FailWithMsgAndReason(c, "角色不存在", err.Error())
		}
		// 查询该用户是否在角色广场
		if !findRole.IsSquare {
			// 不在角色广场就要查是否是用户自己的角色
			if userClaims.UserId != findRole.UserID {
				return res.FailWithMsg(c, "当前用户无角色")
			}
		}
	}

	// 创建会话
	createSession, err := session.CreatSession(req.SessionName, userClaims.UserId, req.RoleId)
	if err != nil {
		return res.FailWithMsgAndReason(c, "会话创建失败", err.Error())
	}
	return res.OkWithMsgAndData(c, "会话创建成功", createSession.ID)
}
