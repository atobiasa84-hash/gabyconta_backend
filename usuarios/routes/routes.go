package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"gabyconta/usuarios/controller"
	"gabyconta/usuarios/service"
)

func UsuariosRoutes(r *gin.Engine, db *gorm.DB) {
	s := service.NewUsuarioService(db)
	c := controller.NewUsuarioController(s)

	api := r.Group("/api/usuarios")
	{
		api.POST("/registro", c.Registro)
		api.POST("/login", c.Login)
	}
}
