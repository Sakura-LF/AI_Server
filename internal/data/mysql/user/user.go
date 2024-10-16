package user

import (
	"AI_Server/init/data"
	"AI_Server/internal/modeles"
	"AI_Server/utils/rand"
	"errors"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func CreateUser(registerSource modeles.RegisterSource, val string) (*modeles.User, error) {
	randomString, err := rand.GetRandomUserName()
	if err != nil {
		return nil, nil
	}
	user := &modeles.User{
		Username:       randomString,
		Nickname:       rand.GetRandomNickName(randomString),
		RegisterSource: registerSource,
		Role:           modeles.UserRoleNormal,
	}
	switch registerSource {
	case modeles.EmailRegister:
		user.Email = val
	case modeles.TelRegister:
		user.Tel = val
	default:
		return nil, errors.New("不支持的注册方式")
	}
	//log.Info().Any("user", user).Msg("创建用户")
	if err = data.DB.Create(user).Error; err != nil {
		log.Error().Err(err).Msg("创建用户失败")
		return nil, err
	}
	return user, nil
}

func FindUserByEmailOrTel(registerSource modeles.RegisterSource, val string) (*modeles.User, error) {
	user := &modeles.User{}
	switch registerSource {
	case modeles.EmailRegister:
		if err := data.DB.Where("email = ?", val).Take(user).Error; err != nil {
			return user, err
		}
	case modeles.TelRegister:
		if err := data.DB.Where("tel = ?", val).Take(user).Error; err != nil {
			return user, err
		}
	default:
		return user, errors.New("不支持的注册方式")
	}
	return user, nil
}

func FindUserByUserId(id uint) (*modeles.User, error) {
	user := &modeles.User{}
	if err := data.DB.Take(user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// DeductUserPoints 扣除用户积分
// 调用这个方法前需要开启数据库事务
func DeductUserPoints(tx *gorm.DB, user *modeles.User, point int) error {
	newPoints := user.Scope - point
	if newPoints < 0 {
		return errors.New("积分不足")
	}
	// 开启数据库事务
	if err := tx.Model(user).Update("scope", newPoints).Error; err != nil {
		return err
	}
	return nil
}

// AddUserPoints 增加用户积分
// 调用这个方法前需要开启数据库事务
func AddUserPoints(tx *gorm.DB, user *modeles.User, point int) error {
	newPoints := user.Scope + point
	// 开启数据库事务
	if err := tx.Model(user).Update("scope", newPoints).Error; err != nil {
		return err
	}
	return nil
}
