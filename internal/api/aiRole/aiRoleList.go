package aiRole

import (
	"AI_Server/common/jwt"
	"AI_Server/internal/models"
	"AI_Server/utils/res"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type AiRoleListRequest struct {
	//models.Page
	Type int8 `form:"type" json:"type" binding:"required,oneof=1 2 3 4"`
	//1. 角色广场
	//2. 热门角色
	//3. 我创建的角色列表
	//4. 后台角色列表
}

type AiRoleListResponse struct {
	models.AiRole
	SessionCount     int    `json:"sessionCount"`
	ChatCount        int    `json:"chatCount"`
	RoleUserNickname string `json:"roleUserNickname,omitempty"` // 创建角色的人
}

func (role *AiRoleApi) AiRoleList(c *fiber.Ctx) error {
	var req AiRoleListRequest
	err := c.BodyParser(&req)
	if err != nil {
		return res.FailWithMsgAndReason(c, "请求参数错误", err.Error())
	}
	// 解析token
	token := c.Get("Authorization")
	claims, err := jwt.ParseToken(token)
	if err != nil {
		claims = &jwt.CustomClaims{}
		log.Info().Any("token", claims).Msg("用户未登录")
		//return res.FailWithMsgAndReason(c, "token解析错误", err.Error())
	}
	switch claims.Role {
	case models.UserRoleUnLogin:
		switch req.Type {
		case 3, 4:
			return res.FailWithMsg(c, "无权限")
		}
	case models.UserRoleNormal:
		switch req.Type {
		case 4:
			return res.FailWithMsg(c, "无权限")
		}
	case models.UserRoleAdmin:
	default:
		return res.FailWithMsg(c, "无权限")
	}

	// todo 分页查询,之后做

	return res.OkWithMsg(c, "ok")
}
