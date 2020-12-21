package models

import (
	"gorm.io/gorm"
)

type Academic_Period struct {
	gorm.Model
	Descripcion     string `json:"descripcion"`
	Estado          string `json:"estado"`
	Sesiones        []Sesiones
	SesionMaterial  []SesionMaterial
	Academic_Plane  []Academic_Plane
	Modulo          []Modulo
	SesionActividad []SesionActividad
	CargaAcademica  []CargaAcademica
}