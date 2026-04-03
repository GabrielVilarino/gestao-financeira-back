package ganho

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	ganhoService "github.com/GabrielVilarino/gestao-financeira-back.git/services/ganho"
)

func GetGanhosController(dataInicio, dataFim *string, idGrupo *int) ([]schemas.GetGanhoResponse, error) {
	return ganhoService.GetGanhosService(dataInicio, dataFim, idGrupo)
}

func IsValidationError(err error) bool {
	_, ok := err.(*ganhoService.ValidationError)
	return ok
}
