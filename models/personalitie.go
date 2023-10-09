package models

type Personalitie struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	BackGround string `json:"background"`
}

var Personalidades []Personalitie
