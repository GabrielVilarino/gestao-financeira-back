package entities

type User struct {
	ID          *int    `json:"id_usuario"`
	Name        string  `json:"nome"`
	Email       string  `json:"email"`
	Password    string  `json:"senha"`
	IsAdmin     *bool   `json:"is_admin"`
	IdGroup     *int    `json:"id_grupo"`
	DataCriacao *string `json:"data_criacao"`
}
