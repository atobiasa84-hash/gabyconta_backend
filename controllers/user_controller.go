package controllers

import (
	"gabyconta/database"
	"gabyconta/models"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// Register permite registrar un nuevo usuario
func Register(c *gin.Context) {
	var input models.User

	// Validar el JSON recibido
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Validar campos obligatorios
	if strings.TrimSpace(input.Name) == "" || strings.TrimSpace(input.Email) == "" ||
		strings.TrimSpace(input.Password) == "" || strings.TrimSpace(input.Role) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todos los campos son obligatorios"})
		return
	}

	// Verificar si ya existe ese email
	var existente models.User
	if err := database.DB.Where("email = ?", input.Email).First(&existente).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "El email ya está registrado"})
		return
	}

	// Hashear la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al encriptar la contraseña"})
		return
	}
	input.Password = string(hashedPassword)

	// Guardar el nuevo usuario
	if err := database.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al registrar usuario"})
		return
	}

	// Retornar sin la contraseña
	input.Password = ""
	c.JSON(http.StatusCreated, gin.H{"message": "Usuario creado exitosamente", "user": input})
}

// Login autentica al usuario y genera un JWT válido
func Login(c *gin.Context) {
	var input models.User
	var user models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Buscar usuario por email
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no encontrado"})
		return
	}

	// Comparar contraseñas
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Contraseña incorrecta"})
		return
	}

	// Generar token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo generar el token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Inicio de sesión exitoso",
		"token":   tokenString,
	})
}

// Profile es una ruta protegida que muestra el usuario autenticado
func Profile(c *gin.Context) {
	userID, _ := c.Get("user_id")
	role, _ := c.Get("role")

	c.JSON(http.StatusOK, gin.H{
		"message": "Ruta protegida con éxito",
		"user_id": userID,
		"role":    role,
	})
}
