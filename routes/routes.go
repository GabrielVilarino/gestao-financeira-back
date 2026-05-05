package routes

import (
	"os"

	"github.com/GabrielVilarino/gestao-financeira-back.git/configs"
	"github.com/GabrielVilarino/gestao-financeira-back.git/routes/auth"
	"github.com/GabrielVilarino/gestao-financeira-back.git/routes/categoria"
	"github.com/GabrielVilarino/gestao-financeira-back.git/routes/dashboard"
	"github.com/GabrielVilarino/gestao-financeira-back.git/routes/despesa"
	"github.com/GabrielVilarino/gestao-financeira-back.git/routes/ganho"
	"github.com/GabrielVilarino/gestao-financeira-back.git/routes/grupo"
	"github.com/GabrielVilarino/gestao-financeira-back.git/routes/subcategoria"
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

	// Rotas de autenticação
	userAuth := v1.Group("/auth")
	{
		userAuth.POST(
			"/login",
			auth.AuthUserRoute,
		)

		userAuth.POST(
			"/logout",
			security.AuthMiddleware(),
			auth.LogoutUserRoute,
		)
	}

	// Rotas de usuário
	userGroup := v1.Group("/users")
	{
		userGroup.POST(
			"/create-user",
			security.BasicAuth(os.Getenv("BASIC_AUTH_USER"), os.Getenv("BASIC_AUTH_PASS")),
			user.CreateUserRoute,
		)
	}

	// Rotas de ganhos
	ganhoGroup := v1.Group("/ganhos")
	ganhoGroup.Use(security.AuthMiddleware())
	{
		ganhoGroup.POST(
			"/create",
			ganho.CreateGanhoRoute,
		)

		ganhoGroup.GET(
			"",
			ganho.GetGanhosRoute,
		)

		ganhoGroup.GET(
			"/:id",
			ganho.GetGanhosByIDRoute,
		)

		ganhoGroup.PUT(
			"/update",
			ganho.UpdateGanhoRoute,
		)

		ganhoGroup.DELETE(
			"/delete/:id",
			ganho.DeleteGanhoRoute,
		)
	}

	// Rotas de despesas
	despesaGroup := v1.Group("/despesas")
	despesaGroup.Use(security.AuthMiddleware())
	{
		despesaGroup.POST(
			"/create",
			despesa.CreateDespesaRoute,
		)

		despesaGroup.GET(
			"",
			despesa.GetDespesasRoute,
		)

		despesaGroup.GET(
			"/:id",
			despesa.GetDespesaByIDRoute,
		)

		despesaGroup.PUT(
			"/update",
			despesa.UpdateDespesaRoute,
		)

		despesaGroup.DELETE(
			"/delete/:id",
			despesa.DeleteDespesaRoute,
		)
	}

	// Rotas de categorias
	categoriaGroup := v1.Group("/categorias")
	categoriaGroup.Use(security.AuthMiddleware())
	{
		categoriaGroup.GET(
			"",
			categoria.GetCategoriasRoute,
		)

		categoriaGroup.POST(
			"/create",
			categoria.CreateCategoriaRoute,
		)

		categoriaGroup.PUT(
			"/update",
			categoria.UpdateCategoriaRoute,
		)

		categoriaGroup.DELETE(
			"/delete/:id",
			categoria.DeleteCategoriaRoute,
		)
	}

	// Rotas de subcategorias
	subcategoriaGroup := v1.Group("/subcategorias")
	subcategoriaGroup.Use(security.AuthMiddleware())
	{
		subcategoriaGroup.GET(
			"",
			subcategoria.GetSubcategoriasRoute,
		)

		subcategoriaGroup.POST(
			"/create",
			subcategoria.CreateSubCategoriaRoute,
		)

		subcategoriaGroup.PUT(
			"/update",
			subcategoria.UpdateSubCategoriaRoute,
		)

		subcategoriaGroup.DELETE(
			"/delete/:id",
			subcategoria.DeleteSubCategoriaRoute,
		)
	}

	// Rotas de grupos
	grupoGroup := v1.Group("/grupos")
	grupoGroup.Use(security.AuthMiddleware())
	{
		grupoGroup.GET(
			"/:id",
			grupo.GetGrupoByIDRoute,
		)
		grupoGroup.POST(
			"/create",
			grupo.CreateGrupoRoute,
		)
		grupoGroup.GET(
			"/:id/participantes",
			grupo.GetParcipantesByIDRoute,
		)
		grupoGroup.POST(
			"/:id/participantes",
			grupo.AddParticipanteRoute,
		)
		grupoGroup.DELETE(
			"/:id/participantes/:id_participante",
			grupo.DeleteParticipanteRoute,
		)
		grupoGroup.PATCH(
			"/:id/participantes/:id_participante/admin",
			grupo.UpdateRoleParticipanteRoute,
		)
	}

	// Rotas de dashboard
	dashboardGroup := v1.Group("/dashboard")
	dashboardGroup.Use(security.AuthMiddleware())
	{
		dashboardGroup.GET(
			"/total-ganhos",
			dashboard.GetDashboardTotalGanhosRoute,
		)

		dashboardGroup.GET(
			"/total-despesas",
			dashboard.GetDashboardTotalDespesasRoute,
		)

		dashboardGroup.GET(
			"/saldo-liquido",
			dashboard.GetDashboardSaldoLiquidoRoute,
		)

		dashboardGroup.GET(
			"/evolucao-mensal",
			dashboard.GetDashboardEvolucaoMensalRoute,
		)
	}

	configs.Log.Infof("==> Servidor Iniciado na porta %s <==", os.Getenv("PORT"))
	// Inicia o servidor na porta definida na variável de ambiente PORT
	if err := router.Run(":" + os.Getenv("PORT")); err != nil {
		configs.Log.Error("Servidor Parou: " + err.Error())
	}
}
