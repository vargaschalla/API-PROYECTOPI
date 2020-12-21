package models

import "gorm.io/gorm"

type CategoriaCurso struct {
	gorm.Model
	Descripcion string `json:"descripcion"`
	Estado      string `json:"estado"`
	Curso       []Curso
}
