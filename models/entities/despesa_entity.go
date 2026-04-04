package entities

type Despesa struct {
	ID                  int     `json:"id_despesa"`
	IDUsuario           int     `json:"id_usuario"`
	IDGrupo             *int    `json:"id_grupo,omitempty"`
	IDCategoria         int     `json:"id_categoria,omitempty"`
	IDSubcategoria      *int    `json:"id_subcategoria,omitempty"`
	TipoTransacao       string  `json:"tipo_transacao"`
	DataPagamento       string  `json:"data_pagamento"`
	DataUltimoPagamento *string `json:"data_ult_pagamento,omitempty"`
	DataCriacao         *string `json:"data_criacao,omitempty"`
	Valor               float64 `json:"valor"`
}
