package models

import "time"

type Tercero struct {
	ID        uint   `gorm:"primaryKey"`
	Nombre    string `gorm:"not null"`
	Tipo      string // Cliente, Proveedor, Empleado, etc.
	Identificacion string
	CreatedAt time.Time
	UpdatedAt time.Time
}

