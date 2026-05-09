package repository

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/config"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
)

func GetDashboardTotalGanhosRepository(idUsuario int, idGrupo *int, dataInicio, dataFim string) (float64, error) {
	query := `
		SELECT COALESCE(SUM(valor), 0)
		FROM transacao
		WHERE tipo = 'RECEITA'
			AND (
				($2::int IS NOT NULL AND id_grupo = $2)
				OR ($2::int IS NULL AND id_usuario = $1 AND id_grupo IS NULL)
			)
			AND competencia BETWEEN $3::date AND $4::date
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
		SELECT COALESCE(SUM(valor), 0)
		FROM transacao
		WHERE tipo = 'DESPESA'
			AND (
				($2::int IS NOT NULL AND id_grupo = $2)
				OR ($2::int IS NULL AND id_usuario = $1 AND id_grupo IS NULL)
			)
			AND competencia BETWEEN $3::date AND $4::date
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
		SELECT
			COALESCE(SUM(CASE WHEN tipo = 'RECEITA' THEN valor ELSE 0 END), 0) -
			COALESCE(SUM(CASE WHEN tipo = 'DESPESA' THEN valor ELSE 0 END), 0)
		FROM transacao
		WHERE (
				($2::int IS NOT NULL AND id_grupo = $2)
				OR ($2::int IS NULL AND id_usuario = $1 AND id_grupo IS NULL)
			)
			AND competencia BETWEEN $3::date AND $4::date
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
		SELECT
			TO_CHAR(date_trunc('month', competencia), 'YYYY-MM') AS mes,
			COALESCE(SUM(CASE WHEN tipo = 'RECEITA' THEN valor ELSE 0 END), 0) AS total_ganhos,
			COALESCE(SUM(CASE WHEN tipo = 'DESPESA' THEN valor ELSE 0 END), 0) AS total_despesas,
			COALESCE(SUM(CASE WHEN tipo = 'RECEITA' THEN valor ELSE 0 END), 0) -
			COALESCE(SUM(CASE WHEN tipo = 'DESPESA' THEN valor ELSE 0 END), 0) AS saldo_liquido
		FROM transacao
		WHERE (
				($2::int IS NOT NULL AND id_grupo = $2)
				OR ($2::int IS NULL AND id_usuario = $1 AND id_grupo IS NULL)
			)
			AND competencia BETWEEN $3::date AND $4::date
		GROUP BY date_trunc('month', competencia)
		ORDER BY date_trunc('month', competencia)
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
