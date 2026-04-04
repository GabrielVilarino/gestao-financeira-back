package subcategoria

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	"github.com/GabrielVilarino/gestao-financeira-back.git/services/subcategoria"
)

func UpdateSubCategoriaController(idUsuario int, request schemas.UpdateSubCategoriaRequest) (*schemas.SubCategoriaResponse, error) {
	response, err := subcategoria.UpdateSubCategoriaService(idUsuario, request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
