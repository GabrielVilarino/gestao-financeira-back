package repository

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/config"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
)

func GetDashboardTotalGanhosRepository(idUsuario int, idGrupo *int, dataInicio, dataFim string) (float64, error) {
	query := `
		WITH meses AS (
			SELECT generate_series(
				date_trunc('month', $3::date),
				date_trunc('month', $4::date),
				interval '1 month'
			)::date AS mes
		)
		SELECT COALESCE(SUM(total_mes), 0)
		FROM (
			SELECT
				m.mes,
				SUM(
					CASE
						WHEN r.tipo_transacao = 'fixa'
							AND r.data_recebimento <= (m.mes + interval '1 month - 1 day')::date
						THEN r.valor
						WHEN r.tipo_transacao = 'variavel'
							AND date_trunc('month', r.data_recebimento)::date = m.mes
						THEN r.valor
						ELSE 0
					END
				) AS total_mes
			FROM meses m
			CROSS JOIN receita r
			WHERE (
					($2::int IS NOT NULL AND r.id_grupo = $2)
					OR ($2::int IS NULL AND r.id_usuario = $1 AND r.id_grupo IS NULL)
				)
				AND (
					r.tipo_transacao = 'fixa'
					OR (
						r.tipo_transacao = 'variavel'
						AND r.data_recebimento BETWEEN $3::date AND $4::date
					)
				)
			GROUP BY m.mes
		) ganhos_por_mes
	`

	var totalGanhos float64
	err := config.DB.QueryRow(query, idUsuario, idGrupo, dataInicio, dataFim).Scan(&totalGanhos)
	if err != nil {
		return 0, err
	}

	return totalGanhos, nil
}

func GetDashboardTotalDespesasRepository(idUsuario int, idGrupo *int, dataInicio, dataFim string) (float64, error) {
	query := `
		WITH meses AS (
			SELECT generate_series(
				date_trunc('month', $3::date),
				date_trunc('month', $4::date),
				interval '1 month'
			)::date AS mes
		)
		SELECT COALESCE(SUM(total_mes), 0)
		FROM (
			SELECT 
				m.mes,
				SUM(
					CASE
						WHEN d.tipo_transacao = 'fixa' THEN d.valor
						WHEN d.tipo_transacao = 'parcelado'
							AND d.data_pagamento <= (m.mes + interval '1 month - 1 day')::date
							AND COALESCE(d.data_ult_pagamento, d.data_pagamento) >= m.mes
						THEN d.valor
						WHEN d.tipo_transacao = 'variavel'
							AND date_trunc('month', d.data_pagamento)::date = m.mes
						THEN d.valor
						ELSE 0
					END
				) AS total_mes
			FROM meses m
			CROSS JOIN despesa d
			WHERE (
					($2::int IS NOT NULL AND d.id_grupo = $2)
					OR ($2::int IS NULL AND d.id_usuario = $1 AND d.id_grupo IS NULL)
				)
				AND (
					d.tipo_transacao = 'fixa'
					OR (
						d.tipo_transacao = 'parcelado'
						AND d.data_pagamento <= $4::date
						AND COALESCE(d.data_ult_pagamento, d.data_pagamento) >= $3::date
					)
					OR (
						d.tipo_transacao = 'variavel'
						AND d.data_pagamento BETWEEN $3::date AND $4::date
					)
				)
			GROUP BY m.mes
		) despesas_por_mes
	`

	var totalDespesas float64
	err := config.DB.QueryRow(query, idUsuario, idGrupo, dataInicio, dataFim).Scan(&totalDespesas)
	if err != nil {
		return 0, err
	}

	return totalDespesas, nil
}

func GetDashboardSaldoLiquidoRepository(idUsuario int, idGrupo *int, dataInicio, dataFim string) (float64, error) {
	query := `
		WITH total_ganhos AS (
			WITH meses AS (
				SELECT generate_series(
					date_trunc('month', $3::date),
					date_trunc('month', $4::date),
					interval '1 month'
				)::date AS mes
			)
			SELECT COALESCE(SUM(total_mes), 0) AS valor
			FROM (
				SELECT
					m.mes,
					SUM(
						CASE
							WHEN r.tipo_transacao = 'fixa'
								AND r.data_recebimento <= (m.mes + interval '1 month - 1 day')::date
							THEN r.valor
							WHEN r.tipo_transacao = 'variavel'
								AND date_trunc('month', r.data_recebimento)::date = m.mes
							THEN r.valor
							ELSE 0
						END
					) AS total_mes
				FROM meses m
				CROSS JOIN receita r
				WHERE (
						($2::int IS NOT NULL AND r.id_grupo = $2)
						OR ($2::int IS NULL AND r.id_usuario = $1 AND r.id_grupo IS NULL)
					)
					AND (
						r.tipo_transacao = 'fixa'
						OR (
							r.tipo_transacao = 'variavel'
							AND r.data_recebimento BETWEEN $3::date AND $4::date
						)
					)
				GROUP BY m.mes
			) ganhos_por_mes
		),
		total_despesas AS (
			WITH meses AS (
				SELECT generate_series(
					date_trunc('month', $3::date),
					date_trunc('month', $4::date),
					interval '1 month'
				)::date AS mes
			)
			SELECT COALESCE(SUM(total_mes), 0) AS valor
			FROM (
				SELECT 
					m.mes,
					SUM(
						CASE
							WHEN d.tipo_transacao = 'fixa' THEN d.valor
							WHEN d.tipo_transacao = 'parcelado'
								AND d.data_pagamento <= (m.mes + interval '1 month - 1 day')::date
								AND COALESCE(d.data_ult_pagamento, d.data_pagamento) >= m.mes
							THEN d.valor
							WHEN d.tipo_transacao = 'variavel'
								AND date_trunc('month', d.data_pagamento)::date = m.mes
							THEN d.valor
							ELSE 0
						END
					) AS total_mes
				FROM meses m
				CROSS JOIN despesa d
				WHERE (
						($2::int IS NOT NULL AND d.id_grupo = $2)
						OR ($2::int IS NULL AND d.id_usuario = $1 AND d.id_grupo IS NULL)
					)
					AND (
						d.tipo_transacao = 'fixa'
						OR (
							d.tipo_transacao = 'parcelado'
							AND d.data_pagamento <= $4::date
							AND COALESCE(d.data_ult_pagamento, d.data_pagamento) >= $3::date
						)
						OR (
							d.tipo_transacao = 'variavel'
							AND d.data_pagamento BETWEEN $3::date AND $4::date
						)
					)
				GROUP BY m.mes
			) despesas_por_mes
		)
		SELECT total_ganhos.valor - total_despesas.valor
		FROM total_ganhos, total_despesas
	`

	var saldoLiquido float64
	err := config.DB.QueryRow(query, idUsuario, idGrupo, dataInicio, dataFim).Scan(&saldoLiquido)
	if err != nil {
		return 0, err
	}

	return saldoLiquido, nil
}

