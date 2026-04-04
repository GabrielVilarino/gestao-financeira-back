package categoria

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	"github.com/GabrielVilarino/gestao-financeira-back.git/services/categoria"
)

func CreateCategoriaController(idUsuario int, request schemas.CreateCategoriaRequest) (*schemas.CategoriaResponse, error) {
	response, err := categoria.CreateCategoriaService(idUsuario, request)

	if err != nil {
		return nil, err
	}

	return response, nil
}
