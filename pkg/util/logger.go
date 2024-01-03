package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func Init() {
	logFile, err := os.OpenFile("logs/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("打开日志文件失败")
	}

	log.Out = logFile
	log.Formatter = &logrus.TextFormatter{}
	log.Level = logrus.InfoLevel
}

func GetLogger() *logrus.Logger {
	return log
}
