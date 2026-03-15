package services

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/entities"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
)

// Conversão entre schemas e entidades para usuário
func UserSchemaToEntity(userRequest schemas.CreateUserRequest) entities.User {
	return entities.User{
		Name:     userRequest.Nome,
		Email:    userRequest.Email,
		Password: userRequest.Senha,
	}
}

func UserEntityToSchema(user entities.User) schemas.CreateUserResponse {
	return schemas.CreateUserResponse{
		ID:    *user.ID,
		Nome:  user.Name,
		Email: user.Email,
	}
}
