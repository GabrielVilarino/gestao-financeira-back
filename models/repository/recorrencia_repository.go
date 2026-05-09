package repository

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/config"
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/entities"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
)

func CreateRecorrenciaRepository(r entities.Recorrencia) (*entities.Recorrencia, error) {
	query := `
		INSERT INTO recorrencia (frequencia, intervalo, data_inicio, data_fim, ativa)
		VALUES ($1, $2, $3, $4, true)
		RETURNING id_recorrencia, created_at
	`

	err := config.DB.QueryRow(
		query,
		r.Frequencia,
		r.Intervalo,
		r.DataInicio,
		r.DataFim,
	).Scan(&r.ID, &r.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func EncerrarRecorrenciaRepository(id int, dataFim string) error {
	query := `
		UPDATE recorrencia
		SET data_fim = $2, ativa = false
		WHERE id_recorrencia = $1
	`

	_, err := config.DB.Exec(query, id, dataFim)
	return err
}

func DeleteTransacoesFuturasByRecorrenciaRepository(idRecorrencia int, dataCorte string) error {
	query := `
		DELETE FROM transacao
		WHERE id_recorrencia = $1
		  AND competencia > $2::date
		  AND status = 'PENDENTE'
	`

	_, err := config.DB.Exec(query, idRecorrencia, dataCorte)
	return err
}

func GetRecorrenciasAtivasRepository() ([]entities.Recorrencia, error) {
	query := `
		SELECT
			r.id_recorrencia,
			t.id_usuario,
			t.id_grupo,
			t.escopo,
			t.id_categoria,
			t.id_subcategoria,
			t.tipo,
			t.descricao,
			t.observacao,
			t.valor,
			r.frequencia,
			r.intervalo,
			r.data_inicio,
			r.data_fim,
			r.ativa
		FROM recorrencia r
		JOIN transacao t ON t.id_recorrencia = r.id_recorrencia
		WHERE r.ativa = true
		  AND (r.data_fim IS NULL OR r.data_fim >= CURRENT_DATE)
		ORDER BY r.id_recorrencia, t.competencia DESC
	`

	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	seen := map[int]bool{}
	var recorrencias []entities.Recorrencia
	for rows.Next() {
		var rec entities.Recorrencia
		err := rows.Scan(
			&rec.ID,
			&rec.IDUsuario,
			&rec.IDGrupo,
			&rec.Escopo,
			&rec.IDCategoria,
			&rec.IDSubcategoria,
			&rec.Tipo,
			&rec.Descricao,
			&rec.Observacao,
			&rec.Valor,
			&rec.Frequencia,
			&rec.Intervalo,
			&rec.DataInicio,
			&rec.DataFim,
			&rec.Ativa,
		)
		if err != nil {
			return nil, err
		}
		if !seen[rec.ID] {
			seen[rec.ID] = true
			recorrencias = append(recorrencias, rec)
		}
	}

	return recorrencias, nil
}

func GetUltimaCompetenciaRecorrenciaRepository(idRecorrencia int) (*string, error) {
	query := `
		SELECT MAX(competencia)::text
		FROM transacao
		WHERE id_recorrencia = $1
	`

	var ultima *string
	err := config.DB.QueryRow(query, idRecorrencia).Scan(&ultima)
	if err != nil {
		return nil, err
	}

	return ultima, nil
}

func GetRecorrenciasByUsuarioRepository(idUsuario int, idGrupo *int) ([]schemas.GetRecorrenciaResponse, error) {
	query := `
		SELECT DISTINCT ON (r.id_recorrencia)
			r.id_recorrencia,
			t.id_categoria,
			t.id_subcategoria,
			t.tipo,
			t.descricao,
			t.observacao,
			t.valor,
			r.frequencia,
			r.intervalo,
			r.data_inicio::text,
			r.data_fim::text,
			r.ativa,
			c.nome AS categoria,
			sc.nome AS subcategoria
		FROM recorrencia r
		JOIN transacao t ON t.id_recorrencia = r.id_recorrencia
		JOIN categoria c ON c.id_categoria = t.id_categoria
		LEFT JOIN subcategoria sc ON sc.id_subcategoria = t.id_subcategoria
		WHERE (
			($2::int IS NOT NULL AND t.id_grupo = $2)
			OR ($2::int IS NULL AND t.id_usuario = $1 AND t.id_grupo IS NULL)
		)
		ORDER BY r.id_recorrencia
	`

	rows, err := config.DB.Query(query, idUsuario, idGrupo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []schemas.GetRecorrenciaResponse
	for rows.Next() {
		var rec schemas.GetRecorrenciaResponse
		err := rows.Scan(
			&rec.ID,
			&rec.IDCategoria,
			&rec.IDSubcategoria,
			&rec.Tipo,
			&rec.Descricao,
			&rec.Observacao,
			&rec.Valor,
			&rec.Frequencia,
			&rec.Intervalo,
			&rec.DataInicio,
			&rec.DataFim,
			&rec.Ativa,
			&rec.Categoria,
			&rec.Subcategoria,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, rec)
	}

	if result == nil {
		result = []schemas.GetRecorrenciaResponse{}
	}

	return result, nil
}
