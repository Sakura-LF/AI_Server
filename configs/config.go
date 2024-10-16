package configs

import (
	"time"
)

type Config struct {
	Server   Server      `mapstructure:"server" json:"server" yaml:"server"`
	Data     Data        `mapstructure:"data" json:"data"  yaml:"data"`
	Log      LogConfig   `mapstructure:"log" json:"log" yaml:"log"`
	SiteInfo SiteInfo    `mapstructure:"siteInfo" json:"siteInfo"  yaml:"siteInfo"`
	Email    EmailConfig `mapstructure:"email" json:"email" yaml:"email"`
	Jwt      Jwt         `mapstructure:"jwt" yaml:"jwt" json:"jwt"`
	AI       AIConfig    `mapstructure:"ai" json:"ai" yaml:"ai"`
}

type Data struct {
	DataBase DBConfig
	Redis    RedisConfig
}

type Server struct {
	Http HttpConfig
}

type HttpConfig struct {
	Addr string `mapstructure:"addr" yaml:"addr"`
}

type DBConfig struct {
	Driver           string        `mapstructure:"driver" yaml:"driver"`
	Source           string        `mapstructure:"source" yaml:"source"`
	ReconnectionNum  int           `mapstructure:"reconnection_num" yaml:"reconnection_num"`
	ReconnectionTime time.Duration `mapstructure:"reconnection_time" yaml:"reconnection_time"`
}

type RedisConfig struct {
	Addr         string        `json:"addr"`
	Username     string        `json:"username"`
	Password     string        `json:"password"`
	ReadTimeout  time.Duration `mapstructure:"readTimeout" yaml:"readTimeout" json:"readTimeout"`
	WriteTimeout time.Duration `mapstructure:"writeTimeout" yaml:"writeTimeout" json:"writeTimeout"`
}

// LogConfig 日志配置
type LogConfig struct {
	ZeroLogConfig ZeroLogConfig `yaml:"ZeroLogConfig"`
	LogRotate     LogRotate     `yaml:"LogRotate"`
}

// ZeroLogConfig 日志配置
type ZeroLogConfig struct {
	Level   string `yaml:"Level"`
	Pattern string `yaml:"Pattern"`
	OutPut  string `yaml:"OutPut"`
}

// LogRotate  日志轮换(分割)配置
type LogRotate struct {
	Filename   string `yaml:"Filename"`
	MaxSize    int    `yaml:"MaxSize"`    // megabytes，M 为单位，达到这个设置数后就进行日志切割
	MaxBackups int    `yaml:"MaxBackups"` // 保留旧文件最大份数
	MaxAge     int    `yaml:"MaxAge"`     // days ， 旧文件最大保存天数
	Compress   bool   `yaml:"Compress"`   // 是否开启压缩,默认关闭
}

type SiteInfo struct {
	Site     Site     `json:"site"`
	Project  Project  `json:"project"`
	SEO      SEO      `json:"seo"`
	Register Register `json:"register"`
	Login    Login    `json:"login"`
}

type Site struct {
	Title    string `mapstructure:"title" json:"title" yaml:"title"`
	EnTitle  string `mapstructure:"enTitle" json:"enTitle" yaml:"enTitle"`
	Logo     string `mapstructure:"logo" json:"logo" yaml:"logo"`
	Slogan   string `mapstructure:"slogan" json:"slogan" yaml:"slogan"`
	Abstract string `mapstructure:"abstract" json:"abstract" yaml:"abstract"`
	Beian    string `mapstructure:"beian" json:"beian" yaml:"beian"`
}

type Project struct {
	Title   string `mapstructure:"title" json:"title" yaml:"title"`
	Icon    string `mapstructure:"icon" json:"icon" yaml:"icon"`
	WebPath string `mapstructure:"webPath" json:"webPath" yaml:"webPath"`
}

type SEO struct {
	Keywords    string `mapstructure:"keywords" json:"keywords" yaml:"keywords"`
	Description string `mapstructure:"description" json:"description" yaml:"description"`
}

type Register struct {
	IsEmailRegister   bool `mapstructure:"isEmailRegister" json:"isEmailRegister" yaml:"isEmailRegister"`
	IsWxLoginRegister bool `mapstructure:"isWxLoginRegister" json:"isWxLoginRegister" yaml:"isWxLoginRegister"`
	IsSmsRegister     bool `mapstructure:"isSmsRegister" json:"isSmsRegister" yaml:"isSmsRegister"`
}

type Login struct {
	IsUserNameLogin bool `mapstructure:"isUserNameLogin" json:"isUserNameLogin" yaml:"isUserNameLogin"`
	IsEmailLogin    bool `mapstructure:"isEmailLogin" json:"isEmailLogin" yaml:"isEmailLogin"`
	IsTelLogin      bool `mapstructure:"isTelLogin" json:"isTelLogin" yaml:"isTelLogin"`
}

type EmailConfig struct {
	Domain   string `mapstructure:"domain" json:"domain" yaml:"domain"`
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`
	UserName string `mapstructure:"userName" json:"userName" yaml:"userName"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}

type Jwt struct {
	Secret  string        `yaml:"secret" json:"secret"`
	Expires time.Duration `yaml:"expires" json:"expires"`
	Issuer  string        `yaml:"issuer" json:"issuer"`
}

type AIConfig struct {
	Model              string `mapstructure:"model" json:"model"`
	ProxyURL           string `mapstructure:"proxyUrl" json:"proxyUrl"`
	APIKey             string `mapstructure:"apiKey" json:"apiKey"`
	ChatScope          int    `mapstructure:"chatScope" json:"chatScope"`
	CreateRoleScope    int    `mapstructure:"createRoleScope" json:"createRoleScope"`       // 创建角色的积分消耗
	UpdateRoleScope    int    `mapstructure:"updateRoleScope" json:"updateRoleScope"`       // 更新角色的积分消耗
	DeleteRoleScope    int    `mapstructure:"deleteRoleScope" json:"deleteRoleScope"`       // 删除角色的积分消耗
	RecommendRoleScope int    `mapstructure:"recommendRoleScope" json:"recommendRoleScope"` // 推荐角色成功的积分赠送
	RegisterUserScope  int    `mapstructure:"registerUserScope" json:"registerUserScope"`   // 注册用户成功的积分赠送
}
