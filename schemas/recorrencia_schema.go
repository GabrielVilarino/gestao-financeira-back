package schemas

type CreateRecorrenciaRequest struct {
	IDGrupo        *int    `json:"id_grupo"`
	IDCategoria    int     `json:"id_categoria" binding:"required"`
	IDSubcategoria *int    `json:"id_subcategoria"`
	Tipo           string  `json:"tipo" binding:"required"`
	Descricao      string  `json:"descricao" binding:"required"`
	Observacao     *string `json:"observacao"`
	Valor          float64 `json:"valor" binding:"required"`
	Frequencia     string  `json:"frequencia" binding:"required"`
	Intervalo      int     `json:"intervalo"`
	DataInicio     string  `json:"data_inicio" binding:"required"`
	DataFim        *string `json:"data_fim"`
}

type CreateRecorrenciaResponse struct {
	ID             int     `json:"id_recorrencia"`
	IDUsuario      int     `json:"id_usuario"`
	IDGrupo        *int    `json:"id_grupo,omitempty"`
	Escopo         string  `json:"escopo"`
	IDCategoria    int     `json:"id_categoria"`
	IDSubcategoria *int    `json:"id_subcategoria,omitempty"`
	Tipo           string  `json:"tipo"`
	Descricao      string  `json:"descricao"`
	Observacao     *string `json:"observacao,omitempty"`
	Valor          float64 `json:"valor"`
	Frequencia     string  `json:"frequencia"`
	Intervalo      int     `json:"intervalo"`
	DataInicio     string  `json:"data_inicio"`
	DataFim        *string `json:"data_fim,omitempty"`
	Ativa          bool    `json:"ativa"`
	CreatedAt      *string `json:"created_at"`
}

type GetRecorrenciaResponse struct {
	ID             int     `json:"id_recorrencia"`
	IDCategoria    int     `json:"id_categoria"`
	IDSubcategoria *int    `json:"id_subcategoria,omitempty"`
	Tipo           string  `json:"tipo"`
	Descricao      string  `json:"descricao"`
	Observacao     *string `json:"observacao,omitempty"`
	Valor          float64 `json:"valor"`
	Frequencia     string  `json:"frequencia"`
	Intervalo      int     `json:"intervalo"`
	DataInicio     string  `json:"data_inicio"`
	DataFim        *string `json:"data_fim,omitempty"`
	Ativa          bool    `json:"ativa"`
	Categoria      string  `json:"categoria"`
	Subcategoria   *string `json:"subcategoria,omitempty"`
}

type EncerrarRecorrenciaResponse struct {
	Message string `json:"message"`
}
