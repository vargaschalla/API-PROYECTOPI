package models

import "gorm.io/gorm"

type MaterialActividad struct {
	gorm.Model
	Descripcion     string `json:"descripcion"`
	Tipo            string `json:"tipo"`
	Estado          string `json:"estado"`
	SesionMaterial  []SesionMaterial
	SesionActividad []SesionActividad
}
