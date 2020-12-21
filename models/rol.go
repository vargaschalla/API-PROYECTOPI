package models

import (
	"gorm.io/gorm"
)

type Rol struct {
	gorm.Model
	Descripcion string `json:"descripcion"`
	Alumno      []Alumno
	Docente     []Docente
	Estado      string `json:"estado"`
}
