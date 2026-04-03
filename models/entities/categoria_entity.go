package entities

type Categoria struct {
	ID               int    `json:"id_categoria"`
	Nome             string `json:"nome"`
	TipoMovimentacao string `json:"tipo_movimentacao"`
}

type Subcategoria struct {
	ID          int    `json:"id_subcategoria"`
	IDCategoria int    `json:"id_categoria"`
	Nome        string `json:"nome"`
}
