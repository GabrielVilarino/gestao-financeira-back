package categoria

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/repository"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	"github.com/GabrielVilarino/gestao-financeira-back.git/utils"
)

func UpdateCategoriaService(idUsuario int, request schemas.UpdateCategoriaRequest) (*schemas.CategoriaResponse, error) {

	// Normaliza o nome removendo acentos e convertendo para maiúsculo
	request.Nome = utils.NormalizarString(request.Nome)

	categoriaEntity := SchemaUpdateToEntity(idUsuario, request)

	categoria, err := repository.UpdateCategoriaRepository(*categoriaEntity)
	if err != nil {
		return nil, err
	}

	reponse := EntityToSchema(*categoria)

	return reponse, nil
}
