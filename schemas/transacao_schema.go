package schemas

type CreateTransacaoRequest struct {
	IDGrupo        *int    `json:"id_grupo"`
	IDCategoria    int     `json:"id_categoria" binding:"required"`
	IDSubcategoria *int    `json:"id_subcategoria"`
	Tipo           string  `json:"tipo" binding:"required"`
	Descricao      string  `json:"descricao" binding:"required"`
	Observacao     *string `json:"observacao"`
	Valor          float64 `json:"valor" binding:"required"`
	Competencia    string  `json:"competencia" binding:"required"`
	DataVencimento *string `json:"data_vencimento"`
	DataPagamento  *string `json:"data_pagamento"`
	Status         string  `json:"status"`
}

type CreateTransacaoResponse struct {
	ID             int     `json:"id_transacao"`
	IDUsuario      int     `json:"id_usuario"`
	IDGrupo        *int    `json:"id_grupo,omitempty"`
	Escopo         string  `json:"escopo"`
	IDCategoria    int     `json:"id_categoria"`
	IDSubcategoria *int    `json:"id_subcategoria,omitempty"`
	Tipo           string  `json:"tipo"`
	Descricao      string  `json:"descricao"`
	Observacao     *string `json:"observacao,omitempty"`
	Valor          float64 `json:"valor"`
	Competencia    string  `json:"competencia"`
	DataVencimento *string `json:"data_vencimento,omitempty"`
	DataPagamento  *string `json:"data_pagamento,omitempty"`
	Status         string  `json:"status"`
	CreatedAt      *string `json:"created_at"`
}

type UpdateTransacaoRequest struct {
	ID             int     `json:"id" binding:"required"`
	IDGrupo        *int    `json:"id_grupo"`
	IDCategoria    int     `json:"id_categoria" binding:"required"`
	IDSubcategoria *int    `json:"id_subcategoria"`
	Tipo           string  `json:"tipo" binding:"required"`
	Descricao      string  `json:"descricao" binding:"required"`
	Observacao     *string `json:"observacao"`
	Valor          float64 `json:"valor" binding:"required"`
	Competencia    string  `json:"competencia" binding:"required"`
	DataVencimento *string `json:"data_vencimento"`
	DataPagamento  *string `json:"data_pagamento"`
	Status         string  `json:"status"`
}

type UpdateTransacaoResponse struct {
	ID             int     `json:"id_transacao"`
	IDUsuario      int     `json:"id_usuario"`
	IDGrupo        *int    `json:"id_grupo,omitempty"`
	Escopo         string  `json:"escopo"`
	IDCategoria    int     `json:"id_categoria"`
	IDSubcategoria *int    `json:"id_subcategoria,omitempty"`
	Tipo           string  `json:"tipo"`
	Descricao      string  `json:"descricao"`
	Observacao     *string `json:"observacao,omitempty"`
	Valor          float64 `json:"valor"`
	Competencia    string  `json:"competencia"`
	DataVencimento *string `json:"data_vencimento,omitempty"`
	DataPagamento  *string `json:"data_pagamento,omitempty"`
	Status         string  `json:"status"`
	CreatedAt      *string `json:"created_at"`
	UpdatedAt      *string `json:"updated_at"`
}

type GetTransacaoResponse struct {
	ID           int     `json:"id"`
	Nome         string  `json:"nome"`
	Tipo         string  `json:"tipo"`
	Descricao    string  `json:"descricao"`
	Valor        float64 `json:"valor"`
	Competencia  string  `json:"competencia"`
	Status       string  `json:"status"`
	Categoria    string  `json:"categoria"`
	Subcategoria *string `json:"subcategoria,omitempty"`
}

type GetTransacaoByIDResponse struct {
	IDCategoria    int     `json:"id_categoria"`
	IDSubcategoria *int    `json:"id_subcategoria,omitempty"`
	Tipo           string  `json:"tipo"`
	Descricao      string  `json:"descricao"`
	Observacao     *string `json:"observacao,omitempty"`
	Valor          float64 `json:"valor"`
	Competencia    string  `json:"competencia"`
	DataVencimento *string `json:"data_vencimento,omitempty"`
	DataPagamento  *string `json:"data_pagamento,omitempty"`
	Status         string  `json:"status"`
}

type DeleteTransacaoResponse struct {
	Message string `json:"message"`
}
