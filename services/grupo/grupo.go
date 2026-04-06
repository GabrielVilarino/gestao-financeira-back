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
