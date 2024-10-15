package aiRole

import (
	"AI_Server/common/jwt"
	"AI_Server/internal/data/mysql/aiRole"
	"AI_Server/utils/res"
	"github.com/gofiber/fiber/v3"
)

type RoleCreateRequest struct {
	Title    string `json:"title"`
	Avatar   string `json:"avatar"`
	Category string `json:"category"` // 角色分类
	Abstract string `json:"abstract"`
	Prompt   string `json:"prompt"` // 提示词
}

func (role *AiRoleApi) RoleCreate(c fiber.Ctx) error {
	var req RoleCreateRequest
	err := c.Bind().Body(&req)
	if err != nil {
		return res.FailWithMsgAndReason(c, "请求参数错误", err.Error())
	}

	claims := c.Locals("claims").(jwt.PayLoad)

	if _, err = aiRole.FindAiRole(claims.UserId, req.Title); err != nil {
		return res.FailWithMsgAndReason(c, "角色已存在,同一个用户不能创建相同名称的角色", err.Error())
	}

	// 开启事务写入

	return nil
}
