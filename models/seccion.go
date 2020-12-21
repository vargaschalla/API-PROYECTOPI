package models

import "gorm.io/gorm"

type Seccion struct {
	gorm.Model
	Descripcion     string `json:"descripcion"`
	Ubicacion       string `json:"ubicacion"`
	Estado          string `json:"estado"`
	Sesiones        []Sesiones
	SesionMaterial  []SesionMaterial
	SesionActividad []SesionActividad
	CargaAcademica  []CargaAcademica
}
