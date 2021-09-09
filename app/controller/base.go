package controller

import (
	"bala/app/db/mysql"
	"bala/app/log"
	"bala/app/myerr"
	"bala/app/protocol"
	"bala/app/protocol/errcode"
	"bala/app/protocol/protocode"
	"encoding/json"
	"net/http"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	AccountSet,
)

const (
	GameDB      = "gameDb"
	ServerId    = "ServerId"
	Token       = "token"
	Proto_Value = "protocol_value"
	Proto_Str   = "protocol_string"
	Packet      = "packet"
)

func getRequest(ctx *gin.Context) interface{} {
	return ctx.MustGet(Packet)
}

func getGameDB(ctx *gin.Context) *mysql.DB {
	return ctx.MustGet(GameDB).(*mysql.DB)
}

// Response 返回响应消息
func Response(ctx *gin.Context, v interface{}) error {
	// 反射
	value := reflect.ValueOf(v)
	pField := value.Elem().FieldByName("Protocol")
	pField.Set(reflect.ValueOf(ctx.MustGet(Proto_Value)))
	sField := value.Elem().FieldByName("ServerTimeTicks")
	sField.Set(reflect.ValueOf(time.Now().Unix()))
	// 序列化json
	str, err := json.Marshal(v)
	if err != nil {
		log.Local().Error("controller Response error:", err)
		return err
	}
	body := &protocol.ResponseBody{Protocol: ctx.GetString(Proto_Str), Packet: string(str)}
	ctx.JSON(http.StatusOK, body)
	return nil
}

// ResponseError 返回错误消息
func ResponseError(ctx *gin.Context, e error) {
	var respErr *protocol.ResponseError
	if err, ok := e.(*myerr.Error); !ok {
		log.Local().Errorf("server error : %s", e.Error())
		respErr = &protocol.ResponseError{Protocol: errcode.None, Reason: e.Error(), ErrorCode: http.StatusOK}
	} else {
		log.Local().Errorf("server error : %s", err.Stack())
		respErr = &protocol.ResponseError{Protocol: err.Code(), Reason: err.Stack(), ErrorCode: http.StatusOK}
	}
	str, err := json.Marshal(respErr)
	if err != nil {
		ctx.JSON(http.StatusOK, err.Error())
	}
	body := &protocol.ResponseBody{Protocol: protocode.Error.String(), Packet: string(str)}
	ctx.JSON(http.StatusOK, body)
}
