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

func FindChats(sessionID uint) ([]models.Chat, error) {
	var chats []models.Chat
	if err := data.DB.Order("created_at").Find(&chats, "session_id =?", sessionID).Error; err != nil {
		return nil, err
	}
	return chats, nil
}

func UpdateChat(content string, chat *models.Chat) error {
	if err := data.DB.Model(chat).Update("ai_content", content).Error; err != nil {
		return err
	}
	return nil
}
