package ganho

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/entities"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
)

func GanhoSchemaToEntity(request schemas.CreateGanhoRequest) entities.Ganho {
	return entities.Ganho{
		IDUsuario:       request.IDUsuario,
		IDGrupo:         request.IDGrupo,
		IDCategoria:     request.IDCategoria,
		IDSubcategoria:  request.IDSubcategoria,
		TipoTransacao:   request.TipoTransacao,
		Valor:           request.Valor,
		DataRecebimento: request.DataRecebimento,
	}
}

func GanhoUpdateSchemaToEntity(request schemas.UpdateGanhoRequest) entities.Ganho {
	return entities.Ganho{
		ID:              request.ID,
		IDUsuario:       request.IDUsuario,
		IDGrupo:         request.IDGrupo,
		IDCategoria:     request.IDCategoria,
		IDSubcategoria:  request.IDSubcategoria,
		TipoTransacao:   request.TipoTransacao,
		Valor:           request.Valor,
		DataRecebimento: request.DataRecebimento,
	}
}

func GanhoEntityToCreateSchema(ganho entities.Ganho) schemas.CreateGanhoResponse {
	return schemas.CreateGanhoResponse{
		ID:              ganho.ID,
		IDUsuario:       ganho.IDUsuario,
		IDGrupo:         ganho.IDGrupo,
		IDCategoria:     ganho.IDCategoria,
		IDSubcategoria:  ganho.IDSubcategoria,
		TipoTransacao:   ganho.TipoTransacao,
		Valor:           ganho.Valor,
		DataRecebimento: ganho.DataRecebimento,
		DataCriacao:     ganho.DataCriacao,
	}
}

func GanhoEntityToUpdateSchema(ganho entities.Ganho) schemas.UpdateGanhoResponse {
	return schemas.UpdateGanhoResponse{
		ID:              ganho.ID,
		IDUsuario:       ganho.IDUsuario,
		IDGrupo:         ganho.IDGrupo,
		IDCategoria:     ganho.IDCategoria,
		IDSubcategoria:  ganho.IDSubcategoria,
		TipoTransacao:   ganho.TipoTransacao,
		Valor:           ganho.Valor,
		DataRecebimento: ganho.DataRecebimento,
		DataCriacao:     ganho.DataCriacao,
	}
}
