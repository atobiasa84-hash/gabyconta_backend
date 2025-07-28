package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	
	"gabyconta/database"


	// Módulos: modelos y rutas
	asientosModel "gabyconta/asientos/model"
	asientosRoutes "gabyconta/asientos/routes"

	cuentasModel "gabyconta/cuentas/model"
	cuentasRoutes "gabyconta/cuentas/routes"

	usuariosModel "gabyconta/usuarios/model"
	usuariosRoutes "gabyconta/usuarios/routes"

	empresasModel "gabyconta/empresas/model"
	empresasRoutes "gabyconta/empresas/routes"

	checklistModel "gabyconta/checklist/model"
	checklistRoutes "gabyconta/checklist/routes"


)

func main() {
	// Conectar a la base de datos
	if _, err := database.ConnectDB(); err != nil {
		log.Fatalf("❌ Error al conectar con la base de datos: %v", err)
	}

	// Obtener instancia de DB
	db := database.GetDB()

	// Migración automática de modelos
	if err := db.AutoMigrate(
		&asientosModel.AsientoContable{},
		&cuentasModel.Cuenta{},
		&usuariosModel.Usuario{},
		&empresasModel.Empresa{},
		&checklistModel.Tarea{},
		); err != nil {
		log.Fatalf("❌ Error en migración automática: %v", err)
	}
	log.Println("✅ Migración automática completada.")

	// Crear router
	router := gin.Default()
	router.Use(cors.Default())

	// ⬇️ Aquí agregas el middleware CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // tu frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))

	// Configurar proxies confiables (localhost)
	if err := router.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		log.Fatalf("❌ Error configurando proxies confiables: %v", err)
	}

	
	// Ruta raíz para prueba
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"mensaje": "Servidor corriendo correctamente."})
	})

	// Rutas de los módulos
	asientosRoutes.RegistrarRutas(router)
	cuentasRoutes.RegistrarRutas(router)
	usuariosRoutes.UsuariosRoutes(router, db)
	empresasRoutes.RegistrarRutas(router)
	checklistGroup := router.Group("/checklist")
    checklistRoutes.RegisterChecklistRoutes(checklistGroup)
	

	
	

	// Puerto desde variable de entorno o por defecto 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		router.Run(":8080")
		
	}

	log.Printf("🚀 Servidor corriendo en http://localhost:%s\n", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("❌ No se pudo iniciar el servidor: %v", err)
	}
}
