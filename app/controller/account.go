package controller

import (
	"bala/app/data"
	"bala/app/protocol"
	"bala/app/service"

	"github.com/google/wire"

	"github.com/gin-gonic/gin"
)

var AccountSet = wire.NewSet(wire.Struct(new(AccountController), "*"))

type AccountController struct {
	Account  *service.Account
	ItemData *data.ItemData
}

// Create 创建账号
func (a *AccountController) Create(ctx *gin.Context) error {
	req := getRequest(ctx).(*protocol.AccountCreateRequest)
	if ip, b := ctx.RemoteIP(); b {
		req.AccessIP = ip.String()
	}
	err := a.Account.Create(req)
	if err != nil {
		return err
	}
	return Response(ctx, &protocol.AccountCreateResponse{})
}

// Login 登录
func (a *AccountController) Login(ctx *gin.Context) error {
	req := getRequest(ctx).(*protocol.AccountAuthRequest)
	if ip, b := ctx.RemoteIP(); b {
		req.AccessIP = ip.String()
	}
	resp, err := a.Account.Login(req)
	if err != nil {
		return err
	}
	return Response(ctx, resp)
}
