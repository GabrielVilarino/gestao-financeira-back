package categoria

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	categoriaService "github.com/GabrielVilarino/gestao-financeira-back.git/services/categoria"
)

func GetCategoriasController() ([]schemas.GetCategoriaResponse, error) {
	return categoriaService.GetCategoriasService()
}

func GetSubcategoriasController(idCategoria *int) ([]schemas.GetSubcategoriaResponse, error) {
	return categoriaService.GetSubcategoriasService(idCategoria)
}
