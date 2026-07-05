package db

import (
	"barberia-cola-virtual/internal/models"
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// InitDB inicializa la base de datos SQLite y realiza las migraciones automáticas.
func InitDB(dbPath string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}

	// Forzar la activación de Foreign Keys en SQLite
	db.Exec("PRAGMA foreign_keys = ON;")

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
