package modeles

import "gorm.io/gorm"

type Log struct {
	gorm.Model
	Ip        string `json:"ip"`
	Addr      string `json:"addr"`
	UserID    uint   `json:"userID"`
	UserModel User   `gorm:"foreignKey:UserID" json:"-"`
	Title     string `json:"title"`   // 标题
	Content   string `json:"content"` // 内容
	Level     int8   `json:"level"`   // info warn error
}
