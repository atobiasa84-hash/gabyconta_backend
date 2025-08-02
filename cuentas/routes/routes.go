package routes

import (
	"github.com/gin-gonic/gin"
	"gabyconta/cuentas/controller"
)

func RegistrarRutas(router *gin.Engine) {
	rutas := router.Group("/api/cuentas")
	{
		rutas.POST("/", controller.CrearCuentaHandler)
		rutas.GET("/", controller.ObtenerCuentasHandler)
		rutas.GET("/:id", controller.ObtenerCuentaPorIDHandler)
		rutas.PUT("/:id", controller.ActualizarCuentaHandler)
		rutas.DELETE("/:id", controller.EliminarCuentaHandler)
	}
}