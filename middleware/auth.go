package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/EkyGalih/dukcapil-web/config"
	"github.com/EkyGalih/dukcapil-web/utils"
)

// requireAuth ensures a valid (non-expired) token cookie exists
func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie(config.COOKIE_NAME)
		if err != nil || token == "" || utils.JWTIsExpired(token, 15) { // 15s skew
			// not authenticated -> redirect to login
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}
		// put token into context for downstream handlers
		c.Set("access_token", token)
		c.Next()
	}
}