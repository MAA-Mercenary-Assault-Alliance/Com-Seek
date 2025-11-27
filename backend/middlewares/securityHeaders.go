package middlewares

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func SecurityHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		frontend := os.Getenv("FRONTEND_ORIGIN")

		c.Writer.Header().Set("X-Frame-Options", "DENY")
		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
		c.Writer.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")

		csp := "default-src 'none'; frame-ancestors 'none';"

		if frontend != "" {
			csp = fmt.Sprintf("default-src 'none'; connect-src 'self' %s; frame-ancestors 'none';", frontend)
		}

		c.Writer.Header().Set("Content-Security-Policy", csp)

		c.Next()

		contentType := c.Writer.Header().Get("Content-Type")

		if strings.Contains(contentType, "image/") || strings.Contains(contentType, "application/pdf") {
			fileCsp := "default-src 'none'; img-src 'self' data:; object-src 'self';"

			c.Writer.Header().Set("Content-Security-Policy", fileCsp)
		}
	}
}
