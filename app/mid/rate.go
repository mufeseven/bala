package mid

import (
	"runtime"

	"github.com/gin-gonic/gin"

	"golang.org/x/time/rate"
)

var (
	limiter = rate.NewLimiter(rate.Limit(300*runtime.NumCPU()), 500*runtime.NumCPU())
)

func RateLimit(c *gin.Context) {
	err := limiter.WaitN(c, 1)
	if err != nil {
		return
	}
}
