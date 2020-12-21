package models

import "gorm.io/gorm"

type SesionMaterial struct {
	gorm.Model
	Titulo              string            `json:"titulo"`
	Descripcion         string            `json:"descripcion"`
	FechaRegistro       string            `json:"fecharegistro"`
	Ruta                string            `json:"ruta"`
	Estado              string            `json:"estado"`
	DocenteID           string            `gorm:"size:191"`
	Docente             Docente           `gorm :"ForeignKey: Sesiones"`
	SeccionID           string            `gorm:"size:191"`
	Seccion             Seccion           `gorm :"ForeignKey: Sesiones"`
	SesionesID          string            `gorm:"size:191"`
	Sesiones            Sesiones          `gorm :"ForeignKey: Sesiones"`
	MaterialActividadID string            `gorm:"size:191"`
	MaterialActividad   MaterialActividad `gorm :"ForeignKey: Sesiones"`
	Academic_PlaneID    string            `gorm:"size:191"`
	Academic_Plane      Academic_Plane    `gorm :"ForeignKey: Sesiones"`
	Academic_PeriodID   string            `gorm:"size:191"`
	Academic_Period     Academic_Period   `gorm :"ForeignKey: Sesiones"`
	SesionActividad     []SesionActividad
}
