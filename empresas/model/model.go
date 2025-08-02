package model

import (
	"time"

	"gorm.io/gorm"
)

type Empresa struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	Nombre          string         `gorm:"not null" json:"nombre"`
	RazonSocial     string         `gorm:"not null" json:"razon_social"`
	RUC             string         `gorm:"uniqueIndex:uni_empresas_ruc;not null" json:"ruc"`
	Direccion       string         `json:"direccion"`
	Telefono        string         `json:"telefono"`
	Email           string         `json:"email"`
	InicioActividad string         `json:"inicio_actividad"`
	CierreActividad string         `json:"cierre_actividad"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`

	UsuariosIDs []uint `gorm:"-" json:"usuarios_ids,omitempty"` // Solo en memoria
}

type UsuarioEmpresa struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	UsuarioID  uint      `gorm:"not null;uniqueIndex:idx_usuario_empresa" json:"usuario_id"`
	EmpresaID  uint      `gorm:"not null;uniqueIndex:idx_usuario_empresa" json:"empresa_id"`
	CreatedAt  time.Time `json:"created_at"`
}
