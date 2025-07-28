package model

import "gorm.io/gorm"

type Tarea struct {
	gorm.Model
	Modulo      string `json:"modulo"`
	Descripcion string `json:"descripcion"`
	Completado  bool   `json:"completado"`
}


