package middleware

import (
	"ginstart/pkg/errno"
	"ginstart/pkg/tools/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// JwtAuth 验证token
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("Authorization")

		parts := strings.SplitN(tokenStr, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, errno.ErrParam)
			c.Abort()
			return
		}

		if parts[1] == "" {
			c.JSON(http.StatusUnauthorized, errno.ErrParam)

			c.Abort()
			return
		}

		// 解析token
		claims, err := jwt.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, errno.ErrSignParam)

			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
