package main

import (
	"AI_Server/init/conf"
	"AI_Server/init/data"
	logz "AI_Server/init/log"
	"fmt"
	"github.com/rs/zerolog/log"
	"time"
)

func main() {
	conf.LoadConfig("config")
	fmt.Println(conf.GlobalConfig)
	// 初始化日志
	logz.InitLog()
	// 初始化数据库
	data.InitDataBase()

	fmt.Println(data.DB)

	for i := 0; i < 100000; i++ {
		time.Sleep(time.Second)
		log.Info().Msg("Sakura")
	}

}
