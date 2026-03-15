package security

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID  int    `json:"user_id"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin"`
	jwt.RegisteredClaims
}

func BearerAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extrai o token do header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token não fornecido"})
			c.Abort()
			return
		}

		// Verifica se o header está no formato "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato de token inválido"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// Valida o token
		claims, err := validateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido ou expirado"})
			c.Abort()
			return
		}

		// Adiciona as informações do usuário no contexto
		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("is_admin", claims.IsAdmin)

		// Continua para o próximo handler
		c.Next()
	}
}

// AdminMiddleware middleware para validar se o usuário é admin
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Verifica se o usuário é admin
		isAdmin, exists := c.Get("is_admin")
		if !exists || !isAdmin.(bool) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Acesso negado: permissão de administrador necessária"})
			c.Abort()
			return
		}

		c.Next()
	}
}
