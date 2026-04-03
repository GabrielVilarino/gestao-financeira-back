package ganho

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/repository"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
)

func CreateGanhoService(request schemas.CreateGanhoRequest) (*schemas.CreateGanhoResponse, error) {
	ganho := GanhoSchemaToEntity(request)

	ganhoCreated, err := repository.CreateGanhoRepository(ganho)
	if err != nil {
		return nil, err
	}

	response := GanhoEntityToCreateSchema(*ganhoCreated)
	return &response, nil
}
