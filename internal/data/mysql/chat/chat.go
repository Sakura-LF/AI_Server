package chat

import (
	"AI_Server/init/data"
	"AI_Server/internal/models"
)

func CreateChat(content string, sessionID, roleID, userID uint) (*models.Chat, error) {
	chat := &models.Chat{
		UserContent: content,
		SessionID:   sessionID,
		RoleID:      roleID,
		UserID:      userID,
	}
	if err := data.DB.Create(chat).Error; err != nil {
		return nil, err
	}
	return chat, nil
}
