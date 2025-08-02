package service

import (
	"errors"

	"gorm.io/gorm"

	"gabyconta/empresas/model"
)

// CrearEmpresa intenta crear una nueva empresa, validando que el RUC no esté duplicado.
func CrearEmpresa(db *gorm.DB, empresa *model.Empresa) (*model.Empresa, error) {
	var existente model.Empresa

	// Validar RUC único
	if err := db.Where("ruc = ?", empresa.RUC).First(&existente).Error; err == nil {
		return nil, errors.New("ya existe una empresa con este RUC")
	}

	if err := db.Create(empresa).Error; err != nil {
		return nil, err
	}

	return empresa, nil
}

// ObtenerEmpresas devuelve todas las empresas
func ObtenerEmpresas(db *gorm.DB) ([]model.Empresa, error) {
	var empresas []model.Empresa
	err := db.Find(&empresas).Error
	return empresas, err
}

// ObtenerEmpresaPorID busca una empresa por su ID
func ObtenerEmpresaPorID(db *gorm.DB, id uint) (*model.Empresa, error) {
	var empresa model.Empresa
	err := db.First(&empresa, id).Error
	return &empresa, err
}

// ActualizarEmpresa actualiza los datos de una empresa específica
func ActualizarEmpresa(db *gorm.DB, id uint, datos *model.Empresa) error {
	return db.Model(&model.Empresa{}).Where("id = ?", id).Updates(datos).Error
}

// EliminarEmpresa elimina una empresa por su ID
func EliminarEmpresa(db *gorm.DB, id uint) error {
	return db.Delete(&model.Empresa{}, id).Error
}
