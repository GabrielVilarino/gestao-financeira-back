package despesa

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/entities"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
)

func DespesaSchemaToEntity(request schemas.CreateDespesaRequest) entities.Despesa {
	return entities.Despesa{
		IDCategoria:         request.IDCategoria,
		IDSubcategoria:      request.IDSubcategoria,
		TipoTransacao:       request.TipoTransacao,
		DataPagamento:       request.DataPagamento,
		DataUltimoPagamento: request.DataUltimoPagamento,
		Valor:               request.Valor,
	}
}

func DespesaEntityToCreateSchema(despesa entities.Despesa) schemas.CreateDespesaResponse {
	return schemas.CreateDespesaResponse{
		ID:                  despesa.ID,
		IDUsuario:           despesa.IDUsuario,
		IDGrupo:             despesa.IDGrupo,
		IDCategoria:         despesa.IDCategoria,
		IDSubcategoria:      despesa.IDSubcategoria,
		TipoTransacao:       despesa.TipoTransacao,
		DataPagamento:       despesa.DataPagamento,
		DataUltimoPagamento: despesa.DataUltimoPagamento,
		DataCriacao:         despesa.DataCriacao,
		Valor:               despesa.Valor,
	}
}

func DespesaUpdateSchemaToEntity(request schemas.UpdateDespesaRequest) entities.Despesa {
	return entities.Despesa{
		ID:                  request.ID,
		IDGrupo:             request.IDGrupo,
		IDCategoria:         request.IDCategoria,
		IDSubcategoria:      request.IDSubcategoria,
		TipoTransacao:       request.TipoTransacao,
		DataPagamento:       request.DataPagamento,
		DataUltimoPagamento: request.DataUltimoPagamento,
		Valor:               request.Valor,
	}
}

func DespesaEntityToUpdateSchema(despesa entities.Despesa) schemas.UpdateDespesaResponse {
	return schemas.UpdateDespesaResponse{
		ID:                  despesa.ID,
		IDUsuario:           despesa.IDUsuario,
		IDGrupo:             despesa.IDGrupo,
		IDCategoria:         despesa.IDCategoria,
		IDSubcategoria:      despesa.IDSubcategoria,
		TipoTransacao:       despesa.TipoTransacao,
		Valor:               despesa.Valor,
		DataPagamento:       despesa.DataPagamento,
		DataUltimoPagamento: despesa.DataUltimoPagamento,
		DataCriacao:         despesa.DataCriacao,
	}
}

func DespesaEntityToGetSchema(despesa entities.Despesa) schemas.GetDespesaByIDResponse {
	return schemas.GetDespesaByIDResponse{
		IDCategoria:         despesa.IDCategoria,
		IDSubcategoria:      despesa.IDSubcategoria,
		TipoTransacao:       despesa.TipoTransacao,
		DataPagamento:       despesa.DataPagamento,
		DataUltimoPagamento: despesa.DataUltimoPagamento,
		Valor:               despesa.Valor,
	}
}
