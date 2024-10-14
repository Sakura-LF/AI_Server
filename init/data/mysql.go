package data

import (
	"AI_Server/init/conf"
	modeles2 "AI_Server/internal/modeles"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	logl "log"
	"os"
	"time"
)

var DB *gorm.DB
var errCount int

func InitDataBase() *gorm.DB {
	// 终端打印输入 sql 执行记录
	newLogger := logger.New(
		logl.New(os.Stdout, "\r\n", logl.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             300 * time.Millisecond, // 慢查询 SQL 阈值
			Colorful:                  true,                   // 是否启动彩色打印
			IgnoreRecordNotFoundError: false,
			LogLevel:                  logger.Error, // Log lever
		},
	)
	db, err := gorm.Open(mysql.Open(conf.GlobalConfig.Data.DataBase.Source), &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		errCount++
		if errCount > conf.GlobalConfig.Data.DataBase.ReconnectionNum {
			panic("failed to connect database:" + err.Error())
		}
		log.Warn().Err(err).Msgf("数据库连接失败,正在进行第%d次重连", errCount)
		time.Sleep(conf.GlobalConfig.Data.DataBase.ReconnectionTime)
		return InitDataBase()
	} else {
		log.Info().Msg("Mysql 连接成功")
	}
	err = db.AutoMigrate(
		&modeles2.User{},
		&modeles2.Session{},
		&modeles2.Chat{},
		&modeles2.AiRole{},
		&modeles2.Image{},
		&modeles2.Order{},
		&modeles2.ChatImage{},
		&modeles2.Log{},
	)
	if err != nil {
		log.Error().Msg("mysql autoMigrate failed")
	} else {
		log.Info().Msg("mysql autoMigrate success")
	}
	DB = db
	return db
}
