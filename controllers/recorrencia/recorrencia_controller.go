package recorrencia

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	recorrenciaService "github.com/GabrielVilarino/gestao-financeira-back.git/services/recorrencia"
)

func CreateRecorrenciaController(idUsuario int, idGrupoJWT *int, req schemas.CreateRecorrenciaRequest) (*schemas.CreateRecorrenciaResponse, error) {
	return recorrenciaService.CreateRecorrenciaService(idUsuario, idGrupoJWT, req)
}

func GetRecorrenciasController(idUsuario int, idGrupoJWT *int, idGrupo *int) ([]schemas.GetRecorrenciaResponse, error) {
	return recorrenciaService.GetRecorrenciasService(idUsuario, idGrupoJWT, idGrupo)
}

func EncerrarRecorrenciaController(id int) error {
	return recorrenciaService.EncerrarRecorrenciaService(id)
}
