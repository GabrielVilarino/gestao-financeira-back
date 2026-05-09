package recorrencia

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/exceptions"
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/entities"
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/repository"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	"github.com/GabrielVilarino/gestao-financeira-back.git/services"
)

var frequenciasValidas = map[string]bool{
	"DIARIA":  true,
	"SEMANAL": true,
	"MENSAL":  true,
	"ANUAL":   true,
}

var tiposValidos = map[string]bool{
	"RECEITA": true,
	"DESPESA": true,
}

func CreateRecorrenciaService(idUsuario int, idGrupoJWT *int, req schemas.CreateRecorrenciaRequest) (*schemas.CreateRecorrenciaResponse, error) {
	if err := services.ValidarIDGrupo(req.IDGrupo, idGrupoJWT); err != nil {
		return nil, err
	}

	if !frequenciasValidas[req.Frequencia] {
		return nil, &exceptions.ValidationError{Message: "frequencia deve ser DIARIA, SEMANAL, MENSAL ou ANUAL"}
	}

	if !tiposValidos[req.Tipo] {
		return nil, &exceptions.ValidationError{Message: "tipo deve ser RECEITA ou DESPESA"}
	}

	rec := recorrenciaSchemaToEntity(req)
	rec.IDUsuario = idUsuario

	if req.IDGrupo != nil {
		rec.IDGrupo = req.IDGrupo
		rec.Escopo = "GRUPO"
	} else {
		rec.Escopo = "PESSOAL"
	}

	created, err := repository.CreateRecorrenciaRepository(rec)
	if err != nil {
		return nil, err
	}

	// Insere a primeira transação com competencia = data_inicio
	primeiraTransacao := entities.Transacao{
		IDUsuario:      rec.IDUsuario,
		IDGrupo:        rec.IDGrupo,
		Escopo:         rec.Escopo,
		IDCategoria:    rec.IDCategoria,
		IDSubcategoria: rec.IDSubcategoria,
		Tipo:           rec.Tipo,
		Descricao:      rec.Descricao,
		Observacao:     rec.Observacao,
		Valor:          rec.Valor,
		Competencia:    rec.DataInicio,
		Status:         "PENDENTE",
	}
	primeiraTransacao.IDRecorrencia = &created.ID

	if _, err := repository.CreateTransacaoComRecorrenciaRepository(primeiraTransacao); err != nil {
		return nil, err
	}

	created.Ativa = true
	response := recorrenciaEntityToCreateSchema(*created)
	return &response, nil
}
