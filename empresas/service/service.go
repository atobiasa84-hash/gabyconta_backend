package service

import (
	"gabyconta/database"
	"gabyconta/empresas/model"
)

func CrearEmpresa(empresa *model.Empresa) error {
	return database.DB.Create(empresa).Error
}

func ObtenerEmpresas() ([]model.Empresa, error) {
	var empresas []model.Empresa
	err := database.DB.Find(&empresas).Error
	return empresas, err
}

func ObtenerEmpresaPorID(id uint) (*model.Empresa, error) {
	var empresa model.Empresa
	err := database.DB.First(&empresa, id).Error
	return &empresa, err
}

func ActualizarEmpresa(id uint, datos *model.Empresa) error {
	return database.DB.Model(&model.Empresa{}).Where("id = ?", id).Updates(datos).Error
}

func EliminarEmpresa(id uint) error {
	return database.DB.Delete(&model.Empresa{}, id).Error
}
