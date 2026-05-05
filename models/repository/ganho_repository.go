package repository

import (
	"fmt"

	"github.com/GabrielVilarino/gestao-financeira-back.git/models/config"
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/entities"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
)

func CreateGanhoRepository(ganho entities.Ganho) (*entities.Ganho, error) {
	query := `
		INSERT INTO receita (id_usuario, id_grupo, id_categoria, id_subcategoria, tipo_transacao, valor, data_recebimento)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id_receita, data_criacao
	`

	err := config.DB.QueryRow(
		query,
		ganho.IDUsuario,
		ganho.IDGrupo,
		ganho.IDCategoria,
		ganho.IDSubcategoria,
		ganho.TipoTransacao,
		ganho.Valor,
		ganho.DataRecebimento,
	).Scan(&ganho.ID, &ganho.DataCriacao)
	if err != nil {
		return nil, err
	}

	return &ganho, nil
}

func UpdateGanhoRepository(ganho entities.Ganho) (*entities.Ganho, error) {
	query := `
		UPDATE receita
		SET id_usuario = $2, id_grupo = $3, id_categoria = $4, id_subcategoria = $5, tipo_transacao = $6, valor = $7, data_recebimento = $8
		WHERE id_receita = $1
		RETURNING data_criacao
	`

	err := config.DB.QueryRow(
		query,
		ganho.ID,
		ganho.IDUsuario,
		ganho.IDGrupo,
		ganho.IDCategoria,
		ganho.IDSubcategoria,
		ganho.TipoTransacao,
		ganho.Valor,
		ganho.DataRecebimento,
	).Scan(&ganho.DataCriacao)
	if err != nil {
		return nil, err
	}

	return &ganho, nil
}

func DeleteGanhoRepository(id int) error {
	query := `
		DELETE FROM receita
		WHERE id_receita = $1
	`

	_, err := config.DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func GetGanhosRepository(idUsuario int, dataInicio, dataFim *string, idGrupo *int) ([]schemas.GetGanhoResponse, error) {
	baseQuery := `
		SELECT 
			r.id_receita as id,
			u.nome,
			r.data_recebimento as data,
			r.valor,
			c.nome as tipo
		FROM receita r 
		JOIN usuario u
		ON u.id_usuario = r.id_usuario
		JOIN categoria c
		ON c.id_categoria = r.id_categoria
	`

	var args []interface{}
	var conditions string
	paramIdx := 1
	addCondition := func(clause string) {
		conditions += " AND " + clause
	}

	if idGrupo != nil {
		conditions = fmt.Sprintf(" WHERE r.id_grupo = $%d", paramIdx)
		args = append(args, *idGrupo)
		paramIdx++
	} else {
		conditions = fmt.Sprintf(" WHERE r.id_usuario = $%d AND r.id_grupo IS NULL", paramIdx)
		args = append(args, idUsuario)
		paramIdx++
	}

	if dataInicio != nil {
		addCondition(fmt.Sprintf("r.data_recebimento >= $%d", paramIdx))
		args = append(args, *dataInicio)
		paramIdx++
	}
	if dataFim != nil {
		addCondition(fmt.Sprintf("r.data_recebimento <= $%d", paramIdx))
		args = append(args, *dataFim)
	}

	rows, err := config.DB.Query(baseQuery+conditions, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ganhos []schemas.GetGanhoResponse
	for rows.Next() {
		var g schemas.GetGanhoResponse
		err := rows.Scan(
			&g.ID,
			&g.Nome,
			&g.DataRecebimento,
			&g.Valor,
			&g.TipoTransacao,
		)
		if err != nil {
			return nil, err
		}
		ganhos = append(ganhos, g)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ganhos, nil
}

func GetGanhosByIDRepository(id string) (*entities.Ganho, error) {
	query := `
		SELECT 
			*
		FROM receita r
		WHERE r.id_receita = $1
	`

	rows, err := config.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ganho entities.Ganho

	if rows.Next() {
		err := rows.Scan(
			&ganho.ID,
			&ganho.IDUsuario,
			&ganho.IDGrupo,
			&ganho.IDCategoria,
			&ganho.IDSubcategoria,
			&ganho.TipoTransacao,
			&ganho.DataRecebimento,
			&ganho.DataCriacao,
			&ganho.Valor,
		)
		if err != nil {
			return nil, err
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &ganho, nil
}
