package model

import (
	"time"

	"gorm.io/gorm"
)

type Cuenta struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	Codigo       string         `gorm:"unique;not null" json:"codigo" binding:"required"`
	Nombre       string         `gorm:"not null" json:"nombre"`
	Tipo         string         `gorm:"not null" json:"tipo"` // Activo, Pasivo, Patrimonio, etc.
	SaldoInicial float64        `json:"saldo_inicial"`
	EmpresaID    uint           `json:"empresa_id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	Empresa      uint           `gorm:"foreignKey:EmpresaID" json:"-"`
	
}
