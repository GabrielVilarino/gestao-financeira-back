package auth

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	"github.com/GabrielVilarino/gestao-financeira-back.git/services/auth"
)

func AuthUserController(authUserRequest schemas.AuthUserRequest) (*schemas.AuthUserResponse, *string, error) {
	// Controller para autenticar um usuário
	user, token, err := auth.AuthUserService(authUserRequest.Email, authUserRequest.Senha)
	if err != nil {
		return nil, nil, err
	}

	return &schemas.AuthUserResponse{
		ID:      *user.ID,
		Nome:    user.Nome,
		Email:   user.Email,
		IsAdmin: *user.IsAdmin,
		IdGroup: user.IdGroup,
	}, token, nil
}
