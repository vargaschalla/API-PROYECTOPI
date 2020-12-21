package models

import "gorm.io/gorm"

type CargaAcademica struct {
	gorm.Model
	NotaFinal         string          `json:"notafinal"`
	FechaRegistro     string          `json:"fecharegistro"`
	Estado            string          `json:"estado"`
	Academic_PlaneID  string          `gorm:"size:191"`
	Academic_Plane    Academic_Plane  `gorm :"ForeignKey: CargaAcademica"`
	Academic_PeriodID string          `gorm:"size:191"`
	Academic_Period   Academic_Period `gorm :"ForeignKey: CargaAcademica"`
	AlumnoID          string          `gorm:"size:191"`
	Alumno            Alumno          `gorm :"ForeignKey: CargaAcademica"`
	SeccionID         string          `gorm:"size:191"`
	Seccion           Seccion         `gorm :"ForeignKey: CargaAcademica"`
}
