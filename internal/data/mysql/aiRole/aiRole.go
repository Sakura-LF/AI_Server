package aiRole

import (
	"AI_Server/init/data"
	"AI_Server/internal/modeles"
	"gorm.io/gorm"
)

// CreateAiRole 用于创建一个角色
func CreateAiRole(tx *gorm.DB, findUser *modeles.User, title, avatar, category, abstract, prompt string) (*modeles.AiRole, error) {
	// 构建角色
	aiRole := &modeles.AiRole{
		UserID:   findUser.ID,
		Title:    title,
		Avatar:   avatar,
		Category: category,
		Abstract: abstract,
		Prompt:   prompt,
	}
	// 事务创建角色
	if err := tx.Create(aiRole).Error; err != nil {
		return nil, err
	}
	return aiRole, nil
}

// FindAiRole 用于寻找该用户是否创建了相同的角色
func FindAiRole(userID uint, title string) (*modeles.AiRole, error) {
	aiRole := &modeles.AiRole{}
	if err := data.DB.Take(aiRole, "user_id = ? and title = ?", userID, title).Error; err != nil {
		return nil, err
	}
	return aiRole, nil
}
