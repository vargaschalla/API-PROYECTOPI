package models

import "gorm.io/gorm"

type Modulo struct {
	gorm.Model
	Descripcion       string          `json:"descripcion"`
	Orden             string          `json:"orden"`
	Estado            string          `json:"estado"`
	Academic_PeriodID string          `gorm:"size:191"`
	Academic_Period   Academic_Period `gorm :"ForeignKey: Modulo"`
	Academic_PlaneID  string          `gorm:"size:191"`
	Academic_Plane    Academic_Plane  `gorm :"ForeignKey: Modulo"`
}
