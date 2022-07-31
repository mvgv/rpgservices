package models

type Personagem struct {
	Nome   string `json:"nome"`
	Classe string `json:"classe"`
}

var Personagens []Personagem
