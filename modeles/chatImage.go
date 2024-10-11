package modeles

import "gorm.io/gorm"

// ChatImage 对话图片关联表
type ChatImage struct {
	gorm.Model
	ChatID  uint  `json:"chatID"`
	Chat    Chat  `gorm:"foreignKey:ChatID" json:"-"`
	ImageID uint  `json:"imageID"`
	Image   Image `gorm:"foreignKey:ImageID" json:"-"`
}
