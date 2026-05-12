package categoria

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	categoriaService "github.com/GabrielVilarino/gestao-financeira-back.git/services/categoria"
)

func GetCategoriasController(idUsuario int, tipo *string, idGrupo *int) ([]schemas.CategoriaResponse, error) {
	return categoriaService.GetCategoriasService(idUsuario, tipo, idGrupo)
}
