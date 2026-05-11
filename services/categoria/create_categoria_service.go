package categoria

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/repository"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	"github.com/GabrielVilarino/gestao-financeira-back.git/utils"
)

func CreateCategoriaService(idUsuario int, idGrupo *int, request schemas.CreateCategoriaRequest) (*schemas.CategoriaResponse, error) {

	// Normaliza o nome removendo acentos e convertendo para maiúsculo
	request.Nome = utils.NormalizarString(request.Nome)

	categoriaEntity := SchemaCreateToEntity(idUsuario, idGrupo, request)

	categoria, err := repository.CreateCategoriaRepository(*categoriaEntity)
	if err != nil {
		return nil, err
	}

	response := EntityToSchema(*categoria)

	return response, nil
}
