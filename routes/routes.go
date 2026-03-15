package routes

import (
	"os"

	"github.com/GabrielVilarino/gestao-financeira-back.git/configs"
	"github.com/GabrielVilarino/gestao-financeira-back.git/routes/user"
	"github.com/GabrielVilarino/gestao-financeira-back.git/security"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes() {
	// Configuração do modo de execução do Gin
	gin.SetMode(gin.ReleaseMode)

	// Criação do router
	router := gin.Default()
	router.SetTrustedProxies(nil)

	// Agrupamento de rotas para a versão da API
	v1 := router.Group("/api/v1")

	// Rotas de usuário
	userGroup := v1.Group("/users")
	{
		userGroup.POST(
			"/create-user",
			security.BasicAuth(os.Getenv("BASIC_AUTH_USER"), os.Getenv("BASIC_AUTH_PASS")),
			user.CreateUserRoute,
		)
	}

	configs.Log.Info("==> Servidor Iniciado <==")
	// Inicia o servidor na porta 8080
	router.Run(":8080")
}
