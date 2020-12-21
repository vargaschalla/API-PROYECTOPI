package models

import (
	"gorm.io/gorm"
)

type PlanAcademico struct {
	gorm.Model
	Nombre      string `json:"nombre"`
	Estado      string `json:"estado"`
	SeccionID   string `gorm:"size:191"`
	PeriodoID   string `gorm:"size:191"`
	CursoID     string `gorm:"size:191"`
	PersonaID   string `gorm:"size:191"`
	Unidad      []Unidad
	Inscripcion []Inscripcion
	Personas    []Persona
	Seccion     Seccion
	Periodo     Periodo
	Persona     Persona
}
