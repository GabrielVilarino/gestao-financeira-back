package transacao

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/repository"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	"github.com/GabrielVilarino/gestao-financeira-back.git/services"
)

func UpdateTransacaoService(idUsuario int, idGrupoJWT *int, request schemas.UpdateTransacaoRequest) (*schemas.UpdateTransacaoResponse, error) {
	if err := services.ValidarIDGrupo(request.IDGrupo, idGrupoJWT); err != nil {
		return nil, err
	}

	transacao := transacaoUpdateSchemaToEntity(request)
	transacao.IDUsuario = idUsuario

	if request.IDGrupo != nil {
		transacao.IDGrupo = request.IDGrupo
		transacao.Escopo = "GRUPO"
	} else {
		transacao.Escopo = "PESSOAL"
	}

	updated, err := repository.UpdateTransacaoRepository(transacao)
	if err != nil {
		return nil, err
	}

	response := transacaoEntityToUpdateSchema(*updated)
	return &response, nil
}
