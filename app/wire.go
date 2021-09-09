// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package app

import (
	"bala/app/config"
	"bala/app/controller"
	"bala/app/data"
	"bala/app/db"
	"bala/app/log"
	"bala/app/model/repo"
	"bala/app/router"
	"bala/app/service"

	"github.com/google/wire"

	_ "bala/app/model/entity"
)

// BuildInjector 生成注入器
func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		config.InitConfig,
		log.InitLog,
		InitGin,
		db.InitManager,
		data.WireSet,
		service.WireSet,
		controller.WireSet,
		repo.WireSet,
		router.WireSet,
		InjectorSet,
	)
	return new(Injector), nil, nil
}
