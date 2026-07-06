package db

import (
	"log"
	"os"

	"barberia-cola-virtual/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitPostgres conecta el proyecto con PostgreSQL usando GORM.
// El DSN viene desde docker-compose.yml mediante la variable DATABASE_DSN.
func InitPostgres() *gorm.DB {
	dsn := os.Getenv("DATABASE_DSN")

	if dsn == "" {
		log.Fatal("DATABASE_DSN no está configurado")
	}

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("No se pudo conectar a PostgreSQL: %v", err)
	}

	err = database.AutoMigrate(
		&models.Usuario{},
		&models.Cliente{},
		&models.PreferenciaPago{},
		&models.PreferenciaCliente{},
		&models.CategoriaServicio{},
		&models.Servicio{},
		&models.Promocion{},
		&models.Turno{},
		&models.SeguimientoTurno{},
		&models.Notificacion{},
	)

	if err != nil {
		log.Fatalf("Error ejecutando migraciones en PostgreSQL: %v", err)
	}

	log.Println("Conectado a PostgreSQL correctamente")

	return database
}