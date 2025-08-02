package models

import "time"

type AsientoContable struct {
	ID          uint      `gorm:"primaryKey"`
	Descripcion string    `gorm:"not null"`
	Fecha       time.Time `gorm:"not null"`
	PeriodoID   uint
	Periodo     PeriodoContable
	TerceroID   *uint
	Tercero     *Tercero
	Detalles    []DetalleAsiento `gorm:"foreignKey:AsientoID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
