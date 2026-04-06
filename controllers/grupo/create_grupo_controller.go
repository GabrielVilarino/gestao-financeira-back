package grupo

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	"github.com/GabrielVilarino/gestao-financeira-back.git/services/grupo"
)

func CreateGrupoController(idUsuario int, request schemas.CreateGrupoRequest) (*schemas.GrupoResponse, error) {
	response, err := grupo.CreateGrupoService(idUsuario, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
