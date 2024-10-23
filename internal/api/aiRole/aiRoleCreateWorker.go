package aiRole

import (
	"AI_Server/common/jwt"
	"AI_Server/internal/data/mysql/aiRole"
	"AI_Server/internal/data/mysql/workOrder"
	"AI_Server/internal/models"
	"AI_Server/utils/res"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

// 用户推荐自己创建的ai角色
// 管理员可以在角色工单列表看到用户推荐的ai角色

type AiRoleWorkCreateRequest struct {
	RoleID           uint           `json:"roleID"`                              // 角色ID
	Reason           string         `json:"reason"`                              // 申请理由
	Type             int8           `json:"type" binding:"required,oneof=1 2 3"` // 1 推荐  2  更新  3 删除
	AiRoleUpdateData *models.AiRole `json:"aiRoleUpdateData"`                    // 角色更新的信息  type=2必填
}

// AiRoleCreateWorker 创建工单,包含推荐角色更新角色和删除角色
func (role *AiRoleApi) AiRoleCreateWorker(c *fiber.Ctx) error {
	// 绑定参数
	var req AiRoleWorkCreateRequest
	//c.BodyParser()
	err := c.BodyParser(&req)
	if err != nil {
		return res.FailWithMsgAndReason(c, "请求参数错误", err.Error())
	}
	log.Info().Any("Request", req).Msg("请求信息")
	claims := c.Locals("claims")
	userClaims, _ := claims.(*jwt.CustomClaims)

	//var workOrder models.AiRoleWorkOrderModel
	//switch req.Type {
	//case 1:
	//case 2:
	//	// 更新角色的工单
	//case 3:
	//
	//}
	// 查找用户有没有创建过这个角色
	findRole, err := aiRole.FindAiRoleByUserIDAndRoleID(userClaims.UserId, req.RoleID)
	if err != nil {
		return res.FailWithMsg(c, "角色不存在")
	}
	log.Info().Any("findRole", findRole).Msg("Role")

	// 查找用户有没有创建过这个工单
	_, err = workOrder.FindWorkOrder(req.RoleID, req.Type)
	if err == nil {
		return res.FailWithMsg(c, "请勿重复提交工单")
	}

	// 创建工单
	_, err = workOrder.CreateWorkOrder(userClaims.UserId, req.RoleID, req.Reason, req.Type, []byte("test"))
	if err != nil {
		return err
	}
	return res.OkWithMsg(c, "提交工单成功")
}
