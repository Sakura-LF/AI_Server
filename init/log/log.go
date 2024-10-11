package log

import (
	"AI_Server/init/conf"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// InitLog  完成Zero 日志的初始化
func InitLog() {
	// 设置日志等级
	switch conf.GlobalConfig.Log.ZeroLogConfig.Level {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	}

	// 日志切割
	logRotate := &lumberjack.Logger{
		Filename:   strings.Join([]string{conf.GlobalConfig.Log.ZeroLogConfig.OutPut, conf.GlobalConfig.Log.LogRotate.Filename}, "/"), // 文件位置
		MaxSize:    100,                                                                                                               // megabytes，M 为单位，达到这个设置数后就进行日志切割
		LocalTime:  true,
		MaxBackups: 3,    // 保留旧文件最大份数
		MaxAge:     7,    //days ， 旧文件最大保存天数
		Compress:   true, // disabled by default，是否压缩日志归档，默认不压缩
	}
	// 调整日志时间格式
	zerolog.TimeFieldFormat = time.StampMilli
	// 开启调用位置打印
	log.Logger = log.With().Caller().Logger()

	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		var builder strings.Builder
		builder.WriteString(filepath.Base(file))
		builder.WriteString(":")
		builder.WriteString(strconv.Itoa(line))
		return builder.String()
	}

	if conf.GlobalConfig.Log.ZeroLogConfig.Pattern == "development" {
		// 控制台输出的输出器
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.StampMilli}
		multi := zerolog.MultiLevelWriter(consoleWriter, logRotate)
		log.Logger = log.Output(multi)
	} else if conf.GlobalConfig.Log.ZeroLogConfig.Pattern == "production" {
		log.Logger = log.Output(logRotate)
	} else {
		panic("log pattern Error")
	}
}
