package userApi

import (
	"AI_Server/common/jwt"
	"AI_Server/internal/data/mysql/user"
	"AI_Server/internal/modeles"
	"AI_Server/utils/res"
	"github.com/gofiber/fiber/v3"
	"time"
)

type UserInfoResponse struct {
	UserID         uint                   `json:"userId"`
	NickName       string                 `json:"nickName"`
	Avatar         string                 `json:"avatar"`
	CreatedAt      time.Time              `json:"createdAt"`
	Tel            string                 `json:"tel"`
	Email          string                 `json:"email"`
	Scope          int                    `json:"score"`
	Role           modeles.UserRole       `json:"role"`
	RegisterSource modeles.RegisterSource `json:"registerSource"`
}

func (userApi *UserApi) UserInfo(c fiber.Ctx) error {
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
		Tel:            findUser.Tel,
		Email:          findUser.Email,
		Scope:          findUser.Scope,
		Role:           findUser.Role,
		RegisterSource: findUser.RegisterSource,
	})
}
