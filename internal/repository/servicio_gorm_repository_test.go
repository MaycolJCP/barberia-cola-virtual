package repository

import (
	"testing"

	"barberia-cola-virtual/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// =======================================
// TEST REPOSITORY GORM - CREAR Y LISTAR SERVICIO
//
// PREPARAR:
// Crear una base SQLite en memoria.
// Ejecutar AutoMigrate para crear la tabla Servicio.
// Crear el repository real con GORM.
//
// EJECUTAR:
// Guardar un servicio.
// Listar los servicios.
//
// VERIFICAR:
// Debe existir un servicio guardado en la base.
// =======================================
func TestServicioGormRepository_CrearYListar(t *testing.T) {

	// PREPARAR: base de datos desechable en memoria
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
	t.Fatalf("no se pudo abrir sqlite en memoria: %v", err)
}

	// Crear tabla desde el struct Servicio
	err = db.AutoMigrate(&models.Servicio{})
	if err != nil {
		t.Fatal("no se pudo ejecutar automigrate")
	}

	repo := ServicioGormRepository{
		DB: db,
	}

	servicio := models.Servicio{
		CategoriaID: 1,
		Nombre:      "Corte Clasico",
		Descripcion: "Corte tradicional",
		Precio:      5,
		Duracion:    30,
	}

	// EJECUTAR: guardar servicio
	err = repo.CrearServicio(servicio)
	if err != nil {
		t.Fatal("no se pudo guardar el servicio")
	}

	// EJECUTAR: listar servicios
	servicios, err := repo.ObtenerServicios()
	if err != nil {
		t.Fatal("no se pudieron listar los servicios")
	}

	// VERIFICAR
	if len(servicios) != 1 {
		t.Errorf("se esperaba 1 servicio, se obtuvo %d", len(servicios))
	}

	if servicios[0].Nombre != "Corte Clasico" {
		t.Error("el nombre del servicio no coincide")
	}
}