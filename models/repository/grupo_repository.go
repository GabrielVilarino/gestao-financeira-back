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

func GetParticipantesByIDRepository(idGrupo int) (*[]entities.User, error) {
	query := `
		SELECT 
			u.id_usuario,
			u.nome,
			u.email,
			u.data_criacao,
			u.is_admin
		FROM 
			usuario u
		WHERE 
			u.id_grupo = $1
	`

	rows, err := config.DB.Query(query, idGrupo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var participantes []entities.User

	for rows.Next() {
		var participante entities.Participante

		err := rows.Scan(
			&participante.ID,
			&participante.Nome,
			&participante.Email,
			&participante.DataCriacao,
			&participante.IsAdmin,
		)
		if err != nil {
			return nil, err
		}

		participantes = append(participantes, entities.User{
			ID:          &participante.ID,
			Nome:        participante.Nome,
			Email:       participante.Email,
			DataCriacao: &participante.DataCriacao,
			IsAdmin:     &participante.IsAdmin,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &participantes, nil
}

func AddParticipanteRepository(email string, idGroup int) error {
	query := `
		UPDATE usuario
		SET id_grupo = $1
		WHERE email = $2
	`

	_, err := config.DB.Exec(query, idGroup, email)
	if err != nil {
		return err
	}

	return nil
}

func DeleteParticipanteRepository(idParticipante int, idGroup int) error {
	query := `
		UPDATE usuario
		SET id_grupo = NULL, is_admin = false
		WHERE id_usuario = $1 AND id_grupo = $2
	`

	_, err := config.DB.Exec(query, idParticipante, idGroup)
	if err != nil {
		return err
	}

	return nil
}

func UpdateRoleParticipanteRepository(isAdmin bool, idParticipante int, idGroup int) error {
	query := `
		UPDATE usuario
		SET is_admin = $1
		WHERE id_usuario = $2 AND id_grupo = $3
	`

	_, err := config.DB.Exec(query, isAdmin, idParticipante, idGroup)
	if err != nil {
		return err
	}

	return nil
}
