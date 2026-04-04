package subcategoria

import "github.com/GabrielVilarino/gestao-financeira-back.git/models/repository"

func DeleteSubCategoriaService(idUsuario int, idSubCategoria int) error {
	return repository.DeleteSubCategoriaRepository(idSubCategoria)
}
