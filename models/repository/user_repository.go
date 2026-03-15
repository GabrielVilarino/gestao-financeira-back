package repository

import (
	"database/sql"
	"fmt"

	"github.com/GabrielVilarino/gestao-financeira-back.git/configs"
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/config"
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/entities"
	"golang.org/x/crypto/bcrypt"
)

func GetUserByEmail(email string) (*entities.User, error) {
	query := `
		SELECT id_usuario, nome, email, senha, is_admin, id_grupo, data_criacao
		FROM usuario
		WHERE email = $1
	`

	var user entities.User
	err := config.DB.QueryRow(query, email).Scan(
		&user.ID,
		&user.Nome,
		&user.Email,
		&user.Password,
		&user.IsAdmin,
		&user.IdGroup,
		&user.DataCriacao,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
func CreateUserRepository(userEntity entities.User) (*entities.User, error) {
	user, err := GetUserByEmail(userEntity.Email)
	if err != nil {
		return nil, err
	}

	if user != nil {
		configs.Log.Warningf("Email %s já existe no banco de dados", userEntity.Email)
		return nil, fmt.Errorf("Erro ao criar email")
	}

	// Hasheia a senha antes de salvar no banco
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userEntity.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("erro ao hashear senha")
	}

	query := `
		INSERT INTO usuario (nome, email, senha)
		VALUES ($1, $2, $3)
		RETURNING id_usuario
	`
	var newUserID int
	err = config.DB.QueryRow(
		query,
		userEntity.Nome,
		userEntity.Email,
		string(hashedPassword),
	).Scan(&newUserID)
	if err != nil {
		return nil, err
	}

	userEntity.ID = &newUserID
	return &userEntity, nil
}
