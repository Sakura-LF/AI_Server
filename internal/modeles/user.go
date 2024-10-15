package modeles

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nickname       string         `gorm:"size:32" json:"nickname"`    // 昵称
	Username       string         `gorm:"size:128" json:"username"`   // 用户名
	Password       string         `gorm:"size:64" json:"-"`           // 密码
	Tel            string         `gorm:"size:12" json:"tel"`         // 电话
	Email          string         `gorm:"size:128" json:"message"`    // 邮箱
	OpenID         string         `gorm:"size:64" json:"openID"`      // 一般是指第三方登陆的唯一ID
	RegisterSource RegisterSource `json:"registerSource"`             // 0 手机号  1 邮箱  2 微信 3 QQ
	Avatar         string         `gorm:"size:256" json:"avatar"`     // 头像
	Scope          int            `json:"scope"`                      // 积分
	Role           UserRole       `json:"role"`                       // 0 普通用户  1 管理员
	AiRoleList     []AiRole       `gorm:"foreignKey:UserID" json:"-"` // 用户创建的角色列表
	SessionList    []Session      `gorm:"foreignKey:UserID" json:"-"` // 用户创建的会话列表
	ChatList       []Chat         `gorm:"foreignKey:UserID" json:"-"` // 用户聊的对话列表
	OrderList      []Order        `gorm:"foreignKey:UserID" json:"-"` // 用户创建的订单列表
}

type RegisterSource int8

const (
	EmailRegister RegisterSource = iota
	TelRegister
	WxRegister
)

type UserRole int8

const (
	UserRoleNormal UserRole = iota + 1
	UserRoleAdmin
)

// EmailDesensitization 设置对邮箱脱敏
func (u *User) EmailDesensitization() string {
	if u.Email == "" {
		return ""
	}
	// 邮箱脱敏
	u.Email = u.Email[:1] + "*****" + u.Email[len(u.Email)-1:]
	return u.Email
}

// PhoneNumberDesensitization 对手机号进行脱敏处理
func (u *User) PhoneNumberDesensitization() string {
	if u.Tel == "" {
		return ""
	}
	// 手机号脱敏
	return u.Tel[:3] + "****" + u.Tel[7:]
}
