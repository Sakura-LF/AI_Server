package main

import (
	"AI_Server/init/conf"
	"AI_Server/init/data"
	logz "AI_Server/init/log"
	"AI_Server/internal/routers"
	"fmt"
	"github.com/bytedance/sonic"
)

func main() {
	conf.LoadConfig("config")
	marshalConfig, err := sonic.MarshalIndent(conf.GlobalConfig, "", "\t")
	if err != nil {
		panic(err)
	}
	// 验证配置
	fmt.Println(string(marshalConfig))
	//log.Info().Str("conf", string(marshalConfig)).Msg("全局配置")
	// 初始化日志
	logz.InitLog()
	// 初始化数据库
	data.InitDataBase()
	data.InitRedis()

	routers.InitRouters()
}
