package service

import (
	"errors"

	"gabyconta/asientos/model"
	"gabyconta/database"

	"gorm.io/gorm"
)

// CrearAsiento guarda un nuevo asiento contable con sus detalles
func CrearAsiento(asiento *model.AsientoContable) error {
	db := database.GetDB()

	// Crear el asiento y sus detalles (genera insert en tablas relacionadas)
	return db.Create(asiento).Error
}

// ObtenerAsientos devuelve todos los asientos con sus detalles
func ObtenerAsientos() ([]model.AsientoContable, error) {
	db := database.GetDB()
	var asientos []model.AsientoContable

	// Preload para cargar detalles relacionados
	if err := db.Preload("Detalles").Find(&asientos).Error; err != nil {
		return nil, err
	}

	return asientos, nil
}

// ObtenerAsientoPorID obtiene un asiento por su ID con detalles
func ObtenerAsientoPorID(id uint) (model.AsientoContable, error) {
	db := database.GetDB()
	var asiento model.AsientoContable

	if err := db.Preload("Detalles").First(&asiento, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return asiento, errors.New("asiento no encontrado")
		}
		return asiento, err
	}

	return asiento, nil
}

// ActualizarAsiento actualiza un asiento existente
func ActualizarAsiento(asiento *model.AsientoContable) error {
	db := database.GetDB()

	// Guardar los cambios del asiento
	if err := db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(asiento).Error; err != nil {
		return err
	}

	return nil
}

// EliminarAsiento elimina un asiento por su ID junto con detalles por cascada
func EliminarAsiento(id uint) error {
	db := database.GetDB()
	asiento := model.AsientoContable{ID: id}

	if err := db.Delete(&asiento).Error; err != nil {
		return err
	}

	return nil
}
