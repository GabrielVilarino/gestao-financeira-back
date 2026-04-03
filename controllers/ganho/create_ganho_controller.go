package ganho

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	ganhoService "github.com/GabrielVilarino/gestao-financeira-back.git/services/ganho"
)

func CreateGanhoController(idUsuario int, idGrupoJWT *int, request schemas.CreateGanhoRequest) (*schemas.CreateGanhoResponse, error) {
	response, err := ganhoService.CreateGanhoService(idUsuario, idGrupoJWT, request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
