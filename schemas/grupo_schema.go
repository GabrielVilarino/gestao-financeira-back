package schemas

type CreateGrupoRequest struct {
	Nome string `json:"nome" binding:"required"`
}

type GrupoResponse struct {
	ID          int    `json:"id"`
	Nome        string `json:"nome"`
	DataCriacao string `json:"data_criacao"`
}
