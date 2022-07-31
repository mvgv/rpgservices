package models

type Personagem struct {
	Id     int    `json:"id"`
	Nome   string `json:"nome"`
	Classe string `json:"classe"`
}

var Personagens []Personagem
