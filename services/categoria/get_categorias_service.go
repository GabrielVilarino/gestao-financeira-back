package categoria

import (
	"fmt"

	"github.com/GabrielVilarino/gestao-financeira-back.git/models/repository"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
)

func GetCategoriasService(idUsuario int, tipo *string) ([]schemas.CategoriaResponse, error) {
	categorias, err := repository.GetCategoriasRepository(idUsuario, tipo)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar categorias: %w", err)
	}

	return categorias, nil
}
