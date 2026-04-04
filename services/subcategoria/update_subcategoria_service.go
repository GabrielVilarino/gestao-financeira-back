package subcategoria

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/repository"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	"github.com/GabrielVilarino/gestao-financeira-back.git/utils"
)

func UpdateSubCategoriaService(idUsuario int, request schemas.UpdateSubCategoriaRequest) (*schemas.SubCategoriaResponse, error) {

	// Normaliza o nome removendo acentos e convertendo para maiúsculo
	request.Nome = utils.NormalizarString(request.Nome)

	subcategoriaEntity := SchemaUpdateToEntity(request)

	subcategoria, err := repository.UpdateSubCategoriaRepository(*subcategoriaEntity)
	if err != nil {
		return nil, err
	}

	reponse := EntityToSchema(*subcategoria)

	return reponse, nil
}
