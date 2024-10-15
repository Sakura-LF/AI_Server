package modeles

import "gorm.io/gorm"

// AiRole 角色表
type AiRole struct {
	gorm.Model
	Title       string    `gorm:"size:32" json:"title"`   // 角色名称
	Avatar      string    `gorm:"size:256" json:"avatar"` // 角色头像
	UserID      uint      `json:"userID"`                 // 角色创建人
	UserModel   User      `gorm:"foreignKey:UserID" json:"-"`
	Status      int8      `json:"status"`                   // 角色状态  0 系统创建 1  能够在角色广场看到的   2  热门角色  3 用户创建，用户可见
	Category    string    `gorm:"size:16" json:"category"`  // 角色分类
	Abstract    string    `gorm:"size:128" json:"abstract"` // 角色简介
	Prompt      string    `gorm:"size:1024" json:"prompt"`  // 提示词
	IsSquare    bool      `json:"isSquare"`                 // 是否在角色广场
	IsSystem    bool      `json:"isSystem"`                 // 是否是系统角色
	IsReview    bool      `json:"isReview"`                 // 是否在审核
	SessionList []Session `gorm:"foreignKey:RoleID" json:"-"`
	ChatList    []Chat    `gorm:"foreignKey:RoleID" json:"-"`
}
