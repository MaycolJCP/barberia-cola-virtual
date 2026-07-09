package repository

import (
	"testing"

	"barberia-cola-virtual/internal/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)
//En Repository pruebo que la entidad Servicio realmente pueda guardarse y consultarse usando GORM. 
// Para la prueba se puede usar una base temporal en memoria, así no se toca la base real de PostgreSQL.”

func TestCatalogoRepository_CrearYBuscarReflejaEnBD(t *testing.T) {
	// 1. PREPARAR: Configurar conexión limpia a SQLite en memoria usando GORM
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("no se pudo conectar a la base de datos en memoria: %v", err)
	}

	// Correr AutoMigrate para crear las tablas reales en la BD efímera
	err = db.AutoMigrate(&models.Servicio{})
	if err != nil {
		t.Fatalf("error al realizar el AutoMigrate: %v", err)
	}

	// Instanciar el repositorio real pasándole la conexión de GORM
	// NOTA: Asegúrate de que el nombre de tu constructor real coincida (ej. NewCatalogoRepository o similar)
	repo := NewCatalogRepository(db)

	servicioOriginal := models.Servicio{
		CategoriaID: 1,
		Nombre:      "Corte Barba Real",
		Descripcion: "Afeitado clásico con toalla caliente",
		Precio:      8,
		Duracion:    25,
	}

	// 2. EJECUTAR: Guardar en la base de datos real
	err = repo.CreateServicio(&servicioOriginal)
	if err != nil {
		t.Fatalf("no se pudo guardar el servicio en el repositorio: %v", err)
	}

	// 3. VERIFICAR: Buscar/Listar de la base de datos para ver si se refleja
	servicioGuardado, err := repo.GetServicioByID(servicioOriginal.ID)
	if err != nil {
		t.Fatalf("error al buscar el servicio recién creado: %v", err)
	}

	// Comprobar que los campos clave sigan intactos
	if servicioGuardado.ID == 0 {
		t.Error("se esperaba que la base de datos asigne un ID autoincremental válido")
	}

	if servicioGuardado.Nombre != "Corte Barba Real" {
		t.Errorf("se esperaba el nombre 'Corte Barba Real', pero se obtuvo: %s", servicioGuardado.Nombre)
	}
}
