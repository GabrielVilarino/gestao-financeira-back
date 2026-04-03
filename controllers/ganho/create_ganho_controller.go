package ganho

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	ganhoService "github.com/GabrielVilarino/gestao-financeira-back.git/services/ganho"
)

func CreateGanhoController(request schemas.CreateGanhoRequest) (*schemas.CreateGanhoResponse, error) {
	response, err := ganhoService.CreateGanhoService(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
