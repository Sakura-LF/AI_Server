package conf

import (
	"AI_Server/configs"
	"AI_Server/utils"
	"github.com/spf13/viper"
	"path/filepath"
)

var GlobalConfig *configs.Config

func LoadConfig(filename string) {
	config := viper.New()

	config.SetConfigName(filename)
	config.AddConfigPath(filepath.Join(utils.GetRootPath(), "configs"))

	err := config.ReadInConfig()
	if err != nil {
		// 如果需要对配置文件不存在错误，做特殊处理，使用：
		panic("配置读取失败" + err.Error())
	}

	//var conf *Config
	err = config.Unmarshal(&GlobalConfig)
	if err != nil {
		panic("配置解析失败" + err.Error())
	}

	return
}
