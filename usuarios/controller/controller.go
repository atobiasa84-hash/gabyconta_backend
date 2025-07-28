package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gabyconta/usuarios/model"
	"gabyconta/usuarios/service"
)

type UsuarioController struct {
	Service *service.UsuarioService
}

func NewUsuarioController(s *service.UsuarioService) *UsuarioController {
	return &UsuarioController{Service: s}
}

func (uc *UsuarioController) Registro(c *gin.Context) {
	var input model.Usuario
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	if err := uc.Service.CrearUsuario(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear el usuario: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"mensaje": "Usuario creado correctamente"})
}

func (uc *UsuarioController) Login(c *gin.Context) {
	var login model.LoginInput
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	usuario, err := uc.Service.Autenticar(login.Email, login.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensaje": "Autenticación exitosa",
		"usuario": gin.H{
			"id":     usuario.ID,
			"nombre": usuario.Nombre,
			"email":  usuario.Email,
		},
	})
}
