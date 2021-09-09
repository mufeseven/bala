package mid

import (
	"bala/app/log"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
)

func RecoveryHandler(c *gin.Context, err interface{}) {
	c.HTML(500, "error.tmpl", gin.H{
		"title": "Error",
		"err":   err,
	})
}

func Recovery(f func(c *gin.Context, err interface{})) gin.HandlerFunc {
	return RecoveryWithWriter(f)
}

func RecoveryWithWriter(f func(c *gin.Context, err interface{})) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				req, _ := httputil.DumpRequest(c.Request, false)
				wrap := errors.Wrap(err, 0)
				log.Local().Errorf("[Recovery] panic recovered:\n %s %s %s", req, wrap.Error(), wrap.Stack())
				f(c, err)
			}
		}()
	}
}
