package model

type Procuct struct {
	ID          int     `json:"id"`
	Name        string  `json:"nome"`
	Price       float64 `json:"preco"`
	Description string  `json:"descricao"`
}
