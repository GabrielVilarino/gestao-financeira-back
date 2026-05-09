package transacao

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/entities"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
)

func transacaoSchemaToEntity(request schemas.CreateTransacaoRequest) entities.Transacao {
	status := request.Status
	if status == "" {
		status = "PENDENTE"
	}

	return entities.Transacao{
		IDCategoria:    request.IDCategoria,
		IDSubcategoria: request.IDSubcategoria,
		Tipo:           request.Tipo,
		Descricao:      request.Descricao,
		Observacao:     request.Observacao,
		Valor:          request.Valor,
		Competencia:    request.Competencia,
		DataVencimento: request.DataVencimento,
		DataPagamento:  request.DataPagamento,
		Status:         status,
	}
}

func transacaoUpdateSchemaToEntity(request schemas.UpdateTransacaoRequest) entities.Transacao {
	status := request.Status
	if status == "" {
		status = "PENDENTE"
	}

	return entities.Transacao{
		ID:             request.ID,
		IDGrupo:        request.IDGrupo,
		IDCategoria:    request.IDCategoria,
		IDSubcategoria: request.IDSubcategoria,
		Tipo:           request.Tipo,
		Descricao:      request.Descricao,
		Observacao:     request.Observacao,
		Valor:          request.Valor,
		Competencia:    request.Competencia,
		DataVencimento: request.DataVencimento,
		DataPagamento:  request.DataPagamento,
		Status:         status,
	}
}

func transacaoEntityToCreateSchema(t entities.Transacao) schemas.CreateTransacaoResponse {
	return schemas.CreateTransacaoResponse{
		ID:             t.ID,
		IDUsuario:      t.IDUsuario,
		IDGrupo:        t.IDGrupo,
		Escopo:         t.Escopo,
		IDCategoria:    t.IDCategoria,
		IDSubcategoria: t.IDSubcategoria,
		Tipo:           t.Tipo,
		Descricao:      t.Descricao,
		Observacao:     t.Observacao,
		Valor:          t.Valor,
		Competencia:    t.Competencia,
		DataVencimento: t.DataVencimento,
		DataPagamento:  t.DataPagamento,
		Status:         t.Status,
		CreatedAt:      t.CreatedAt,
	}
}

func transacaoEntityToUpdateSchema(t entities.Transacao) schemas.UpdateTransacaoResponse {
	return schemas.UpdateTransacaoResponse{
		ID:             t.ID,
		IDUsuario:      t.IDUsuario,
		IDGrupo:        t.IDGrupo,
		Escopo:         t.Escopo,
		IDCategoria:    t.IDCategoria,
		IDSubcategoria: t.IDSubcategoria,
		Tipo:           t.Tipo,
		Descricao:      t.Descricao,
		Observacao:     t.Observacao,
		Valor:          t.Valor,
		Competencia:    t.Competencia,
		DataVencimento: t.DataVencimento,
		DataPagamento:  t.DataPagamento,
		Status:         t.Status,
		CreatedAt:      t.CreatedAt,
		UpdatedAt:      t.UpdatedAt,
	}
}

func transacaoEntityToGetByIDSchema(t entities.Transacao) schemas.GetTransacaoByIDResponse {
	return schemas.GetTransacaoByIDResponse{
		IDCategoria:    t.IDCategoria,
		IDSubcategoria: t.IDSubcategoria,
		Tipo:           t.Tipo,
		Descricao:      t.Descricao,
		Observacao:     t.Observacao,
		Valor:          t.Valor,
		Competencia:    t.Competencia,
		DataVencimento: t.DataVencimento,
		DataPagamento:  t.DataPagamento,
		Status:         t.Status,
	}
}
