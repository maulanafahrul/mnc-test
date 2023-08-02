package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/maulanafahrul/mnc-test/utils"
)

type authHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

func RequireToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// check exist token
		h := &authHeader{}
		if err := c.ShouldBindHeader(&h); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": "Unauthorize",
			})
			c.Abort()
			return
		}

		tokenString := strings.Replace(h.AuthorizationHeader, "Bearer ", "", -1)

		// check token kosong
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": "Unauthorize",
			})
			c.Abort()
			return
		}

		// check verify token
		token, err := utils.VerifyAccessToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": "Unauthorize",
			})
			c.Abort()
			return
		}

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": "Unauthorize",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
