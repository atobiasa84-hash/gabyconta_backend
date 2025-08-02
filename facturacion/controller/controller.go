package controller

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "gabyconta/facturacion/model"
    "gabyconta/facturacion/service"
)

type FacturacionController struct {
    Service *service.FacturacionService
}

func NewFacturacionController(s *service.FacturacionService) *FacturacionController {
    return &FacturacionController{Service: s}
}

func (c *FacturacionController) CrearFactura(ctx *gin.Context) {
    var factura model.Factura
    if err := ctx.ShouldBindJSON(&factura); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := c.Service.CrearFactura(&factura); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear la factura"})
        return
    }

    ctx.JSON(http.StatusOK, factura)
}

func (c *FacturacionController) ObtenerFacturas(ctx *gin.Context) {
    empresaIDStr := ctx.Query("empresa_id")
    empresaID, err := strconv.Atoi(empresaIDStr)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Empresa ID inválido"})
        return
    }

    facturas, err := c.Service.ObtenerFacturasPorEmpresa(uint(empresaID))
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener las facturas"})
        return
    }

    ctx.JSON(http.StatusOK, facturas)
}
