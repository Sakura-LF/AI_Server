package userApi

import (
	"AI_Server/common/jwt"
	"AI_Server/internal/data/mysql/user"
	"AI_Server/internal/models"
	"AI_Server/utils/res"
	"github.com/gofiber/fiber/v2"
	"time"
)

type UserInfoResponse struct {
	UserID         uint                  `json:"userId"`
	NickName       string                `json:"nickName"`
	Avatar         string                `json:"avatar"`
	CreatedAt      time.Time             `json:"createdAt"`
	Tel            string                `json:"tel"`
	Email          string                `json:"email"`
	Scope          int                   `json:"score"`
	Role           models.UserRole       `json:"role"`
	RegisterSource models.RegisterSource `json:"registerSource"`
}

func (userApi *UserApi) UserInfo(c *fiber.Ctx) error {
	claims := c.Locals("claims").(jwt.PayLoad)

	findUser, err := user.FindUserByUserId(claims.UserId)
	if err != nil {
		return res.FailWithMsgAndReason(c, "用户不存在", err.Error())
	}

	return res.OkWithData(c, UserInfoResponse{
		UserID:         findUser.ID,
		Avatar:         findUser.Avatar,
		NickName:       findUser.Nickname,
		CreatedAt:      findUser.CreatedAt,
		Tel:            findUser.PhoneNumberDesensitization(),
		Email:          findUser.Email,
		Scope:          findUser.Scope,
		Role:           findUser.Role,
		RegisterSource: findUser.RegisterSource,
	})
}
