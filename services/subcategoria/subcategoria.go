package subcategoria

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/entities"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
)

func SchemaCreateToEntity(request schemas.CreateSubCategoriaRequest) *entities.SubCategoria {
	return &entities.SubCategoria{
		IDCategoria: request.IDCategoria,
		Nome:        request.Nome,
	}
}

func SchemaUpdateToEntity(request schemas.UpdateSubCategoriaRequest) *entities.SubCategoria {
	return &entities.SubCategoria{
		ID:   request.ID,
		Nome: request.Nome,
	}
}

func EntityToSchema(subcategoria entities.SubCategoria) *schemas.SubCategoriaResponse {
	return &schemas.SubCategoriaResponse{
		ID:          subcategoria.ID,
		IDCategoria: subcategoria.IDCategoria,
		Nome:        subcategoria.Nome,
	}
}
