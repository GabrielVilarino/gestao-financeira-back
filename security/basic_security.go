package security

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func BasicAuth(expectedUser, expectedPassword string) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			unauthorized(c)
			return
		}

		// Esperado: "Basic base64(user:pass)"
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Basic" {
			unauthorized(c)
			return
		}

		decoded, err := base64.StdEncoding.DecodeString(parts[1])
		if err != nil {
			unauthorized(c)
			return
		}

		credentials := strings.SplitN(string(decoded), ":", 2)
		if len(credentials) != 2 {
			unauthorized(c)
			return
		}

		username := credentials[0]
		password := credentials[1]

		if username != expectedUser || password != expectedPassword {
			unauthorized(c)
			return
		}

		c.Next()
	}
}

func unauthorized(c *gin.Context) {
	c.Header("WWW-Authenticate", `Basic realm="restricted"`)
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"error": "unauthorized",
	})
}
