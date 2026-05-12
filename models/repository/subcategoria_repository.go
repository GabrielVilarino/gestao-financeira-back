package repository

import (
	"fmt"

	"github.com/GabrielVilarino/gestao-financeira-back.git/models/config"
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/entities"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
)

func GetSubcategoriasRepository(idCategoria *int, idGrupo *int) ([]schemas.SubCategoriaResponse, error) {
	var query string
	var args []interface{}
	argIdx := 1

	if idGrupo != nil {
		query = `
			SELECT id_subcategoria, id_categoria, nome, id_grupo
			FROM subcategoria
			WHERE id_grupo = $1
		`
		args = []interface{}{*idGrupo}
		argIdx = 2

		if idCategoria != nil {
			query += fmt.Sprintf(" AND id_categoria = $%d", argIdx)
			args = append(args, *idCategoria)
		}
	} else {
		query = `
			SELECT id_subcategoria, id_categoria, nome, id_grupo
			FROM subcategoria
			WHERE id_grupo IS NULL
		`
		args = []interface{}{}

		if idCategoria != nil {
			query += fmt.Sprintf(" AND id_categoria = $%d", argIdx)
			args = append(args, *idCategoria)
		}
	}

	query += " ORDER BY nome"

	rows, err := config.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subcategorias []schemas.SubCategoriaResponse
	for rows.Next() {
		var s schemas.SubCategoriaResponse
		if err := rows.Scan(&s.ID, &s.IDCategoria, &s.Nome, &s.IDGrupo); err != nil {
			return nil, err
		}
		subcategorias = append(subcategorias, s)
	}

	return subcategorias, nil
}

func CreateSubcategoriaRepository(subcategoria entities.SubCategoria) (*entities.SubCategoria, error) {
	query := `
		INSERT INTO subcategoria (id_categoria, id_grupo, nome)
		VALUES ($1, $2, $3)
		RETURNING id_subcategoria, id_categoria, nome, id_grupo
	`

	var subcategoriaResponse entities.SubCategoria
	err := config.DB.QueryRow(
		query,
		subcategoria.IDCategoria,
		subcategoria.IDGrupo,
		subcategoria.Nome,
	).Scan(
		&subcategoriaResponse.ID,
		&subcategoriaResponse.IDCategoria,
		&subcategoriaResponse.Nome,
		&subcategoriaResponse.IDGrupo,
	)

	if err != nil {
		return nil, err
	}

	return &subcategoriaResponse, nil
}

func UpdateSubCategoriaRepository(subcategoria entities.SubCategoria) (*entities.SubCategoria, error) {
	query := `
		UPDATE subcategoria
		SET nome = $2
		WHERE id_subcategoria = $1
		RETURNING id_subcategoria, id_categoria, nome, id_grupo
	`

	var subcategoriaResponse entities.SubCategoria
	err := config.DB.QueryRow(
		query,
		subcategoria.ID,
		subcategoria.Nome,
	).Scan(
		&subcategoriaResponse.ID,
		&subcategoriaResponse.IDCategoria,
		&subcategoriaResponse.Nome,
		&subcategoriaResponse.IDGrupo,
	)

	if err != nil {
		return nil, err
	}

	return &subcategoriaResponse, nil
}

func DeleteSubCategoriaRepository(idSubcategoria int) error {
	query := `
		DELETE FROM subcategoria
		WHERE id_subcategoria = $1
	`

	result, err := config.DB.Exec(query, idSubcategoria)

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("subcategoria não encontrada")
	}

	return nil
}
