package controller

import (
	"bala/app/db"
	"bala/app/db/mysql"
	"bala/app/log"

	"github.com/gin-gonic/gin"
)

// 消息处理
type InvokeHandler func(ctx *gin.Context) error

func HandlerInvokeNone(invoke InvokeHandler) gin.HandlerFunc {
	return func(context *gin.Context) {
		if err := invoke(context); err != nil {
			ResponseError(context, err)
		}
	}
}

func HandlerInvokeTrans(invoke InvokeHandler, dbm *db.Manager) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		dbId := ctx.GetInt(ServerId)
		if dbId == 0 {
			log.Local().Error("controller HandlerInvokeTrans dbId == 0")
			return
		}
		gameDB := dbm.GetGameDB(dbId)
		if err := gameDB.Transaction(func(tx *mysql.DB) error {
			ctx.Set(GameDB, gameDB)
			return invoke(ctx)
		}); err != nil {
			ResponseError(ctx, err)
		}
	}
}
