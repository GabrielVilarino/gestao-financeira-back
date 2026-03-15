package main

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/configs"
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/config"
	"github.com/GabrielVilarino/gestao-financeira-back.git/routes"
)

func main() {
	// Garante que a conexão com o banco será fechada ao finalizar
	defer config.CloseDatabase()

	// Carrega as configurações
	err := configs.InitializeConfigs()
	if err != nil {
		panic("Erro ao carregar as configurações: " + err.Error())
	}

	// Inicializa o banco de dados
	err = config.InitializeDatabase()
	if err != nil {
		panic(err.Error())
	}

	// Inicia o servidor
	routes.InitializeRoutes()
}
