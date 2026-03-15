package user

import (
	"fmt"

	"github.com/GabrielVilarino/gestao-financeira-back.git/models/entities"
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/repository"
	"github.com/GabrielVilarino/gestao-financeira-back.git/security"
	"golang.org/x/crypto/bcrypt"
)

func AuthUserService(email, senha string) (*entities.User, *string, error) {
	// Service para autenticar um usuário
	user, err := repository.GetUserByEmail(email)
	if err != nil {
		return nil, nil, fmt.Errorf("Erro ao buscar usuário")
	}

	if user == nil {
		return nil, nil, fmt.Errorf("Usuário não encontrado")
	}

	// Verifica se a senha está correta comparando com o hash armazenado
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(senha))
	if err != nil {
		return nil, nil, fmt.Errorf("Credenciais inválidas")
	}

	// Gera o token JWT
	token, err := security.GenerateToken(*user.ID, user.Email, *user.IsAdmin)
	if err != nil {
		return nil, nil, fmt.Errorf("Erro ao gerar token")
	}

	return user, &token, nil
}
