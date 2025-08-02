package service

import (
	"errors"
	"gabyconta/cuentas/model"
	"gabyconta/database"

	"gorm.io/gorm"
)

func CrearCuenta(cuenta *model.Cuenta) error {
	return database.DB.Create(cuenta).Error
}

func ObtenerCuentas() ([]model.Cuenta, error) {
	var cuentas []model.Cuenta
	err := database.DB.Find(&cuentas).Error
	return cuentas, err
}

func ObtenerCuentaPorID(id uint) (*model.Cuenta, error) {
	var cuenta model.Cuenta
	err := database.DB.First(&cuenta, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &cuenta, err
}

func ActualizarCuenta(id uint, datos *model.Cuenta) error {
	cuentaExistente, err := ObtenerCuentaPorID(id)
	if err != nil {
		return err
	}
	if cuentaExistente == nil {
		return gorm.ErrRecordNotFound
	}

	return database.DB.Model(&cuentaExistente).Updates(datos).Error
}

func EliminarCuenta(id uint) error {
	return database.DB.Delete(&model.Cuenta{}, id).Error
}
