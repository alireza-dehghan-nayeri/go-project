package middlewares

import (
	"log"
	"net/http"

	"github.com/alireza-dehghan-nayeri/go-project/util"
	"github.com/gin-gonic/gin"
	limiter "github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := util.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}

func RateLimiterMiddleware() gin.HandlerFunc {
	rate, err := limiter.NewRateFromFormatted("4-M")
	if err != nil {
		log.Fatal(err)
	}
	store := memory.NewStore()
	middleware := mgin.NewMiddleware(limiter.New(store, rate))
	return middleware
}
