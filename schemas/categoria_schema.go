package schemas

type GetCategoriaResponse struct {
	ID               int    `json:"id_categoria"`
	Nome             string `json:"nome"`
	TipoMovimentacao string `json:"tipo_movimentacao"`
}

type GetSubcategoriaResponse struct {
	ID          int    `json:"id_subcategoria"`
	IDCategoria int    `json:"id_categoria"`
	Nome        string `json:"nome"`
}
