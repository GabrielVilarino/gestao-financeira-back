package user

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	"github.com/GabrielVilarino/gestao-financeira-back.git/services/user"
)

func AuthUserController(authUserRequest schemas.AuthUserRequest) (*schemas.AuthUserResponse, error) {
	// Controller para autenticar um usuário
	token, err := user.AuthUserService(authUserRequest.Email, authUserRequest.Senha)
	if err != nil {
		return nil, err
	}

	return &schemas.AuthUserResponse{
		Token: *token,
	}, nil
}
