package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"gabyconta/empresas/controller"
)
func RegisterEmpresaRoutes(r *gin.RouterGroup, db *gorm.DB) {
	r.POST("/", controller.CrearEmpresa(db))
	r.GET("/", controller.ObtenerEmpresas(db))
	r.GET("/:id", controller.ObtenerEmpresaPorID(db))
	r.PUT("/:id", controller.ActualizarEmpresa(db))
	r.DELETE("/:id", controller.EliminarEmpresa(db))
}
