package modeles

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	UserID    uint   `json:"userID"`
	UserModel User   `gorm:"foreignKey:UserID" json:"-"`
	Filename  string `gorm:"size:64" json:"filename"`  // 文件名
	FilePath  string `gorm:"size:256" json:"filePath"` // 文件路径
	Size      int64  `json:"size"`                     // 单位是字节
}
