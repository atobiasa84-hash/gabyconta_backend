package routes

import (
	"github.com/gin-gonic/gin"
	"gabyconta/empresas/controller"
)

func RegistrarRutas(r *gin.Engine) {
	ctrl := controller.EmpresaController{}
	grupo := r.Group("/api/empresas")
	{
		grupo.POST("/", ctrl.Crear)
		grupo.GET("/", ctrl.Listar)
		grupo.GET("/:id", ctrl.ObtenerPorID)
		grupo.PUT("/:id", ctrl.Actualizar)
		grupo.DELETE("/:id", ctrl.Eliminar)
	}
}
