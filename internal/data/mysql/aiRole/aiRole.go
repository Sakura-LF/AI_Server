package aiRole

import (
	"AI_Server/init/data"
	"AI_Server/internal/models"
	"gorm.io/gorm"
)

// CreateAiRole 用于创建一个角色
func CreateAiRole(tx *gorm.DB, findUser *models.User, title, avatar, category, abstract, prompt string) (*models.AiRole, error) {
	// 构建角色
	aiRole := &models.AiRole{
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

func FinAiRole(roleID uint) (*models.AiRole, error) {
	aiRole := &models.AiRole{}
	if err := data.DB.Take(aiRole, roleID).Error; err != nil {
		return nil, err
	}
	return aiRole, nil
}

// FindAiRoleByUserID 用于寻找该用户是否创建了相同的角色
func FindAiRoleByUserID(userID uint, title string) (*models.AiRole, error) {
	aiRole := &models.AiRole{}
	if err := data.DB.Take(aiRole, "user_id = ? and title = ?", userID, title).Error; err != nil {
		return nil, err
	}
	return aiRole, nil
}

// FindAiRoleByUserIDAndRoleID  用于寻找该用户是否创建了相同的角色
func FindAiRoleByUserIDAndRoleID(userID uint, roleID uint) (*models.AiRole, error) {
	aiRole := &models.AiRole{}
	if err := data.DB.Take(aiRole, "user_id = ? and id = ?", userID, roleID).Error; err != nil {
		return nil, err
	}
	return aiRole, nil
}

func FinAiRoleIsSystem() (*models.AiRole, error) {
	aiRole := &models.AiRole{}
	if err := data.DB.First(&aiRole, "is_system = ?", true).Error; err != nil {
		return nil, err
	}
	return aiRole, nil
}
