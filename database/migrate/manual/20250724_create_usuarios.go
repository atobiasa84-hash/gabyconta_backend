
package manual

import (
	"log"

	"gabyconta/database"
	"gabyconta/usuarios/model"
)

// Up20250724CreateUsuarios crea la tabla usuarios
func Up20250724CreateUsuarios() error {
	db := database.GetDB()

	err := db.AutoMigrate(&model.Usuario{})
	if err != nil {
		return err
	}

	log.Println("✅ Migración 20250724_create_usuarios aplicada correctamente")
	return nil
}
