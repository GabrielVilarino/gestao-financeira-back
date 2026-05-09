package entities

type Transacao struct {
	ID             int     `json:"id_transacao"`
	IDUsuario      int     `json:"id_usuario"`
	IDGrupo        *int    `json:"id_grupo,omitempty"`
	Escopo         string  `json:"escopo"`
	IDCategoria    int     `json:"id_categoria"`
	IDSubcategoria *int    `json:"id_subcategoria,omitempty"`
	IDRecorrencia  *int    `json:"id_recorrencia,omitempty"`
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
