package repository

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/config"
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/entities"
)

func CreateGrupoRepository(idUsuario int, grupoEntity entities.Grupo) (*entities.Grupo, error) {
	queryAddGrupo := `
		INSERT INTO grupo (nome)
		VALUES ($1)
		RETURNING id_grupo, nome, data_criacao
	`

	err := config.DB.QueryRow(
		queryAddGrupo,
		grupoEntity.Nome,
	).Scan(&grupoEntity.ID, &grupoEntity.Nome, &grupoEntity.DataCriacao)

	if err != nil {
		return nil, err
	}

	queryUpdateUser := `
		UPDATE usuario
		SET 
			id_grupo = $1,
			is_admin = true
		WHERE 
			id_usuario = $2
	`
	_, err = config.DB.Exec(
		queryUpdateUser,
		grupoEntity.ID,
		idUsuario,
	)

	if err != nil {
		return nil, err
	}
	return &grupoEntity, nil
}

func GetGrupoByIDRepository(idGrupo int) (*entities.Grupo, error) {
	query := `
		SELECT id_grupo, nome, data_criacao
		FROM grupo
		WHERE id_grupo = $1
	`

	var grupo entities.Grupo
	err := config.DB.QueryRow(query, idGrupo).Scan(&grupo.ID, &grupo.Nome, &grupo.DataCriacao)

	if err != nil {
		return nil, err
	}

	return &grupo, nil
}
