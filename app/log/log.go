package log

import (
	"bala/app/config"

	"github.com/sirupsen/logrus"
)

var (
	log = new(Log)
)

type Log struct {
	local *logrus.Logger
}

func InitLog(config *config.Config) (*Log, error) {
	p := config.Get().Log.Path + "/"
	// 设置日志格式为json格式
	logger, err := initLocalLogger(p)
	if err != nil {
		return nil, err
	}
	log.local = logger
	// 设置日志级别
	level, err := logrus.ParseLevel(config.Get().Log.Level)
	if err != nil {
		return nil, err
	}
	log.local.SetLevel(level)
	return log, nil
}

func (l *Log) Local() *logrus.Logger {
	return l.local
}

func Local() *logrus.Logger {
	return log.local
}
