package model

import (
	"time"

	"gorm.io/gorm"
)

type Empresa struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Nombre    string         `gorm:"not null" json:"nombre"`
	RUC       string         `gorm:"unique;not null" json:"ruc"`
	Direccion string         `json:"direccion"`
	Telefono  string         `json:"telefono"`
	Email     string         `json:"email"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
