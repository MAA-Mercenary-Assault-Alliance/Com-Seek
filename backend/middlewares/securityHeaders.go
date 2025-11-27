package middlewares

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func SecurityHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		frontend := os.Getenv("FRONTEND_ORIGIN")

		c.Writer.Header().Set("X-Frame-Options", "DENY")
		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
		c.Writer.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")

		baseConnect := "'self'"
		if frontend != "" {
			baseConnect = fmt.Sprintf("'self' %s", frontend)
		}

		csp := fmt.Sprintf("default-src 'self'; "+
			"script-src 'self' https://www.google.com https://www.gstatic.com; "+
			"style-src 'self' https://fonts.googleapis.com; "+
			"font-src 'self' https://fonts.gstatic.com; "+
			"img-src 'self' data:; "+
			"connect-src %s; "+
			"frame-ancestors 'none';", baseConnect)

		c.Writer.Header().Set("Content-Security-Policy", csp)

		c.Next()
	}
}
