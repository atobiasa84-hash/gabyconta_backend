package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"gabyconta/database"

	asientosModel "gabyconta/asientos/model"
	asientosRoutes "gabyconta/asientos/routes"

	cuentasModel "gabyconta/cuentas/model"
	cuentasRoutes "gabyconta/cuentas/routes"

	usuariosModel "gabyconta/usuarios/model"
	usuariosRoutes "gabyconta/usuarios/routes"

	empresasModel "gabyconta/empresas/model"
	empresasRoutes "gabyconta/empresas/routes"

	chec "gabyconta/checklist/model"
	checklistRoutes "gabyconta/checklist/routes"
)

func main() {
	if _, err := database.ConnectDB(); err != nil {
		log.Fatalf("❌ Error al conectar con la base de datos: %v", err)
	}

	db := database.GetDB()

	if err := db.AutoMigrate(
		&asientosModel.AsientoContable{},
		&cuentasModel.Cuenta{},
		&usuariosModel.Usuario{},
		&empresasModel.Empresa{},
		&empresasModel.UsuarioEmpresa{},
		&chec.Tarea{},
	); err != nil {
		log.Fatalf("❌ Error en migración automática: %v", err)
	}
	log.Println("✅ Migración automática completada.")
	fmt.Println("✅ Migración completada")

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))

	if err := router.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		log.Fatalf("❌ Error configurando proxies confiables: %v", err)
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"mensaje": "Servidor corriendo correctamente."})
	})

	asientosRoutes.RegistrarRutas(router)
	cuentasRoutes.RegistrarRutas(router)
	usuariosRoutes.UsuariosRoutes(router, db)
	empresaGroup := router.Group("/empresas")
    empresasRoutes.RegisterEmpresaRoutes(empresaGroup, db)
	
	checklistGroup := router.Group("/checklist")
	checklistRoutes.RegisterChecklistRoutes(checklistGroup)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Servidor corriendo en http://localhost:%s\n", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("❌ No se pudo iniciar el servidor: %v", err)
	}
}
