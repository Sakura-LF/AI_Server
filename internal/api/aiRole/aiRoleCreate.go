package aiRole

import (
	"AI_Server/common/jwt"
	"AI_Server/internal/data/mysql/aiRole"
	"AI_Server/internal/data/mysql/user"
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

	// 验证角色是否存在
	findUser, err := user.FindUserByUserId(claims.UserId)
	if err != nil {
		return res.FailWithMsgAndReason(c, "用户不存在", err.Error())
	}
	if _, err = aiRole.CreateAiRole(findUser, req.Title, req.Avatar, req.Category, req.Abstract, req.Prompt); err != nil {
		return res.FailWithMsgAndReason(c, "创建角色失败", err.Error())
	}

	return res.OkWithMsg(c, "创建角色成功")
}
