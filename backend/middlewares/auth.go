package middlewares

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	ID uint `json:"id"`
	jwt.RegisteredClaims
}

func CheckAuth(c *gin.Context) {
	tokenString, err := c.Cookie("token")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token not found in cookie"})
		return
	}

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, http.ErrNotSupported
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		return
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
		return
	}

	c.Set("userID", claims.ID)

	c.Next()
}
