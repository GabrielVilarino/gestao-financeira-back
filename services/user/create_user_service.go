package user

import (
	"fmt"

	"github.com/GabrielVilarino/gestao-financeira-back.git/models/repository"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	"github.com/GabrielVilarino/gestao-financeira-back.git/services"
)

func CreateUserService(createUserRequest schemas.CreateUserRequest) (*schemas.CreateUserResponse, error) {
	// Service para criar um usuário
	if !validatePassword(createUserRequest.Senha) {
		return nil, fmt.Errorf("senha não atende os requisitos de segurança")
	}

	user := services.UserSchemaToEntity(createUserRequest)

	userCreated, err := repository.CreateUserRepository(user)
	if err != nil {
		return nil, err
	}

	userSchema := services.UserEntityToSchema(*userCreated)

	return &userSchema, nil
}
