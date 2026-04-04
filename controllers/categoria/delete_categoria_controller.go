package categoria

import "github.com/GabrielVilarino/gestao-financeira-back.git/services/categoria"

func DeleteCategoriaController(idUsuario int, idCategoria int) error {
	return categoria.DeleteCategoriaService(idUsuario, idCategoria)
}
