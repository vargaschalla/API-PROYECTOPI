package models

import "gorm.io/gorm"

type Sesiones struct {
	gorm.Model
	Descripcion       string          `json:"descripcion"`
	Orden             string          `json:"orden"`
	SeccionID         string          `gorm:"size:191"`
	Seccion           Seccion         `gorm :"ForeignKey: Sesiones"`
	Academic_PlaneID  string          `gorm:"size:191"`
	Academic_Plane    Academic_Plane  `gorm :"ForeignKey: Sesiones"`
	Academic_PeriodID string          `gorm:"size:191"`
	Academic_Period   Academic_Period `gorm :"ForeignKey: Sesiones"`
	FechaPlanificada  string          `json:"fechaplanificada"`
	Estado            string          `json:"estado"`
	SesionMaterial    []SesionMaterial
	SesionActividad   []SesionActividad
}
