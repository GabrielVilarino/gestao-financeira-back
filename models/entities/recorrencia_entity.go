package entities

type Recorrencia struct {
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
	CreatedAt      *string `json:"created_at,omitempty"`
}
