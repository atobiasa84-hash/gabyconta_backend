package migrate

import (
	"log"

	"gabyconta/database/migrate/manual"
)

// RunManualMigrations ejecuta todos los scripts manuales (de forma ordenada)
func RunManualMigrations() error {
	migrations := []func() error{
		manual.Up20250724CreateUsuarios,
		// manual.Up20250725CreateProductos,
		// manual.Up20250726CreateAsientos,
	}

	for _, m := range migrations {
		if err := m(); err != nil {
			log.Printf("❌ Error en una migración manual: %v", err)
			return err
		}
	}
	log.Println("✅ Todas las migraciones manuales aplicadas")
	return nil
}
