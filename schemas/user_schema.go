package schemas

type AuthUserRequest struct {
	Email string `json:"email" binding:"required,email"`
	Senha string `json:"senha" binding:"required"`
}

type AuthUserResponse struct {
	ID      int    `json:"id_usuario"`
	Nome    string `json:"nome"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin"`
	IdGroup *int   `json:"id_grupo"`
}

type CreateUserRequest struct {
	Nome  string `json:"nome" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Senha string `json:"senha" binding:"required,min=6"`
}

type CreateUserResponse struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}
