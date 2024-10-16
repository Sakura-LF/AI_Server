package aiRole

import (
	"AI_Server/common/jwt"
	"AI_Server/init/conf"
	"AI_Server/init/data"
	"AI_Server/internal/data/mysql/aiRole"
	"AI_Server/internal/data/mysql/user"
	"AI_Server/utils/res"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type RoleCreateRequest struct {
	Title    string `json:"title"`
	Avatar   string `json:"avatar"`
	Category string `json:"category"` // 角色分类
	Abstract string `json:"abstract"`
	Prompt   string `json:"prompt"` // 提示词
}

func (role *AiRoleApi) RoleCreate(c *fiber.Ctx) error {
	var req RoleCreateRequest
	err := c.BodyParser(&req)
	if err != nil {
		return res.FailWithMsgAndReason(c, "请求参数错误", err.Error())
	}
	claims := c.Locals("claims").(jwt.PayLoad)

	// 1.验证角色是否存在
	findUser, err := user.FindUserByUserId(claims.UserId)
	if err != nil {
		return res.FailWithMsgAndReason(c, "用户不存在", err.Error())
	}
	// 2.查找该用户是否创建了其他角色
	_, err = aiRole.FindAiRole(findUser.ID, req.Title)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 记录未找到，继续创建角色
			log.Info().Msg("角色不存在,创建角色")
		} else {
			// 其他错误
			return res.FailWithMsgAndReason(c, "查找角色时发送错误", err.Error())
		}
	} else {
		// 角色已存在
		return res.FailWithMsg(c, "角色已存在")
	}

	// 3.开启事务创建角色
	if err = data.DB.Transaction(func(tx *gorm.DB) error {
		_, err = aiRole.CreateAiRole(tx, findUser, req.Title, req.Avatar, req.Category, req.Abstract, req.Prompt)
		if err != nil {
			return err
		}
		// 扣除用户的积分
		err = user.DeductUserPoints(tx, findUser, conf.GlobalConfig.AI.CreateRoleScope)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return res.FailWithMsgAndReason(c, "创建角色失败", err.Error())
	}
	return res.OkWithMsg(c, "创建角色成功")
}
