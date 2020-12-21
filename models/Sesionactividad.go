package models

import "gorm.io/gorm"

type SesionActividad struct {
	gorm.Model
	Titulo              string            `json:"titulo"`
	Descripcion         string            `json:"descripcion"`
	FechaRegistro       string            `json:"fecharegistro"`
	Ruta                string            `json:"ruta"`
	Estado              string            `json:"estado"`
	AlumnoID            string            `gorm:"size:191"`
	Alumno              Alumno            `gorm :"ForeignKey: SesionActividad"`
	MaterialActividadID string            `gorm:"size:191"`
	MaterialActividad   MaterialActividad `gorm :"ForeignKey: SesionActividad"`
	SesionesID          string            `gorm:"size:191"`
	Sesiones            Sesiones          `gorm :"ForeignKey: SesionActividad"`
	SeccionID           string            `gorm:"size:191"`
	Seccion             Seccion           `gorm :"ForeignKey: SesionActividad"`
	Academic_PeriodID   string            `gorm:"size:191"`
	Academic_Period     Academic_Period   `gorm :"ForeignKey: SesionActividad"`
	Academic_PlaneID    string            `gorm:"size:191"`
	Academic_Plane      Academic_Plane    `gorm :"ForeignKey: SesionActividad"`
	SesionMaterialID    string            `gorm:"size:191"`
	SesionMaterial      SesionMaterial    `gorm :"ForeignKey: SesionActividad"`
}
