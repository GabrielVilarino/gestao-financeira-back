package categoria

import "github.com/GabrielVilarino/gestao-financeira-back.git/models/repository"

func DeleteCategoriaService(idUsuario int, idCategoria int) error {
	return repository.DeleteCategoriaRepository(idUsuario, idCategoria)
}
