package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LogoutUserRoute(c *gin.Context) {
	// Rota para logout do usuário
	c.SetCookie(
		"auth_token",
		"",
		-1,
		"/",
		"",
		false,
		true,
	)
	c.Status(http.StatusOK)
}