func GetDashboardEvolucaoMensalRepository(idUsuario int, idGrupo *int, dataInicio, dataFim string) ([]schemas.GetDashboardEvolucaoMensalItemResponse, error) {
	query := `
		WITH meses AS (
			SELECT generate_series(
				date_trunc('month', $3::date),
				date_trunc('month', $4::date),
				interval '1 month'
			)::date AS mes
		),
		ganhos AS (
			SELECT 
				m.mes,
				COALESCE(SUM(
					CASE
						WHEN r.tipo_transacao = 'fixa'
							AND r.data_recebimento <= (m.mes + interval '1 month - 1 day')::date
						THEN r.valor
						WHEN r.tipo_transacao = 'variavel'
							AND date_trunc('month', r.data_recebimento)::date = m.mes
						THEN r.valor
						ELSE 0
					END
				), 0) AS total_ganhos
			FROM meses m
			CROSS JOIN receita r
			WHERE (
					($2::int IS NOT NULL AND r.id_grupo = $2)
					OR ($2::int IS NULL AND r.id_usuario = $1 AND r.id_grupo IS NULL)
				)
				AND (
					r.tipo_transacao = 'fixa'
					OR (
						r.tipo_transacao = 'variavel'
						AND r.data_recebimento BETWEEN $3::date AND $4::date
					)
				)
			GROUP BY m.mes
		),
		despesas AS (
			SELECT 
				m.mes,
				COALESCE(SUM(
					CASE
						WHEN d.tipo_transacao = 'fixa' THEN d.valor
						WHEN d.tipo_transacao = 'parcelado'
							AND d.data_pagamento <= (m.mes + interval '1 month - 1 day')::date
							AND COALESCE(d.data_ult_pagamento, d.data_pagamento) >= m.mes
						THEN d.valor
						WHEN d.tipo_transacao = 'variavel'
							AND date_trunc('month', d.data_pagamento)::date = m.mes
						THEN d.valor
						ELSE 0
					END
				), 0) AS total_despesas
			FROM meses m
			CROSS JOIN despesa d
			WHERE (
					($2::int IS NOT NULL AND d.id_grupo = $2)
					OR ($2::int IS NULL AND d.id_usuario = $1 AND d.id_grupo IS NULL)
				)
				AND (
					d.tipo_transacao = 'fixa'
					OR (
						d.tipo_transacao = 'parcelado'
						AND d.data_pagamento <= $4::date
						AND COALESCE(d.data_ult_pagamento, d.data_pagamento) >= $3::date
					)
					OR (
						d.tipo_transacao = 'variavel'
						AND d.data_pagamento BETWEEN $3::date AND $4::date
					)
				)
			GROUP BY m.mes
		)
		SELECT 
			TO_CHAR(m.mes, 'YYYY-MM') AS mes,
			COALESCE(g.total_ganhos, 0) AS total_ganhos,
			COALESCE(d.total_despesas, 0) AS total_despesas,
			COALESCE(g.total_ganhos, 0) - COALESCE(d.total_despesas, 0) AS saldo_liquido
		FROM meses m
		LEFT JOIN ganhos g ON g.mes = m.mes
		LEFT JOIN despesas d ON d.mes = m.mes
		ORDER BY m.mes
	`

	rows, err := config.DB.Query(query, idUsuario, idGrupo, dataInicio, dataFim)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var evolucaoMensal []schemas.GetDashboardEvolucaoMensalItemResponse
	for rows.Next() {
		var item schemas.GetDashboardEvolucaoMensalItemResponse
		if err := rows.Scan(
			&item.Mes,
			&item.TotalGanhos,
			&item.TotalDespesas,
			&item.SaldoLiquido,
		); err != nil {
			return nil, err
		}
		evolucaoMensal = append(evolucaoMensal, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return evolucaoMensal, nil
}
