package configs

import (
	"log"

	"github.com/joho/godotenv"
)

func initializeEnvVars() {

	err := godotenv.Overload(".env")

	if err != nil {
		log.Println("Erro ao carregar .env: ", err)
	}
}
