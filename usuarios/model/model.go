package model

import (
	"time"

	"gorm.io/gorm"
	"gabyconta/empresas/model" // Ajusta según tu estructura
)

// Usuario representa un usuario en el sistema
type Usuario struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Nombre    string         `json:"nombre"`
	Email     string         `gorm:"unique;not null" json:"email"`
	Password  string         `json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Empresas []model.Empresa `gorm:"many2many:usuario_empresas;" json:"empresas"`
}

// LoginInput estructura para login
type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
