package models

import "gorm.io/gorm"

type CargaAcademica struct {
	gorm.Model
	Nombre  string `json:"nombre"`
	Paterno string `json:"paterno"`
	Materno string `json:"materno"`
	Email   string `json:"email"`
	Edad    string `json:"edad"`
	//Edad            int `json:"edad"`
	Celular string `json:"celular"`
	//Celular         int `json:"celular"`
	Fechanacimiento string `json:"fechanacimiento"`
	//Fechanacimiento time.Time `json:"fechanacimiento"`
	DNI string `json:"dni"`
	//DNI    int `json:"dni"`
	Estado string `json:"estado"`
}
