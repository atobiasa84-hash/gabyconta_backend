package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gabyconta/asientos/model"
	"gabyconta/asientos/service"
)

// CrearAsientoHandler crea un nuevo asiento contable
func CrearAsientoHandler(c *gin.Context) {
	var asiento model.AsientoContable
	if err := c.ShouldBindJSON(&asiento); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos", "detalle": err.Error()})
		return
	}

	if err := service.CrearAsiento(&asiento); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear el asiento", "detalle": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, asiento)
}

// ObtenerAsientosHandler lista todos los asientos con sus detalles
func ObtenerAsientosHandler(c *gin.Context) {
	asientos, err := service.ObtenerAsientos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener asientos", "detalle": err.Error()})
		return
	}
	c.JSON(http.StatusOK, asientos)
}

// ObtenerAsientoPorIDHandler obtiene un asiento por su ID
func ObtenerAsientoPorIDHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	asiento, err := service.ObtenerAsientoPorID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Asiento no encontrado", "detalle": err.Error()})
		return
	}

	c.JSON(http.StatusOK, asiento)
}

// ActualizarAsientoHandler actualiza un asiento existente
func ActualizarAsientoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var datosActualizar model.AsientoContable
	if err := c.ShouldBindJSON(&datosActualizar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos", "detalle": err.Error()})
		return
	}

	// Se asegura que el ID del objeto sea el mismo recibido
	datosActualizar.ID = uint(id)

	if err := service.ActualizarAsiento(&datosActualizar); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar asiento", "detalle": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Asiento actualizado correctamente", "asiento": datosActualizar})
}

// EliminarAsientoHandler elimina un asiento por su ID
func EliminarAsientoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := service.EliminarAsiento(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar asiento", "detalle": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Asiento eliminado correctamente"})
}
