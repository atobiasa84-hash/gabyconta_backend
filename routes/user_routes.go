package routes

import (
	"github.com/gin-gonic/gin"
	"gabyconta/controllers"
	"gabyconta/middlewares"
)

// RegisterUserRoutes define las rutas relacionadas con los usuarios
func RegisterUserRoutes(router *gin.Engine) {
	api := router.Group("/api")

	// Rutas públicas
	api.POST("/register", controllers.Register)
	api.POST("/login", controllers.Login)

	// Rutas protegidas con middleware JWT
	protected := api.Group("/")
	protected.Use(middlewares.AuthMiddleware()) // Se aplica el middleware a todas las rutas dentro
	{
		protected.GET("/profile", controllers.Profile)
	}
}

