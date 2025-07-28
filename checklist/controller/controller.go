package controller

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type ChecklistController struct{}
// GET /checklist/modulos
func GetModulos(c *gin.Context) {
    data, err := os.ReadFile("checklist.json")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error leyendo el checklist"})
        return
    }

    var jsonData struct {
        Modulos []string `json:"modulos"`
    }

    if err := json.Unmarshal(data, &jsonData); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parseando el checklist"})
        return
    }

    c.JSON(http.StatusOK, jsonData.Modulos)
}

// POST /checklist/modulos
func AddModulo(c *gin.Context) {
    var nuevo struct {
        Modulo string `json:"modulo"`
    }

    if err := c.BindJSON(&nuevo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
        return
    }

    data, err := os.ReadFile("checklist.json")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error leyendo archivo"})
        return
    }

    var jsonData struct {
        Secciones []interface{} `json:"secciones"`
        Modulos   []string      `json:"modulos"`
    }

    if err := json.Unmarshal(data, &jsonData); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parseando JSON"})
        return
    }

    for _, mod := range jsonData.Modulos {
        if mod == nuevo.Modulo {
            c.JSON(http.StatusOK, gin.H{"mensaje": "Módulo ya existe"})
            return
        }
    }

    jsonData.Modulos = append(jsonData.Modulos, nuevo.Modulo)

    updated, _ := json.MarshalIndent(jsonData, "", "  ")
    os.WriteFile("checklist.json", updated, 0644)

    c.JSON(http.StatusOK, gin.H{"mensaje": "Módulo guardado"})
}

func NewChecklistController() *ChecklistController {
	return &ChecklistController{}
}

func (cc *ChecklistController) GetChecklistJSON(c *gin.Context) {
	// Leer archivo
	data, err := os.ReadFile("checklist/checklist.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo leer el archivo JSON"})
		return
	}

	// Parsear JSON en estructura genérica
	var checklist interface{}
	if err := json.Unmarshal(data, &checklist); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "JSON inválido"})
		return
	}

	// Retornar JSON correctamente serializado
	c.JSON(http.StatusOK, checklist)
}

