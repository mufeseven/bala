package log

import (
	"bala/app/config"
	"errors"
	"io"
	"os"
	"runtime"
	"time"

	rotate "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

/* 日志轮转相关函数
`WithLinkName` 为最新的日志建立软连接
`WithRotationTime` 设置日志分割的时间，隔多久分割一次
`WithMaxAge 和 WithRotationCount二者只能设置一个
`WithMaxAge` 设置文件清理前的最长保存时间
`WithRotationCount` 设置文件清理前最多保存的个数
*/
func initLocalLogger(path string) (*logrus.Logger, error) {
	var writer *rotate.RotateLogs
	var err error
	switch goos := runtime.GOOS; goos {
	case "windows":
		writer, err = rotate.New(
			path+"%Y-%m-%d.log",
			//rotate.WithLinkName(p),
			rotate.WithMaxAge(-1),
			rotate.WithRotationTime(24*time.Hour),
		)
	case "linux":
		writer, err = rotate.New(
			path+"%Y-%m-%d.log",
			rotate.WithLinkName(path),
			rotate.WithMaxAge(-1),
			rotate.WithRotationTime(24*time.Hour),
		)
	default:
		return nil, errors.New("the os is unknown" + goos)
	}
	if err != nil {
		return nil, err
	}
	logger := new(logrus.Logger)
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05.999999999",
	})
	// 日志消息输出可以是任意的io.writer类型
	if config.GoTest {
		logger.SetOutput(os.Stdout)
	} else {
		logger.SetOutput(io.MultiWriter(os.Stdout, writer))
	}
	return logger, nil
}
