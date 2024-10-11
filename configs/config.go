package configs

import (
	"time"
)

type Config struct {
	Server Server
	Data   Data
	Log    LogConfig
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
	ReadTimeout  time.Duration `mapstructure:"read_timeout" json:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout" json:"write_timeout"`
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
