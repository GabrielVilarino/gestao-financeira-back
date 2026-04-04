package repository

import (
	"fmt"

	"github.com/GabrielVilarino/gestao-financeira-back.git/models/config"
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/entities"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
)

func GetCategoriasRepository(idUsuario int, tipo *string) ([]schemas.CategoriaResponse, error) {
	query := `
		SELECT id_categoria, nome, tipo_movimentacao
		FROM categoria
		WHERE id_usuario = $1
	`
	args := []interface{}{idUsuario}

	if tipo != nil {
		query += " AND tipo_movimentacao = $2"
		args = append(args, *tipo)
	}

	query += " ORDER BY nome"

	rows, err := config.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categorias []schemas.CategoriaResponse
	for rows.Next() {
		var c schemas.CategoriaResponse
		if err := rows.Scan(&c.ID, &c.Nome, &c.TipoMovimentacao); err != nil {
			return nil, err
		}
		categorias = append(categorias, c)
	}

	return categorias, nil
}

func CreateCategoriaRepository(request entities.Categoria) (*entities.Categoria, error) {
	query := `
		INSERT INTO categoria (id_usuario, nome, tipo_movimentacao)
		VALUES ($1, $2, $3)
		RETURNING id_categoria, nome, tipo_movimentacao, id_usuario
	`

	var categoria entities.Categoria
	err := config.DB.QueryRow(
		query,
		request.IDUsuario,
		request.Nome,
		request.TipoMovimentacao,
	).Scan(
		&categoria.ID,
		&categoria.Nome,
		&categoria.TipoMovimentacao,
		&categoria.IDUsuario,
	)

	if err != nil {
		return nil, err
	}

	return &categoria, nil
}

func UpdateCategoriaRepository(request entities.Categoria) (*entities.Categoria, error) {
	query := `
		UPDATE categoria
		SET nome = $1, tipo_movimentacao = $2
		WHERE id_categoria = $3
		AND id_usuario = $4
		RETURNING id_categoria, nome, tipo_movimentacao, id_usuario
	`

	var categoria entities.Categoria
	err := config.DB.QueryRow(
		query,
		request.Nome,
		request.TipoMovimentacao,
		request.ID,
		request.IDUsuario,
	).Scan(
		&categoria.ID,
		&categoria.Nome,
		&categoria.TipoMovimentacao,
		&categoria.IDUsuario,
	)

	if err != nil {
		return nil, err
	}

	return &categoria, nil
}

func DeleteCategoriaRepository(idUsuario int, idCategoria int) error {
	query := `
		DELETE FROM categoria
		WHERE id_categoria = $1
		AND id_usuario = $2
	`

	result, err := config.DB.Exec(query, idCategoria, idUsuario)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("categoria não encontrada ou não pertence ao usuário")
	}

	return nil
}
