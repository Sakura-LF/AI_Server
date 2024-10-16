package aiRole

import (
	"AI_Server/init/data"
	"AI_Server/internal/data/mysql/user"
	"AI_Server/internal/modeles"
	"errors"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

// CreateAiRole 用于创建一个角色
func CreateAiRole(findUser *modeles.User, title, avatar, category, abstract, prompt string) (*modeles.AiRole, error) {
	// 1.查找该用户是否创建了相同的角色
	if _, err := FindAiRole(findUser.ID, title); err == nil {
		//if errors.Is(err, gorm.ErrRecordNotFound) {
		//	// 角色不存在，继续创建角色
		//	log.Info().Msg("角色不存在,创建角色")
		//} else if err != nil {
		//	log.Info().Msg("角色已存在")
		//	return nil, errors.New("角色已存在,同一个用户不能创建相同名称的角色")
		//}
		return nil, errors.New("角色已存在,同一个用户不能创建相同名称的角色")
	}
	// 2.构建角色
	aiRole := &modeles.AiRole{
		UserID:   findUser.ID,
		Title:    title,
		Avatar:   avatar,
		Category: category,
		Abstract: abstract,
		Prompt:   prompt,
	}
	// 3.开启事务创建角色
	err := data.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(aiRole).Error; err != nil {
			return err
		}
		// 4.扣除用户的积分
		if err := user.DeductUserPoints(tx, findUser, 100); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return aiRole, nil
}

// FindAiRole 用于寻找该用户是否创建了相同的角色
func FindAiRole(userID uint, title string) (*modeles.AiRole, error) {
	aiRole := &modeles.AiRole{}
	log.Info().Any("userId", userID).Any("title", title).Msg("查找角色")
	if err := data.DB.Take(aiRole, "user_id = ? and title = ?", userID, title).Error; err != nil {
		return nil, err
	}
	return aiRole, nil
}
