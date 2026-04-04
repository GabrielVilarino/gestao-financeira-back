package entities

type Categoria struct {
	ID               int    `json:"id_categoria"`
	Nome             string `json:"nome"`
	TipoMovimentacao string `json:"tipo_movimentacao"`
	IDUsuario        int    `json:"id_usuario"`
}
