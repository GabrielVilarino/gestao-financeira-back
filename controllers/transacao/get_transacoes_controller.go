package transacao

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	transacaoService "github.com/GabrielVilarino/gestao-financeira-back.git/services/transacao"
)

func GetTransacoesController(idUsuario int, dataInicio, dataFim *string, tipo *string, idGrupo *int) ([]schemas.GetTransacaoResponse, error) {
	return transacaoService.GetTransacoesService(idUsuario, dataInicio, dataFim, tipo, idGrupo)
}

func GetTransacaoByIDController(id string) (*schemas.GetTransacaoByIDResponse, error) {
	return transacaoService.GetTransacaoByIDService(id)
}
