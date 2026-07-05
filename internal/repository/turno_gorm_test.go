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

	// 2. Migración
	_ = db.AutoMigrate(&models.Turno{})

	// 3. Insertar datos de prueba
	turno := models.Turno{ClienteID: 2, ServicioID: 3, Estado: "ESPERANDO"}
	db.Create(&turno)

	// 4. Verificar listado
	var lista []models.Turno
	db.Find(&lista)

	if len(lista) != 1 {
		t.Errorf("Se esperaba 1 turno en la lista, se obtuvieron %d", len(lista))
	}
}
