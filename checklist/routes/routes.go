package routes

import (
	"github.com/gin-gonic/gin"
	"gabyconta/checklist/controller"
)

func RegisterChecklistRoutes(router *gin.RouterGroup) {
	ctrl := controller.NewChecklistController()
	router.GET("/json", ctrl.GetChecklistJSON)
    router.GET("/checklist/modulos", controller.GetModulos)
    router.POST("/checklist/modulos", controller.AddModulo)
}

