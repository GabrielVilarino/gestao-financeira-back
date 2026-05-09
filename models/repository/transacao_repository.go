package repository

import (
	"fmt"

	"github.com/GabrielVilarino/gestao-financeira-back.git/models/config"
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/entities"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
)

func CreateTransacaoRepository(transacao entities.Transacao) (*entities.Transacao, error) {
	return createTransacaoInternal(transacao)
}

func CreateTransacaoComRecorrenciaRepository(transacao entities.Transacao) (*entities.Transacao, error) {
	return createTransacaoInternal(transacao)
}

func createTransacaoInternal(transacao entities.Transacao) (*entities.Transacao, error) {
	query := `
		INSERT INTO transacao (id_usuario, id_grupo, escopo, id_categoria, id_subcategoria, id_recorrencia, tipo, descricao, observacao, valor, competencia, data_vencimento, data_pagamento, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
		RETURNING id_transacao, created_at
	`

	err := config.DB.QueryRow(
		query,
		transacao.IDUsuario,
		transacao.IDGrupo,
		transacao.Escopo,
		transacao.IDCategoria,
		transacao.IDSubcategoria,
		transacao.IDRecorrencia,
		transacao.Tipo,
		transacao.Descricao,
		transacao.Observacao,
		transacao.Valor,
		transacao.Competencia,
		transacao.DataVencimento,
		transacao.DataPagamento,
		transacao.Status,
	).Scan(&transacao.ID, &transacao.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &transacao, nil
}

func UpdateTransacaoRepository(transacao entities.Transacao) (*entities.Transacao, error) {
	query := `
		UPDATE transacao
		SET id_categoria = $2, id_subcategoria = $3, tipo = $4, descricao = $5, observacao = $6, valor = $7,
		    competencia = $8, data_vencimento = $9, data_pagamento = $10, status = $11, updated_at = CURRENT_TIMESTAMP
		WHERE id_transacao = $1
		RETURNING created_at, updated_at
	`

	err := config.DB.QueryRow(
		query,
		transacao.ID,
		transacao.IDCategoria,
		transacao.IDSubcategoria,
		transacao.Tipo,
		transacao.Descricao,
		transacao.Observacao,
		transacao.Valor,
		transacao.Competencia,
		transacao.DataVencimento,
		transacao.DataPagamento,
		transacao.Status,
	).Scan(&transacao.CreatedAt, &transacao.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &transacao, nil
}

func DeleteTransacaoRepository(id int) error {
	query := `DELETE FROM transacao WHERE id_transacao = $1`

	_, err := config.DB.Exec(query, id)
	return err
}

func GetTransacoesRepository(idUsuario int, dataInicio, dataFim *string, tipo *string, idGrupo *int) ([]schemas.GetTransacaoResponse, error) {
	baseQuery := `
		SELECT
			t.id_transacao AS id,
			u.nome,
			t.tipo,
			t.descricao,
			t.valor,
			t.competencia,
			t.status,
			c.nome AS categoria,
			sc.nome AS subcategoria
		FROM transacao t
		JOIN usuario u ON u.id_usuario = t.id_usuario
		JOIN categoria c ON c.id_categoria = t.id_categoria
		LEFT JOIN subcategoria sc ON sc.id_subcategoria = t.id_subcategoria
	`

	var args []interface{}
	var conditions string
	paramIdx := 1
	addCondition := func(clause string) {
		conditions += " AND " + clause
	}

	if idGrupo != nil {
		conditions = fmt.Sprintf(" WHERE t.id_grupo = $%d", paramIdx)
		args = append(args, *idGrupo)
		paramIdx++
	} else {
		conditions = fmt.Sprintf(" WHERE t.id_usuario = $%d AND t.id_grupo IS NULL", paramIdx)
		args = append(args, idUsuario)
		paramIdx++
	}

	if dataInicio != nil && dataFim != nil {
		addCondition(fmt.Sprintf("t.competencia BETWEEN $%d AND $%d", paramIdx, paramIdx+1))
		args = append(args, *dataInicio, *dataFim)
		paramIdx += 2
	} else if dataInicio != nil {
		addCondition(fmt.Sprintf("t.competencia >= $%d", paramIdx))
		args = append(args, *dataInicio)
		paramIdx++
	} else if dataFim != nil {
		addCondition(fmt.Sprintf("t.competencia <= $%d", paramIdx))
		args = append(args, *dataFim)
		paramIdx++
	}

	if tipo != nil {
		addCondition(fmt.Sprintf("t.tipo = $%d", paramIdx))
		args = append(args, *tipo)
		paramIdx++
	}

	conditions += " ORDER BY t.competencia DESC"

	rows, err := config.DB.Query(baseQuery+conditions, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transacoes []schemas.GetTransacaoResponse
	for rows.Next() {
		var t schemas.GetTransacaoResponse
		err := rows.Scan(
			&t.ID,
			&t.Nome,
			&t.Tipo,
			&t.Descricao,
			&t.Valor,
			&t.Competencia,
			&t.Status,
			&t.Categoria,
			&t.Subcategoria,
		)
		if err != nil {
			return nil, err
		}
		transacoes = append(transacoes, t)
	}

	if transacoes == nil {
		transacoes = []schemas.GetTransacaoResponse{}
	}

	return transacoes, nil
}

func GetTransacaoByIDRepository(id string) (*entities.Transacao, error) {
	query := `
		SELECT id_categoria, id_subcategoria, tipo, descricao, observacao, valor, competencia, data_vencimento, data_pagamento, status
		FROM transacao
		WHERE id_transacao = $1
	`

	var t entities.Transacao
	err := config.DB.QueryRow(query, id).Scan(
		&t.IDCategoria,
		&t.IDSubcategoria,
		&t.Tipo,
		&t.Descricao,
		&t.Observacao,
		&t.Valor,
		&t.Competencia,
		&t.DataVencimento,
		&t.DataPagamento,
		&t.Status,
	)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
