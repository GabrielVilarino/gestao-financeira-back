package repository

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/config"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
)

func GetCategoriasRepository(tipo *string) ([]schemas.GetCategoriaResponse, error) {
	query := `
		SELECT id_categoria, nome, tipo_movimentacao
		FROM categoria
	`
	args := []interface{}{}

	if tipo != nil {
		query += " WHERE tipo_movimentacao = $1"
		args = append(args, *tipo)
	}

	query += " ORDER BY nome"

	rows, err := config.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categorias []schemas.GetCategoriaResponse
	for rows.Next() {
		var c schemas.GetCategoriaResponse
		if err := rows.Scan(&c.ID, &c.Nome, &c.TipoMovimentacao); err != nil {
			return nil, err
		}
		categorias = append(categorias, c)
	}

	return categorias, nil
}

func GetSubcategoriasRepository(idCategoria *int) ([]schemas.GetSubcategoriaResponse, error) {
	query := `
		SELECT id_subcategoria, id_categoria, nome
		FROM subcategoria
	`
	args := []interface{}{}

	if idCategoria != nil {
		query += " WHERE id_categoria = $1"
		args = append(args, *idCategoria)
	}

	query += " ORDER BY nome"

	rows, err := config.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subcategorias []schemas.GetSubcategoriaResponse
	for rows.Next() {
		var s schemas.GetSubcategoriaResponse
		if err := rows.Scan(&s.ID, &s.IDCategoria, &s.Nome); err != nil {
			return nil, err
		}
		subcategorias = append(subcategorias, s)
	}

	return subcategorias, nil
}
