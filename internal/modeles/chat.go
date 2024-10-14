package modeles

import "gorm.io/gorm"

type Chat struct {
	gorm.Model
	SessionID   uint        `json:"sessionID"`
	Session     Session     `gorm:"foreignKey:SessionID" json:"-"`
	RoleID      uint        `json:"roleID"`
	Role        AiRole      `gorm:"foreignKey:RoleID" json:"-"`
	UserID      uint        `json:"userID"`
	User        User        `gorm:"foreignKey:UserID" json:"-"`
	UserContent string      `json:"userContent"`                // 用户的聊天内容
	AiContent   string      `json:"aiContent"`                  // 机器人的聊天内容
	ImageList   []ChatImage `gorm:"foreignKey:ChatID" json:"-"` // 对话关联的图片列表
}
