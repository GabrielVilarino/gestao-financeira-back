package categoria

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	"github.com/GabrielVilarino/gestao-financeira-back.git/services/categoria"
)

func UpdateCategoriaController(idUsuario int, request schemas.UpdateCategoriaRequest) (*schemas.CategoriaResponse, error) {
	response, err := categoria.UpdateCategoriaService(idUsuario, request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
