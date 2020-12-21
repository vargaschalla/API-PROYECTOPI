package models

import "gorm.io/gorm"

type Inscripcion struct {
	gorm.Model
	PlanAcademicoID string `gorm:"size:191"`
	PlanAcademico   PlanAcademico
	PersonaID       string `gorm:"size:191"`
	Persona         Persona
}
