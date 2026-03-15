package schemas

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
