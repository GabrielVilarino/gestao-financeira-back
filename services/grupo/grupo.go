package grupo

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/entities"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
)

func CreateGrupoSchemaToEntity(request schemas.CreateGrupoRequest) entities.Grupo {
	return entities.Grupo{
		Nome: request.Nome,
	}
}

func EntityToGrupoResponse(grupoEntity entities.Grupo) schemas.GrupoResponse {
	return schemas.GrupoResponse{
		ID:          grupoEntity.ID,
		Nome:        grupoEntity.Nome,
		DataCriacao: grupoEntity.DataCriacao,
	}
}

func EntityToParticipanteResponse(userEntity []entities.User) schemas.ParticipanteResponse {
	participantes := make([]schemas.Participante, len(userEntity))
	for i, user := range userEntity {
		participantes[i] = schemas.Participante{
			ID:          *user.ID,
			Nome:        user.Nome,
			Email:       user.Email,
			IsAdmin:     *user.IsAdmin,
			DataCriacao: *user.DataCriacao,
		}
	}
	return schemas.ParticipanteResponse{Participantes: participantes}
}
