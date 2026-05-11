package subcategoria

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	"github.com/GabrielVilarino/gestao-financeira-back.git/services/subcategoria"
)

func CreateSubCategoriaController(idUsuario int, idGrupo *int, request schemas.CreateSubCategoriaRequest) (*schemas.SubCategoriaResponse, error) {
	response, err := subcategoria.CreateSubCategoriaService(idUsuario, idGrupo, request)

	if err != nil {
		return nil, err
	}

	return response, nil
}
