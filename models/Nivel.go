package models

import "gorm.io/gorm"

type Nivel struct {
	gorm.Model
	Descripcion string `json:"descripcion"`
	Grado       []Grado
	Estado      string `json:"estado"`
}
