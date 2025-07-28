package models

import "time"

type CuentaContable struct {
	ID       uint   `gorm:"primaryKey"`
	Codigo   string `gorm:"unique;not null"` // Ej: 1.1.01
	Nombre   string `gorm:"not null"`        // Ej: Caja general
	Tipo     string `gorm:"not null"`        // Activo, Pasivo, etc.
	Nivel    int
	PadreID  *uint
	Padre    *CuentaContable `gorm:"foreignKey:PadreID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}