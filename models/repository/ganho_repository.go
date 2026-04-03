package repository

import (
	"fmt"

	"github.com/GabrielVilarino/gestao-financeira-back.git/models/config"
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/entities"
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

func GetGanhosRepository(dataInicio, dataFim *string, idGrupo *int) ([]entities.Ganho, error) {
	baseQuery := `
		SELECT id_receita, id_usuario, id_grupo, id_categoria, id_subcategoria, tipo_transacao, valor, data_recebimento, data_criacao
		FROM receita
	`

	args := []interface{}{}
	conditions := ""
	paramIdx := 1
	addCondition := func(clause string) {
		if paramIdx == 1 {
			conditions += " WHERE " + clause
		} else {
			conditions += " AND " + clause
		}
	}

	if dataInicio != nil {
		addCondition(fmt.Sprintf("data_recebimento >= $%d", paramIdx))
		args = append(args, *dataInicio)
		paramIdx++
	}
	if dataFim != nil {
		addCondition(fmt.Sprintf("data_recebimento <= $%d", paramIdx))
		args = append(args, *dataFim)
		paramIdx++
	}
	if idGrupo != nil {
		addCondition(fmt.Sprintf("id_grupo = $%d", paramIdx))
		args = append(args, *idGrupo)
	}

	rows, err := config.DB.Query(baseQuery+conditions, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ganhos []entities.Ganho
	for rows.Next() {
		var g entities.Ganho
		err := rows.Scan(
			&g.ID,
			&g.IDUsuario,
			&g.IDGrupo,
			&g.IDCategoria,
			&g.IDSubcategoria,
			&g.TipoTransacao,
			&g.Valor,
			&g.DataRecebimento,
			&g.DataCriacao,
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
