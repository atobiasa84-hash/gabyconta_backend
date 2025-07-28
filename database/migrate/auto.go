
package migrate

import (
	"log"

	"gabyconta/database"
	"gabyconta/usuarios/model"
)

// AutoMigrateModels ejecuta migraciones automáticas de GORM para desarrollo rápido
func AutoMigrateModels() error {
	db := database.GetDB()

	models := []interface{}{
		&model.Usuario{},
		// &asientos.Modelo{},
		// &productos.Modelo{},
	}

	if err := db.AutoMigrate(models...); err != nil {
		return err
	}

	log.Println("📦 Migraciones automáticas ejecutadas")
	return nil
}
