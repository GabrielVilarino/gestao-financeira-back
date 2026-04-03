package ganho

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/repository"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
)

func UpdateGanhoService(request schemas.UpdateGanhoRequest) (*schemas.UpdateGanhoResponse, error) {
	ganho := GanhoUpdateSchemaToEntity(request)

	ganhoUpdated, err := repository.UpdateGanhoRepository(ganho)
	if err != nil {
		return nil, err
	}

	response := GanhoEntityToUpdateSchema(*ganhoUpdated)
	return &response, nil
}
