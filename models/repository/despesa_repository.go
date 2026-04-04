package repository

import (
	"fmt"

	"github.com/GabrielVilarino/gestao-financeira-back.git/models/config"
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/entities"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
)

func CreateDespesaRepository(despesa entities.Despesa) (*entities.Despesa, error) {
	query := `
		INSERT INTO despesa (id_usuario, id_grupo, id_categoria, id_subcategoria, tipo_transacao, valor, data_pagamento, data_ult_pagamento)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id_despesa, data_criacao
	`

	err := config.DB.QueryRow(
		query,
		despesa.IDUsuario,
		despesa.IDGrupo,
		despesa.IDCategoria,
		despesa.IDSubcategoria,
		despesa.TipoTransacao,
		despesa.Valor,
		despesa.DataPagamento,
		despesa.DataUltimoPagamento,
	).Scan(&despesa.ID, &despesa.DataCriacao)

	if err != nil {
		return nil, err
	}

	return &despesa, nil
}

func UpdateDespesaRepository(despesa entities.Despesa) (*entities.Despesa, error) {
	query := `
		UPDATE despesa
		SET id_categoria = $2, id_subcategoria = $3, tipo_transacao = $4, valor = $5, data_pagamento = $6, data_ult_pagamento = $7
		WHERE id_despesa = $1
		RETURNING data_criacao
	`

	err := config.DB.QueryRow(
		query,
		despesa.ID,
		despesa.IDCategoria,
		despesa.IDSubcategoria,
		despesa.TipoTransacao,
		despesa.Valor,
		despesa.DataPagamento,
		despesa.DataUltimoPagamento,
	).Scan(&despesa.DataCriacao)
	if err != nil {
		return nil, err
	}

	return &despesa, nil
}

func DeleteDespesaRepository(id int) error {
	query := `
		DELETE FROM despesa
		WHERE id_despesa = $1
	`

	_, err := config.DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func GetDespesasRepository(idUsuario int, dataInicio, dataFim *string, idGrupo *int) ([]schemas.GetDespesaResponse, error) {
	baseQuery := `
		SELECT 
			d.id_despesa as id,
			u.nome,
			CASE 
				WHEN d.tipo_transacao = 'fixa' 
				THEN date_trunc('month', CURRENT_DATE)
				ELSE d.data_pagamento
			END as data_pagamento,
			d.data_ult_pagamento,
			d.valor,
			c.nome as categoria,
			sc.nome as subcategoria,
			d.tipo_transacao
		FROM despesa d 
		JOIN usuario u 
		ON u.id_usuario = d.id_usuario
		JOIN categoria c 
		ON c.id_categoria = d.id_categoria 
		LEFT JOIN subcategoria sc 
		ON sc.id_subcategoria = d.id_subcategoria
	`

	args := []interface{}{idUsuario}
	conditions := " WHERE d.id_usuario = $1"
	paramIdx := 2
	addCondition := func(clause string) {
		conditions += " AND " + clause
	}

	if idGrupo != nil {
		addCondition(fmt.Sprintf("d.id_grupo = $%d", paramIdx))
		args = append(args, *idGrupo)
		paramIdx++
	} else {
		addCondition("d.id_grupo IS NULL")
	}

	if dataInicio != nil && dataFim != nil {
		addCondition(fmt.Sprintf(`
			(
				d.tipo_transacao = 'fixa'
				OR (
					d.tipo_transacao = 'parcelado'
					AND d.data_pagamento <= $%d
					AND d.data_ult_pagamento >= $%d
				)
				OR (
					d.tipo_transacao = 'variavel'
					AND d.data_pagamento BETWEEN $%d AND $%d
				)
			)
		`, paramIdx+1, paramIdx, paramIdx, paramIdx+1))

		args = append(args, *dataInicio, *dataFim)
		paramIdx += 2

	} else if dataInicio != nil {
		addCondition(fmt.Sprintf(`
			(
				d.tipo_transacao = 'fixa'
				OR (
					d.tipo_transacao = 'parcelado'
					AND d.data_ult_pagamento >= $%d
				)
				OR (
					d.tipo_transacao = 'variavel'
					AND d.data_pagamento >= $%d
				)
			)
		`, paramIdx, paramIdx))

		args = append(args, *dataInicio)
		paramIdx++

	} else if dataFim != nil {
		addCondition(fmt.Sprintf(`
			(
				d.tipo_transacao = 'fixa'
				OR (
					d.tipo_transacao = 'parcelado'
					AND d.data_pagamento <= $%d
				)
				OR (
					d.tipo_transacao = 'variavel'
					AND d.data_pagamento <= $%d
				)
			)
		`, paramIdx, paramIdx))

		args = append(args, *dataFim)
		paramIdx++
	}

	rows, err := config.DB.Query(baseQuery+conditions, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var despesas []schemas.GetDespesaResponse
	for rows.Next() {
		var d schemas.GetDespesaResponse
		err := rows.Scan(
			&d.ID,
			&d.Nome,
			&d.DataPagamento,
			&d.DataUltimoPagamento,
			&d.Valor,
			&d.Categoria,
			&d.Subcategoria,
			&d.TipoTransacao,
		)
		if err != nil {
			return nil, err
		}
		despesas = append(despesas, d)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return despesas, nil
}

func GetDespesasByIDRepository(id string) (*entities.Despesa, error) {
	query := `
		SELECT 
			*
		FROM despesa d
		WHERE d.id_despesa = $1
	`

	rows, err := config.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var despesa entities.Despesa

	if rows.Next() {
		err := rows.Scan(
			&despesa.ID,
			&despesa.IDUsuario,
			&despesa.IDGrupo,
			&despesa.IDCategoria,
			&despesa.IDSubcategoria,
			&despesa.TipoTransacao,
			&despesa.DataPagamento,
			&despesa.DataUltimoPagamento,
			&despesa.DataCriacao,
			&despesa.Valor,
		)
		if err != nil {
			return nil, err
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &despesa, nil
}
