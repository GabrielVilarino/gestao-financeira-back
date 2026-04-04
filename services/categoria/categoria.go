package categoria

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/entities"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
)

func SchemaCreateToEntity(idUsuario int, request schemas.CreateCategoriaRequest) *entities.Categoria {
	return &entities.Categoria{
		Nome:             request.Nome,
		TipoMovimentacao: request.TipoMovimentacao,
		IDUsuario:        idUsuario,
	}
}

func SchemaUpdateToEntity(idUsuario int, request schemas.UpdateCategoriaRequest) *entities.Categoria {
	return &entities.Categoria{
		ID:               request.ID,
		Nome:             request.Nome,
		TipoMovimentacao: request.TipoMovimentacao,
		IDUsuario:        idUsuario,
	}
}

func EntityToSchema(categoria entities.Categoria) *schemas.CategoriaResponse {
	return &schemas.CategoriaResponse{
		ID:               categoria.ID,
		Nome:             categoria.Nome,
		TipoMovimentacao: categoria.TipoMovimentacao,
		IDUsuario:        &categoria.IDUsuario,
	}
}
