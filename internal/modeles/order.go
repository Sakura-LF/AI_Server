package modeles

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	UserID    uint      `json:"userID"`
	User      User      `gorm:"foreignKey:UserID" json:"-"`
	No        string    `gorm:"size:64" json:"no"` // 订单号
	Price     int64     `json:"price"`             // 单位是分
	Scope     int       `json:"scope"`             // 订单范围
	PayStatus int8      `json:"payStatus"`         // 支付状态 0  1  2  3
	PaySource int8      `json:"paySource"`         // 支付源  0 微信 1 支付宝
	PayedTime time.Time `json:"payedTime"`         // 支付时间
}
