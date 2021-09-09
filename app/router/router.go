package router

import (
	c "bala/app/controller"
	"bala/app/db"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var WireSet = wire.NewSet(wire.Struct(new(Router), "*"), wire.Bind(new(IRouter), new(*Router)))

// IRouter 注册路由
type IRouter interface {
	Register(e *gin.Engine) error
	Prefixes() []string
}

// Router 路由管理器
type Router struct {
	DBManager *db.Manager
	Account   *c.AccountController
}

// Register 注册路由
func (r *Router) Register(e *gin.Engine) error {
	r.RegisterApi(e)
	return nil
}

// Prefixes 路由前缀列表
func (r *Router) Prefixes() []string {
	return []string{
		"/api/",
	}
}

// 不带自动事务的handler
func (r *Router) Post(ir gin.IRoutes, relativePath string, handler c.InvokeHandler) {
	ir.POST(relativePath, c.HandlerInvokeNone(handler))
}

// 带自动事务的handler
func (r *Router) PostWithTrans(ir gin.IRoutes, relativePath string, handler c.InvokeHandler) {
	ir.POST(relativePath, c.HandlerInvokeTrans(handler, r.DBManager))
}
