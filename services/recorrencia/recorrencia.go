package recorrencia

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/entities"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
)

func recorrenciaSchemaToEntity(req schemas.CreateRecorrenciaRequest) entities.Recorrencia {
	intervalo := req.Intervalo
	if intervalo <= 0 {
		intervalo = 1
	}

	return entities.Recorrencia{
		IDCategoria:    req.IDCategoria,
		IDSubcategoria: req.IDSubcategoria,
		Tipo:           req.Tipo,
		Descricao:      req.Descricao,
		Observacao:     req.Observacao,
		Valor:          req.Valor,
		Frequencia:     req.Frequencia,
		Intervalo:      intervalo,
		DataInicio:     req.DataInicio,
		DataFim:        req.DataFim,
	}
}

func recorrenciaEntityToCreateSchema(r entities.Recorrencia) schemas.CreateRecorrenciaResponse {
	return schemas.CreateRecorrenciaResponse{
		ID:             r.ID,
		IDUsuario:      r.IDUsuario,
		IDGrupo:        r.IDGrupo,
		Escopo:         r.Escopo,
		IDCategoria:    r.IDCategoria,
		IDSubcategoria: r.IDSubcategoria,
		Tipo:           r.Tipo,
		Descricao:      r.Descricao,
		Observacao:     r.Observacao,
		Valor:          r.Valor,
		Frequencia:     r.Frequencia,
		Intervalo:      r.Intervalo,
		DataInicio:     r.DataInicio,
		DataFim:        r.DataFim,
		Ativa:          r.Ativa,
		CreatedAt:      r.CreatedAt,
	}
}
