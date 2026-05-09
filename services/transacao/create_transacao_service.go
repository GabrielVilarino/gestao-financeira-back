package transacao

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/repository"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	"github.com/GabrielVilarino/gestao-financeira-back.git/services"
)

func CreateTransacaoService(idUsuario int, idGrupoJWT *int, request schemas.CreateTransacaoRequest) (*schemas.CreateTransacaoResponse, error) {
	if err := services.ValidarIDGrupo(request.IDGrupo, idGrupoJWT); err != nil {
		return nil, err
	}

	transacao := transacaoSchemaToEntity(request)
	transacao.IDUsuario = idUsuario

	if request.IDGrupo != nil {
		transacao.IDGrupo = request.IDGrupo
		transacao.Escopo = "GRUPO"
	} else if idGrupoJWT != nil {
		transacao.IDGrupo = idGrupoJWT
		transacao.Escopo = "GRUPO"
	} else {
		transacao.Escopo = "PESSOAL"
	}

	created, err := repository.CreateTransacaoRepository(transacao)
	if err != nil {
		return nil, err
	}

	response := transacaoEntityToCreateSchema(*created)
	return &response, nil
}
