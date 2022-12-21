package rate_limit

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/grin-ch/dever-box-api/pkg/error_enum"
	"golang.org/x/time/rate"
)

var (
	limiter *rate.Limiter
)

func RateLimiter(limit float64, burst int) []gin.HandlerFunc {
	limiter = rate.NewLimiter(rate.Limit(limit), burst)
	return []gin.HandlerFunc{
		rateLimiter,
	}
}

func rateLimiter(ctx *gin.Context) {
	if limiter.Allow() {
		ctx.Next()
		return
	}
	ctx.JSON(http.StatusOK, error_enum.EnumData(error_enum.ServerBusy))
	ctx.Abort()
}
