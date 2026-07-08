package repository

import (
	"testing"

	"barberia-cola-virtual/internal/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func TestTurnoGORM_ListarYCrear(t *testing.T) {
	// 1. Conexión en memoria
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error al abrir base de datos en memoria: %v", err)
	}

	// 2. Migración del modelo
	_ = db.AutoMigrate(&models.Turno{})

	// 3. Inicializar el repositorio real de producción
	repo := NewTurnosRepository(db)

	// 4. Insertar datos de prueba usando tu método Create
	turno := models.Turno{ClienteID: 2, ServicioID: 3, Estado: "ESPERANDO"}
	err = repo.Create(&turno)
	if err != nil {
		t.Fatalf("Error al crear el turno usando el repositorio: %v", err)
	}

	// 5. Verificar listado usando tu método GetAll
	lista, err := repo.GetAll()
	if err != nil {
		t.Fatalf("Error al obtener los turnos usando el repositorio: %v", err)
	}

	if len(lista) != 1 {
		t.Errorf("Se esperaba 1 turno en la lista, se obtuvieron %d", len(lista))
	}
}
