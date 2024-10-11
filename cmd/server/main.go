package main

import (
	"AI_Server/init/conf"
	"AI_Server/init/data"
	logz "AI_Server/init/log"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
)

func main() {
	conf.LoadConfig("config")
	log.Info().Interface("全局配置", conf.GlobalConfig).Msg("全局配置")
	// 初始化日志
	logz.InitLog()
	// 初始化数据库
	data.InitDataBase()

	fmt.Println(data.DB)
	log.Info().Err(errors.New("th")).Msg("Sara")

	//for i := 0; i < 100000; i++ {
	//	time.Sleep(time.Second)
	//	log.Info().Msg("Sakura")
	//}

}
