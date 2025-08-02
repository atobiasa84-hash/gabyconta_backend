package database

import (
	"context"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB inicializa la conexión a PostgreSQL con contexto y timeout
func ConnectDB() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		dsn = "host=localhost user=admin password=superuser dbname=gabyconta port=5432 sslmode=disable"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err := sqlDB.PingContext(ctx); err != nil {
		return nil, err
	}

	DB = db
	log.Println("✅ Conexión exitosa a PostgreSQL")
	return DB, nil
}

// GetDB retorna la instancia global, o falla si no está inicializada
func GetDB() *gorm.DB {
	if DB == nil {
		log.Fatal("❌ Base de datos no inicializada. Ejecuta ConnectDB() primero.")
	}
	return DB
}
