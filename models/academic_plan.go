package models

import (
	"gorm.io/gorm"
)

type Academic_Plane struct {
	gorm.Model
	Descripcion       string `json:"descripcion"`
	Fecha             string `json:"fecha"`
	Estado            string `json:"estado"`
	Sesiones          []Sesiones
	SesionMaterial    []SesionMaterial
	Modulo            []Modulo
	SesionActividad   []SesionActividad
	CargaAcademica    []CargaAcademica
	GradoID           string          `gorm:"size:191"`
	Grado             Grado           `gorm :"ForeignKey: Academic_Plane"`
	Academic_PeriodID string          `gorm:"size:191"`
	Academic_Period   Academic_Period `gorm :"ForeignKey: Academic_Plane"`
	CursoID           string          `gorm:"size:191"`
	Curso             Curso           `gorm :"ForeignKey: Academic_Plane"`
}
