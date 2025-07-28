package controller

import (
	"net/http"
	"strconv"

	"gabyconta/empresas/model"
	"gabyconta/empresas/service"

	"github.com/gin-gonic/gin"
)

type EmpresaController struct{}

func (ctrl *EmpresaController) Crear(c *gin.Context) {
	var empresa model.Empresa
	if err := c.ShouldBindJSON(&empresa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.CrearEmpresa(&empresa); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear la empresa"})
		return
	}
	c.JSON(http.StatusCreated, empresa)
}

func (ctrl *EmpresaController) Listar(c *gin.Context) {
	empresas, err := service.ObtenerEmpresas()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener las empresas"})
		return
	}
	c.JSON(http.StatusOK, empresas)
}

func (ctrl *EmpresaController) ObtenerPorID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	empresa, err := service.ObtenerEmpresaPorID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Empresa no encontrada"})
		return
	}
	c.JSON(http.StatusOK, empresa)
}

func (ctrl *EmpresaController) Actualizar(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var datos model.Empresa
	if err := c.ShouldBindJSON(&datos); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.ActualizarEmpresa(uint(id), &datos); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar la empresa"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"mensaje": "Empresa actualizada correctamente"})
}

func (ctrl *EmpresaController) Eliminar(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := service.EliminarEmpresa(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar la empresa"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"mensaje": "Empresa eliminada correctamente"})
}
