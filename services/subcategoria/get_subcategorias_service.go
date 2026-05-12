package subcategoria

import (
	"fmt"

	"github.com/GabrielVilarino/gestao-financeira-back.git/models/repository"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
)

func GetSubcategoriasService(idCategoria *int, idGrupo *int) ([]schemas.SubCategoriaResponse, error) {
	subcategorias, err := repository.GetSubcategoriasRepository(idCategoria, idGrupo)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar subcategorias: %w", err)
	}

	return subcategorias, nil
}
