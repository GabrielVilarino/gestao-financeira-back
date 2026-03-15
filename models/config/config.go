package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitializeDatabase() error {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	// Valor padrão para SSL Mode
	if dbSSLMode == "" {
		dbSSLMode = "disable"
	}

	// Validação das variáveis obrigatórias
	if dbHost == "" || dbPort == "" || dbUser == "" || dbPassword == "" || dbName == "" {
		return fmt.Errorf("variáveis de ambiente do banco de dados não configuradas corretamente")
	}

	// String de conexão PostgreSQL
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode,
	)

	// Abre a conexão com o banco
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("erro ao abrir conexão com o banco: %v", err)
	}

	// Testa a conexão
	err = db.Ping()
	if err != nil {
		return fmt.Errorf("erro ao conectar com o banco: %v", err)
	}

	// Configurações de pool de conexões
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	DB = db
	fmt.Println("Conexão com banco de dados estabelecida com sucesso")

	return nil
}

func CloseDatabase() {
	if DB != nil {
		DB.Close()
		fmt.Println("Conexão com banco de dados fechada")
	}
}
