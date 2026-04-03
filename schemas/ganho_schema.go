package schemas

type CreateGanhoRequest struct {
	IDGrupo         *int    `json:"id_grupo"`
	IDCategoria     int     `json:"id_categoria" binding:"required"`
	IDSubcategoria  *int    `json:"id_subcategoria"`
	TipoTransacao   string  `json:"tipo_transacao" binding:"required"`
	Valor           float64 `json:"valor" binding:"required"`
	DataRecebimento string  `json:"data_recebimento" binding:"required"`
}

type CreateGanhoResponse struct {
	ID              int     `json:"id_receita"`
	IDUsuario       int     `json:"id_usuario"`
	IDGrupo         *int    `json:"id_grupo,omitempty"`
	IDCategoria     int     `json:"id_categoria"`
	IDSubcategoria  *int    `json:"id_subcategoria,omitempty"`
	TipoTransacao   string  `json:"tipo_transacao"`
	Valor           float64 `json:"valor"`
	DataRecebimento string  `json:"data_recebimento"`
	DataCriacao     *string `json:"data_criacao"`
}

type UpdateGanhoRequest struct {
	ID              int     `json:"id" binding:"required"`
	IDGrupo         *int    `json:"id_grupo"`
	IDCategoria     int     `json:"id_categoria" binding:"required"`
	IDSubcategoria  *int    `json:"id_subcategoria"`
	TipoTransacao   string  `json:"tipo_transacao" binding:"required"`
	Valor           float64 `json:"valor" binding:"required"`
	DataRecebimento string  `json:"data_recebimento" binding:"required"`
}

type UpdateGanhoResponse struct {
	ID              int     `json:"id_receita"`
	IDUsuario       int     `json:"id_usuario"`
	IDGrupo         *int    `json:"id_grupo,omitempty"`
	IDCategoria     int     `json:"id_categoria"`
	IDSubcategoria  *int    `json:"id_subcategoria,omitempty"`
	TipoTransacao   string  `json:"tipo_transacao"`
	Valor           float64 `json:"valor"`
	DataRecebimento string  `json:"data_recebimento"`
	DataCriacao     *string `json:"data_criacao"`
}

type GetGanhoResponse struct {
	ID              int     `json:"id"`
	Nome            string  `json:"nome"`
	DataRecebimento string  `json:"data"`
	Valor           float64 `json:"valor"`
	TipoTransacao   string  `json:"tipo"`
}

type GetGanhoByIDResponse struct {
	IDCategoria     int     `json:"id_categoria"`
	IDSubcategoria  *int    `json:"id_subcategoria"`
	Valor           float64 `json:"valor"`
	TipoTransacao   string  `json:"tipo_transacao"`
	DataRecebimento string  `json:"data_recebimento"`
}

type DeleteGanhoResponse struct {
	Message string `json:"message"`
}
