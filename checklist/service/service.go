package service

import (
	"gabyconta/checklist/model"
	"gabyconta/database"
)

func CrearTarea(tarea *model.Tarea) error {
	return database.DB.Create(tarea).Error
}

func ObtenerTareas() ([]model.Tarea, error) {
	var tareas []model.Tarea
	err := database.DB.Find(&tareas).Error
	return tareas, err
}

func ActualizarTarea(tarea *model.Tarea) error {
	return database.DB.Save(tarea).Error
}

func EliminarTarea(id uint) error {
	return database.DB.Delete(&model.Tarea{}, id).Error
}

