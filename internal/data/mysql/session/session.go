package session

import (
	"AI_Server/init/data"
	"AI_Server/internal/models"
	"github.com/rs/zerolog/log"
)

func CreatSession(title string, userID, roleID uint) (*models.Session, error) {
	session := &models.Session{
		Title:  title,
		UserID: userID,
		RoleID: roleID,
	}
	if err := data.DB.Create(session).Error; err != nil {
		return nil, err
	}
	return session, nil
}

// FinSession 查找session
func FinSession(id uint) (*models.Session, error) {
	var session models.Session
	if err := data.DB.Where("id = ?", id).First(&session).Error; err != nil {
		return nil, err
	}
	return &session, nil
}

func FinSessionByUserID(userID uint, sessionID uint) (*models.Session, error) {
	session := &models.Session{}
	if err := data.DB.Where("id =? and user_id =? ", sessionID, userID).Take(session).Error; err != nil {
		return nil, err
	}
	return session, nil
}

func FinSessionByRoleIDAndUserID(roleID uint, userID uint) (*models.Session, error) {
	var session *models.Session
	if err := data.DB.Order("created_at desc").
		Take(&session, "role_id = ? and user_id = ?", roleID, userID).Error; err != nil {
		return nil, err
	}
	return session, nil
}

func UpdateSessionTitle(session *models.Session, title string) (*models.Session, error) {
	session.Title = title
	if err := data.DB.Save(session).Error; err != nil {
		return nil, err
	}
	return session, nil
}

func FindSessions(idList []uint) (*[]models.Session, error) {
	var sessions []models.Session
	log.Info().Uints("idList", idList).Msg("查找session")
	if err := data.DB.Find(&sessions, "id in ?", idList).Error; err != nil {
		return nil, err
	}
	return &sessions, nil
}

func DeleteSessions(sessionList *[]models.Session) error {
	if err := data.DB.Delete(sessionList).Error; err != nil {
		log.Info().Err(err).Msg("删除session")
		return err
	}
	return nil
}
