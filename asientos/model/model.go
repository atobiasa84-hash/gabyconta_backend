package model

import (
	"time"

	"gorm.io/gorm"
)

// AsientoContable representa el asiento contable principal
type AsientoContable struct {
	ID          uint             `gorm:"primaryKey;autoIncrement" json:"id"`
	Fecha       time.Time        `json:"fecha" binding:"required"`
	Glosa       string           `json:"glosa" binding:"required"`
	Descripcion string           `json:"descripcion" binding:"required"`
	Detalles    []DetalleAsiento `gorm:"foreignKey:AsientoID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"detalles,omitempty"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
	DeletedAt   gorm.DeletedAt   `gorm:"index" json:"-"`
}

// DetalleAsiento representa el detalle de cada asiento contable
type DetalleAsiento struct {
	ID          uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	AsientoID   uint           `json:"asiento_id" binding:"required"` // FK hacia AsientoContable
	CuentaID    uint           `json:"cuenta_id" binding:"required"`  // FK hacia CuentaContable
	Debe        float64        `json:"debe" binding:"gte=0"`          // monto debitado
	Haber       float64        `json:"haber" binding:"gte=0"`         // monto acreditado
	Descripcion string         `json:"descripcion"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
