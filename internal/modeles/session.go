package modeles

import "gorm.io/gorm"

type Session struct {
	gorm.Model
	UserID   uint   `json:"userID"`
	User     User   `gorm:"foreignKey:UserID" json:"-"`
	RoleID   uint   `json:"roleID"`
	Role     AiRole `gorm:"foreignKey:RoleID" json:"-"`
	Title    string `json:"title"` // 会话名称
	ChatList []Chat `gorm:"foreignKey:SessionID" json:"-"`
}
