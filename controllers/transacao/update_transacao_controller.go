package transacao

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	transacaoService "github.com/GabrielVilarino/gestao-financeira-back.git/services/transacao"
)

func UpdateTransacaoController(idUsuario int, idGrupoJWT *int, request schemas.UpdateTransacaoRequest) (*schemas.UpdateTransacaoResponse, error) {
	return transacaoService.UpdateTransacaoService(idUsuario, idGrupoJWT, request)
}
