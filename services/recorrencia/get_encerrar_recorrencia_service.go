package recorrencia

import (
	"fmt"
	"time"

	"github.com/GabrielVilarino/gestao-financeira-back.git/models/repository"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
)

func GetRecorrenciasService(idUsuario int, idGrupoJWT *int, idGrupo *int) ([]schemas.GetRecorrenciaResponse, error) {
	idGrupoResolved := idGrupo
	if idGrupoResolved == nil {
		idGrupoResolved = idGrupoJWT
	}

	recorrencias, err := repository.GetRecorrenciasByUsuarioRepository(idUsuario, idGrupoResolved)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar recorrências: %w", err)
	}

	return recorrencias, nil
}

func EncerrarRecorrenciaService(id int) error {
	dataFim := time.Now().Format("2006-01-02")

	if err := repository.EncerrarRecorrenciaRepository(id, dataFim); err != nil {
		return fmt.Errorf("erro ao encerrar recorrência: %w", err)
	}

	if err := repository.DeleteTransacoesFuturasByRecorrenciaRepository(id, dataFim); err != nil {
		return fmt.Errorf("erro ao remover transações futuras: %w", err)
	}

	return nil
}
