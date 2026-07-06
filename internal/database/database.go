package database

import (
	"log"
	"os"

	"barberia-cola-virtual/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB será la conexión global a PostgreSQL.
// Si está en nil, el proyecto puede seguir usando storage en memoria.
var DB *gorm.DB

// ConnectDB conecta la API con PostgreSQL cuando DB_DRIVER=postgres.
// En Docker, esas variables vienen desde docker-compose.yml.
func ConnectDB() {
	driver := os.Getenv("DB_DRIVER")
	dsn := os.Getenv("DATABASE_DSN")

	// Si no estamos en modo postgres, no conectamos base.
	// Así el proyecto local puede seguir funcionando con memoria.
	if driver != "postgres" {
		log.Println("DB_DRIVER no es postgres: usando almacenamiento en memoria")
		return
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error conectando a PostgreSQL:", err)
	}

	// AutoMigrate crea las tablas si no existen.
	err = db.AutoMigrate(
		&models.Servicio{},
		&models.CategoriaServicio{},
		&models.Promocion{},
		&models.Usuario{},
		&models.Turno{},
	)
	if err != nil {
		log.Fatal("Error ejecutando AutoMigrate:", err)
	}

	DB = db
	log.Println("Conectado a PostgreSQL correctamente")
}