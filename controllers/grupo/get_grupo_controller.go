package grupo

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	"github.com/GabrielVilarino/gestao-financeira-back.git/services/grupo"
)

func GetGrupoByIDController(idGrupo int) (*schemas.GrupoResponse, error) {
	grupo, err := grupo.GetGrupoServiceByID(idGrupo)

	if err != nil {
		return nil, err
	}

	return grupo, nil
}
