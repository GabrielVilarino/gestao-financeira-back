package transacao

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	transacaoService "github.com/GabrielVilarino/gestao-financeira-back.git/services/transacao"
)

func CreateTransacaoController(idUsuario int, idGrupoJWT *int, request schemas.CreateTransacaoRequest) (*schemas.CreateTransacaoResponse, error) {
	return transacaoService.CreateTransacaoService(idUsuario, idGrupoJWT, request)
}
