package entities

type SubCategoria struct {
	ID          int    `json:"id_subcategoria"`
	IDCategoria int    `json:"id_categoria"`
	Nome        string `json:"nome"`
	IDGrupo     *int   `json:"id_grupo,omitempty"`
}
