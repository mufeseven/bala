module bala

go 1.16

require (
	github.com/gin-gonic/gin v1.7.2
	github.com/go-errors/errors v1.4.0
	github.com/gomodule/redigo v1.8.5
	github.com/google/wire v0.5.0
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/lestrrat-go/strftime v1.0.5 // indirect
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/viper v1.8.1
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0
	gorm.io/driver/mysql v1.1.2
	gorm.io/gorm v1.21.14
)

// replace gorm.io/gorm => ./library/gorm
