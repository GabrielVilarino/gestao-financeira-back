package subcategoria

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/repository"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	"github.com/GabrielVilarino/gestao-financeira-back.git/utils"
)

func CreateSubCategoriaService(idUsuario int, idGrupo *int, request schemas.CreateSubCategoriaRequest) (*schemas.SubCategoriaResponse, error) {

	// Normaliza o nome removendo acentos e convertendo para maiúsculo
	request.Nome = utils.NormalizarString(request.Nome)

	subcategoriaEntity := SchemaCreateToEntity(idGrupo, request)

	subcategoria, err := repository.CreateSubcategoriaRepository(*subcategoriaEntity)
	if err != nil {
		return nil, err
	}

	response := EntityToSchema(*subcategoria)

	return response, nil
}
