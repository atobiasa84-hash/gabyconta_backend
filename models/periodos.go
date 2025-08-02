package models

import "time"

type PeriodoContable struct {
	ID         uint      `gorm:"primaryKey"`
	Nombre     string    `gorm:"not null"` // Ej: Enero 2025
	FechaInicio time.Time
	FechaFin    time.Time
	Activo      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
