package service

import (
    "gabyconta/facturacion/model"
    "gorm.io/gorm"
)

type FacturacionService struct {
    DB *gorm.DB
}

func NewFacturacionService(db *gorm.DB) *FacturacionService {
    return &FacturacionService{DB: db}
}

func (s *FacturacionService) CrearFactura(factura *model.Factura) error {
    return s.DB.Create(factura).Error
}

func (s *FacturacionService) ObtenerFacturasPorEmpresa(empresaID uint) ([]model.Factura, error) {
    var facturas []model.Factura
    err := s.DB.Preload("Detalles").Where("empresa_id = ?", empresaID).Find(&facturas).Error
    return facturas, err
}
