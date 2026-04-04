package categoria

import (
	"fmt"

	"github.com/GabrielVilarino/gestao-financeira-back.git/models/repository"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
)

func GetCategoriasService(tipo *string) ([]schemas.GetCategoriaResponse, error) {
	categorias, err := repository.GetCategoriasRepository(tipo)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar categorias: %w", err)
	}

	return categorias, nil
}

func GetSubcategoriasService(idCategoria *int) ([]schemas.GetSubcategoriaResponse, error) {
	subcategorias, err := repository.GetSubcategoriasRepository(idCategoria)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar subcategorias: %w", err)
	}

	return subcategorias, nil
}
