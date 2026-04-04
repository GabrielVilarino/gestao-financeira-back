package services

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/exceptions"
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/entities"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
)

// Conversão entre schemas e entidades para usuário
func UserSchemaToEntity(userRequest schemas.CreateUserRequest) entities.User {
	return entities.User{
		Nome:     userRequest.Nome,
		Email:    userRequest.Email,
		Password: userRequest.Senha,
	}
}

func UserEntityToSchema(user entities.User) schemas.CreateUserResponse {
	return schemas.CreateUserResponse{
		ID:    *user.ID,
		Nome:  user.Nome,
		Email: user.Email,
	}
}

func ValidarIDGrupo(idGrupoRequest *int, idGrupoJWT *int) error {
	if idGrupoRequest == nil {
		return nil
	}
	jwtGroupID := 0
	if idGrupoJWT != nil {
		jwtGroupID = *idGrupoJWT
	}
	if *idGrupoRequest != jwtGroupID {
		return &exceptions.ValidationError{Message: "id_grupo não corresponde ao grupo do usuário autenticado"}
	}
	return nil
}
