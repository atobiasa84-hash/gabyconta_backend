package models

import "time"

type DetalleAsiento struct {
	ID        uint `gorm:"primaryKey"`
	AsientoID uint `gorm:"not null"`
	CuentaID  uint `gorm:"not null"`
	Cuenta    CuentaContable
	Debe      float64
	Haber     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
