package models

import "gorm.io/gorm"

type Grado struct {
	gorm.Model
	Descripcion    string `json:"descripcion"`
	NivelID        string `gorm:"size:191"`
	Nivel          Nivel
	Estado         string `json:"estado"`
	Academic_Plane []Academic_Plane
}
