package schemas

type CreateGrupoRequest struct {
	Nome string `json:"nome" binding:"required"`
}

type GrupoResponse struct {
	ID          int    `json:"id"`
	Nome        string `json:"nome"`
	DataCriacao string `json:"data_criacao"`
}

type AddParticipanteRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type UpdateRoleParticipanteRequest struct {
	IsAdmin bool `json:"is_admin"`
}
