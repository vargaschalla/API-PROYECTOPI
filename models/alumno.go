package models

import "gorm.io/gorm"

type Alumno struct {
	gorm.Model
	Usuario         string `json:"usuario"`
	Password        string `json:"password"`
	RolID           string `gorm:"size:191"`
	Rol             Rol    `gorm :"ForeignKey: Alumno"`
	Estado          string `json:"estado"`
	SesionActividad []SesionActividad

	CargaAcademica []CargaAcademica
}
