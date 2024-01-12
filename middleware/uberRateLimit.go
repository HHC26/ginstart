package middleware

import (
	"ginstart/global"

	"github.com/gin-gonic/gin"
	"go.uber.org/ratelimit"
)

func UberRateLimit() gin.HandlerFunc {
	number := global.Conf.System.UberlimitCount

	var Rl = ratelimit.New(number)
	return func(c *gin.Context) {
		Rl.Take()
	}
}
