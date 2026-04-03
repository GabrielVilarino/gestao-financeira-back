package entities

type Ganho struct {
	ID              int     `json:"id_receita"`
	IDUsuario       int     `json:"id_usuario"`
	IDGrupo         *int    `json:"id_grupo,omitempty"`
	IDCategoria     int     `json:"id_categoria,omitempty"`
	IDSubcategoria  *int    `json:"id_subcategoria,omitempty"`
	TipoTransacao   string  `json:"tipo_transacao"`
	Valor           float64 `json:"valor"`
	DataRecebimento string  `json:"data_recebimento"`
	DataCriacao     *string `json:"data_criacao"`
}
