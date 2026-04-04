package categoria

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	categoriaService "github.com/GabrielVilarino/gestao-financeira-back.git/services/categoria"
)

func GetCategoriasController(tipo *string) ([]schemas.GetCategoriaResponse, error) {
	return categoriaService.GetCategoriasService(tipo)
}

func GetSubcategoriasController(idCategoria *int) ([]schemas.GetSubcategoriaResponse, error) {
	return categoriaService.GetSubcategoriasService(idCategoria)
}
