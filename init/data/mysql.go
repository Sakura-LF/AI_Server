package data

import (
	"AI_Server/init/conf"
	"AI_Server/modeles"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	logl "log"
	"os"
	"time"
)

var DB *gorm.DB

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
		panic("failed to connect database")
	} else {
		log.Info().Msg("mysql connect success")
	}
	err = db.AutoMigrate(
		&modeles.User{},
		&modeles.Session{},
		&modeles.Chat{},
		&modeles.AiRole{},
		&modeles.Image{},
		&modeles.Order{},
		&modeles.ChatImage{},
		&modeles.Log{},
	)
	if err != nil {
		log.Error().Msg("mysql autoMigrate failed")
	} else {
		log.Info().Msg("mysql autoMigrate success")
	}
	DB = db
	return db
}
