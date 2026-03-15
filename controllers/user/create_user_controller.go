package user

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	"github.com/GabrielVilarino/gestao-financeira-back.git/services/user"
)

func CreateUserController(createUserRequest schemas.CreateUserRequest) (*schemas.CreateUserResponse, error) {
	// Controller para criar um usuário
	user, err := user.CreateUserService(createUserRequest)
	if err != nil {
		return nil, err
	}

	return &schemas.CreateUserResponse{
		ID:    user.ID,
		Nome:  user.Nome,
		Email: user.Email,
	}, nil
}
