package models

import "gorm.io/gorm"

type Curso struct {
	gorm.Model
	Descripcion      string `json:"descripcion"`
	Estado           string `json:"estado"`
	Academic_Plane   []Academic_Plane
	CategoriaCursoID string         `gorm:"size:191"`
	CategoriaCurso   CategoriaCurso `gorm :"ForeignKey: Curso"`
}
