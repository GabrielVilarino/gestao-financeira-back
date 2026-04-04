package subcategoria

import "github.com/GabrielVilarino/gestao-financeira-back.git/services/subcategoria"

func DeleteSubCategoriaController(idUsuario int, idSubCategoria int) error {
	return subcategoria.DeleteSubCategoriaService(idUsuario, idSubCategoria)
}
