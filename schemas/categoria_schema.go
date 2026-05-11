package schemas

type CreateCategoriaRequest struct {
	Nome             string `json:"nome" binding:"required"`
	TipoMovimentacao string `json:"tipo_movimentacao" binding:"required,oneof=receita despesa"`
}

type UpdateCategoriaRequest struct {
	ID               int    `json:"id_categoria" binding:"required"`
	Nome             string `json:"nome" binding:"required"`
	TipoMovimentacao string `json:"tipo_movimentacao" binding:"required,oneof=receita despesa"`
}

type CategoriaResponse struct {
	ID               int    `json:"id_categoria"`
	Nome             string `json:"nome"`
	TipoMovimentacao string `json:"tipo_movimentacao"`
	IDUsuario        *int   `json:"id_usuario"`
	IDGrupo          *int   `json:"id_grupo,omitempty"`
}
