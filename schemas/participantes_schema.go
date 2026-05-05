package schemas

type Participante struct {
	ID          int    `json:"id"`
	Nome        string `json:"nome"`
	Email       string `json:"email"`
	DataCriacao string `json:"data_criacao"`
	IsAdmin     bool   `json:"is_admin"`
}
type ParticipanteResponse struct {
	Participantes []Participante `json:"participantes"`
}
