package ganho

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	ganhoService "github.com/GabrielVilarino/gestao-financeira-back.git/services/ganho"
)

func GetGanhosController(idUsuario int, dataInicio, dataFim *string, idGrupo *int) ([]schemas.GetGanhoResponse, error) {
	return ganhoService.GetGanhosService(idUsuario, dataInicio, dataFim, idGrupo)
}

func GetGanhosByIDController(id string) (*schemas.GetGanhoByIDResponse, error) {
	return ganhoService.GetGanhosByIDService(id)
}

func IsValidationError(err error) bool {
	_, ok := err.(*ganhoService.ValidationError)
	return ok
}
