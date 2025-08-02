package controller

import (
	"net/http"
	"strconv"

		    
	
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"gabyconta/cuentas/model"
	"gabyconta/cuentas/service"
)

func CrearCuentaHandler(c *gin.Context) {
	var cuenta model.Cuenta
	if err := c.ShouldBindJSON(&cuenta); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.CrearCuenta(&cuenta); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la cuenta"})
		return
	}

	c.JSON(http.StatusCreated, cuenta)
}

func ObtenerCuentasHandler(c *gin.Context) {
	cuentas, err := service.ObtenerCuentas()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las cuentas"})
		return
	}
	c.JSON(http.StatusOK, cuentas)
}

func ObtenerCuentaPorIDHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	cuenta, err := service.ObtenerCuentaPorID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener la cuenta"})
		return
	}
	if cuenta == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cuenta no encontrada"})
		return
	}
	c.JSON(http.StatusOK, cuenta)
}

func ActualizarCuentaHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var datos model.Cuenta
	if err := c.ShouldBindJSON(&datos); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.ActualizarCuenta(uint(id), &datos); err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			c.JSON(http.StatusNotFound, gin.H{"error": "Cuenta no encontrada"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar la cuenta"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Cuenta actualizada correctamente"})
}

func EliminarCuentaHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := service.EliminarCuenta(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar la cuenta"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"mensaje": "Cuenta eliminada correctamente"})
}
