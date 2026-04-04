package subcategoria

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	"github.com/GabrielVilarino/gestao-financeira-back.git/services/subcategoria"
)

func GetSubcategoriasController(idCategoria *int) ([]schemas.SubCategoriaResponse, error) {
	return subcategoria.GetSubcategoriasService(idCategoria)
}
