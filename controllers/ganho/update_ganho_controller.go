package ganho

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	ganhoService "github.com/GabrielVilarino/gestao-financeira-back.git/services/ganho"
)

func UpdateGanhoController(request schemas.UpdateGanhoRequest) (*schemas.UpdateGanhoResponse, error) {
	response, err := ganhoService.UpdateGanhoService(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
