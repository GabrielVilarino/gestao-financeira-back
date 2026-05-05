package entities

type Grupo struct {
	ID            int            `json:"id"`
	Nome          string         `json:"nome"`
	DataCriacao   string         `json:"data_criacao"`
	Participantes []Participante `json:"participantes"`
}

type Participante struct {
	ID          int    `json:"id_participante"`
	Nome        string `json:"nome"`
	Email       string `json:"email"`
	DataCriacao string `json:"data_criacao"`
	IsAdmin     bool   `json:"is_admin"`
}
