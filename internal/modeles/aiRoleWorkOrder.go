package modeles

import "gorm.io/gorm"

// AiRoleWorkOrderModel 角色工单表
type AiRoleWorkOrderModel struct {
	gorm.Model
	Reason      string `gorm:"size:256" json:"reason"` // 申请理由
	UserID      uint   `json:"userID"`
	UserModel   User   `gorm:"foreignKey:UserID" json:"-"`
	AiRoleID    uint   `json:"aiRoleID"`
	AiRoleModel AiRole `gorm:"foreignKey:AiRoleID" json:"-"`
	Type        int8   `json:"type"`   // 1 推荐  2  更新  3 删除
	Status      int8   `json:"status"` // 1  未处理   2  拒绝  3 同意
}
