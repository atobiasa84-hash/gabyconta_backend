package model

import (
    "time"

    "gorm.io/gorm"
)

type Factura struct {
    ID           uint           `gorm:"primaryKey" json:"id"`
    Numero       string         `json:"numero"`
    Fecha        time.Time      `json:"fecha"`
    Cliente      string         `json:"cliente"`
    RUC          string         `json:"ruc"`
    Direccion    string         `json:"direccion"`
    Total        float64        `json:"total"`
    EmpresaID    uint           `json:"empresa_id"`
    UsuarioID    uint           `json:"usuario_id"`
    Estado       string         `json:"estado"` // borrador, emitida, autorizada, anulada
    XMLGenerado  bool           `json:"xml_generado"`
    CreatedAt    time.Time
    UpdatedAt    time.Time
    DeletedAt    gorm.DeletedAt `gorm:"index"`
    Detalles     []DetalleFactura `gorm:"foreignKey:FacturaID"`
}

type DetalleFactura struct {
    ID          uint    `gorm:"primaryKey" json:"id"`
    FacturaID   uint    `json:"factura_id"`
    Producto    string  `json:"producto"`
    Cantidad    int     `json:"cantidad"`
    PrecioUnit  float64 `json:"precio_unit"`
    Subtotal    float64 `json:"subtotal"`
}
