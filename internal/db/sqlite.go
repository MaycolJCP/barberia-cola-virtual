package db

import (
	"barberia-cola-virtual/internal/models"
	"log"
	"os"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB inicializa la base de datos de manera dinámica (Postgres o SQLite) según el entorno.
func InitDB(dbPath string) *gorm.DB {
	var db *gorm.DB
	var err error

	driver := os.Getenv("DB_DRIVER")
	dsn := os.Getenv("DB_DSN")

	if driver == "postgres" {
		log.Println("Conectando a la base de datos PostgreSQL en Docker...")
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	} else {
		log.Printf("Conectando a la base de datos SQLite local en: %s", dbPath)
		db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	}

	if err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}

	// Forzar la activación de Foreign Keys solo si estamos usando SQLite
	if driver != "postgres" {
		db.Exec("PRAGMA foreign_keys = ON;")
	}

	// Ejecutar la migración automática de todas las tablas
	err = db.AutoMigrate(
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
		log.Fatalf("Error ejecutando la migración estructural de GORM: %v", err)
	}

	return db
}
