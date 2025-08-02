package routes

import (
	"gabyconta/asientos/controller"

	"github.com/gin-gonic/gin"
)

// RegistrarRutas configura las rutas para el módulo Asientos
func RegistrarRutas(r *gin.Engine) {
	asientos := r.Group("/api/asientos")
	{
		asientos.POST("/", controller.CrearAsientoHandler)
		asientos.GET("/", controller.ObtenerAsientosHandler)
		asientos.GET("/:id", controller.ObtenerAsientoPorIDHandler)
		asientos.PUT("/:id", controller.ActualizarAsientoHandler)
		asientos.DELETE("/:id", controller.EliminarAsientoHandler)
	}
}
