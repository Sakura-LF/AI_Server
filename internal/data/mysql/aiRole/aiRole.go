package aiRole

import (
	"AI_Server/init/data"
	"AI_Server/internal/modeles"
)

func CreateAiRole() {

}

// FindAiRole 用于寻找该用户是否创建了相同的角色
func FindAiRole(userID uint, title string) (*modeles.AiRole, error) {
	aiRole := &modeles.AiRole{}
	if err := data.DB.Take(aiRole, "user_id = ? and title = ?", userID, title).Error; err != nil {
		return nil, err
	}
	return aiRole, nil
}
