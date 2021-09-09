package app

import (
	"bala/app/config"
	"bala/app/db"
	"bala/app/log"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// InjectorSet 注入Injector
var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

// Injector 注入器(用于初始化完成之后的引用)
type Injector struct {
	Config *config.Config
	Log    *log.Log
	Engine *gin.Engine
	Shards *db.Manager
}
