package schemas

type CreateDespesaRequest struct {
	IDGrupo             *int    `json:"id_grupo"`
	IDCategoria         int     `json:"id_categoria" binding:"required"`
	IDSubcategoria      *int    `json:"id_subcategoria"`
	TipoTransacao       string  `json:"tipo_transacao" binding:"required"`
	DataPagamento       string  `json:"data_pagamento" binding:"required"`
	DataUltimoPagamento *string `json:"data_ultimo_pagamento"`
	Valor               float64 `json:"valor" binding:"required"`
}

type CreateDespesaResponse struct {
	ID                  int     `json:"id_despesa"`
	IDUsuario           int     `json:"id_usuario"`
	IDGrupo             *int    `json:"id_grupo,omitempty"`
	IDCategoria         int     `json:"id_categoria"`
	IDSubcategoria      *int    `json:"id_subcategoria,omitempty"`
	TipoTransacao       string  `json:"tipo_transacao"`
	DataPagamento       string  `json:"data_pagamento"`
	DataUltimoPagamento *string `json:"data_ultimo_pagamento,omitempty"`
	DataCriacao         *string `json:"data_criacao"`
	Valor               float64 `json:"valor"`
}

type UpdateDespesaRequest struct {
	ID                  int     `json:"id" binding:"required"`
	IDGrupo             *int    `json:"id_grupo"`
	IDCategoria         int     `json:"id_categoria" binding:"required"`
	IDSubcategoria      *int    `json:"id_subcategoria"`
	TipoTransacao       string  `json:"tipo_transacao" binding:"required"`
	DataPagamento       string  `json:"data_pagamento" binding:"required"`
	DataUltimoPagamento *string `json:"data_ultimo_pagamento"`
	Valor               float64 `json:"valor" binding:"required"`
}

type UpdateDespesaResponse struct {
	ID                  int     `json:"id_despesa"`
	IDUsuario           int     `json:"id_usuario"`
	IDGrupo             *int    `json:"id_grupo,omitempty"`
	IDCategoria         int     `json:"id_categoria"`
	IDSubcategoria      *int    `json:"id_subcategoria,omitempty"`
	TipoTransacao       string  `json:"tipo_transacao"`
	Valor               float64 `json:"valor"`
	DataPagamento       string  `json:"data_pagamento"`
	DataUltimoPagamento *string `json:"data_ult_pagamento,omitempty"`
	DataCriacao         *string `json:"data_criacao,omitempty"`
}

type DeleteDespesaResponse struct {
	Message string `json:"message"`
}

type GetDespesaResponse struct {
	ID                  int     `json:"id"`
	Nome                string  `json:"nome"`
	DataPagamento       string  `json:"data"`
	DataUltimoPagamento *string `json:"data_ult_pagamento,omitempty"`
	Valor               float64 `json:"valor"`
	Categoria           string  `json:"categoria"`
	Subcategoria        *string `json:"subcategoria,omitempty"`
	TipoTransacao       string  `json:"tipo"`
}

type GetDespesaByIDResponse struct {
	IDCategoria         int     `json:"id_categoria"`
	IDSubcategoria      *int    `json:"id_subcategoria"`
	DataPagamento       string  `json:"data_pagamento"`
	DataUltimoPagamento *string `json:"data_ult_pagamento,omitempty"`
	Valor               float64 `json:"valor"`
	TipoTransacao       string  `json:"tipo_transacao"`
}
