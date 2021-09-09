package router

import (
	"github.com/gin-gonic/gin"
)

// RegisterApi register api group router
func (r *Router) RegisterApi(e *gin.Engine) {
	api := e.Group("/api")
	{
		acc := api.Group("/account")
		{
			r.Post(acc, "/create", r.Account.Create)
			r.Post(acc, "/auth", r.Account.Login)
		}
	}
}
