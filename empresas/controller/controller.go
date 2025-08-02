package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"gabyconta/empresas/model"
	"gabyconta/empresas/service"
)

// CrearEmpresa maneja la solicitud para registrar una nueva empresa
func CrearEmpresa(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var empresa model.Empresa
		if err := c.ShouldBindJSON(&empresa); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos", "detalles": err.Error()})
			return
		}

		nuevaEmpresa, err := service.CrearEmpresa(db, &empresa)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No se pudo crear la empresa", "detalles": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, nuevaEmpresa)
	}
}

// ObtenerEmpresas maneja la solicitud para listar todas las empresas
func ObtenerEmpresas(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		empresas, err := service.ObtenerEmpresas(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener las empresas", "detalles": err.Error()})
			return
		}
		c.JSON(http.StatusOK, empresas)
	}
}

// ObtenerEmpresaPorID maneja la solicitud para obtener una empresa por su ID
func ObtenerEmpresaPorID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}

		empresa, err := service.ObtenerEmpresaPorID(db, uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Empresa no encontrada", "detalles": err.Error()})
			return
		}
		c.JSON(http.StatusOK, empresa)
	}
}

// ActualizarEmpresa maneja la solicitud para actualizar los datos de una empresa
func ActualizarEmpresa(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}

		var datos model.Empresa
		if err := c.ShouldBindJSON(&datos); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos", "detalles": err.Error()})
			return
		}

		err = service.ActualizarEmpresa(db, uint(id), &datos)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar la empresa", "detalles": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"mensaje": "Empresa actualizada correctamente"})
	}
}

// EliminarEmpresa maneja la solicitud para eliminar una empresa por ID
func EliminarEmpresa(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}

		err = service.EliminarEmpresa(db, uint(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar la empresa", "detalles": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"mensaje": "Empresa eliminada correctamente"})
	}
}
