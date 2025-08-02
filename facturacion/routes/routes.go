package routes

import (
    "github.com/gin-gonic/gin"
    "gabyconta/facturacion/controller"
    "gabyconta/facturacion/service"
    "gorm.io/gorm"
)

func RegisterFacturacionRoutes(r *gin.RouterGroup, db *gorm.DB) {
    svc := service.NewFacturacionService(db)
    ctrl := controller.NewFacturacionController(svc)

    fact := r.Group("/facturas")
    {
        fact.POST("/", ctrl.CrearFactura)
        fact.GET("/", ctrl.ObtenerFacturas)
    }
}
