package user

import (
	"AI_Server/init/data"
	"AI_Server/internal/modeles"
	"AI_Server/utils/rand"
	"errors"
	"github.com/rs/zerolog/log"
)

func CreateUser(registerSource modeles.RegisterSource, val string) error {
	randomString, err := rand.GetRandomUserName()
	if err != nil {
		return err
	}
	user := &modeles.User{
		Username:       randomString,
		Nickname:       rand.GetRandomNickName(randomString),
		RegisterSource: registerSource,
	}

	log.Info().Any("user", user).Msg("创建用户")

	switch registerSource {
	case modeles.EmailRegister:
		user.Email = val
		_, err := FindUserByEmail(val)
		if err == nil {
			return errors.New("邮箱已存在")
		}
	default:
		return errors.New("不支持的注册方式")
	}

	if err = data.DB.Create(user).Error; err != nil {
		log.Error().Err(err).Msg("创建用户失败")
		return errors.New("创建用户失败,请联系网站管理员")
	}
	return nil
}

func FindUserByEmail(email string) (modeles.User, error) {
	var user modeles.User
	if err := data.DB.Where("email = ?", email).Take(&user).Error; err != nil {
		//log.Error().Err(err).Msg("查询用户失败")
		return user, err
	}
	return user, nil

}
