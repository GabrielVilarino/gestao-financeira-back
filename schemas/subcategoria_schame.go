package schemas

type CreateSubCategoriaRequest struct {
	IDCategoria int    `json:"id_categoria" binding:"required"`
	Nome        string `json:"nome" binding:"required"`
}

type UpdateSubCategoriaRequest struct {
	ID   int    `json:"id_subcategoria" binding:"required"`
	Nome string `json:"nome" binding:"required"`
}

type SubCategoriaResponse struct {
	ID          int    `json:"id_subcategoria"`
	IDCategoria int    `json:"id_categoria"`
	Nome        string `json:"nome"`
}
