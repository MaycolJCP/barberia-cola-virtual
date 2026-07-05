package repository

import (
	"testing"

	"barberia-cola-virtual/internal/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func TestTurnoRepository_Persistencia(t *testing.T) {
	// 1. Conexión aislada en memoria
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("No se pudo conectar a la base de datos de prueba: %v", err)
	}

	// 2. AutoMigrate
	err = db.AutoMigrate(&models.Turno{})
	if err != nil {
		t.Fatalf("Error al migrar el modelo Turno: %v", err)
	}

	// 3. Crear registro
	turnoNuevo := models.Turno{
		ClienteID:  1,
		ServicioID: 2,
		Estado:     "ESPERANDO",
	}

	if err := db.Create(&turnoNuevo).Error; err != nil {
		t.Fatalf("Error al guardar el turno: %v", err)
	}

	// 4. Buscar / Listar lo refleja
	var turnoResultado models.Turno
	if err := db.First(&turnoResultado, turnoNuevo.ID).Error; err != nil {
		t.Fatalf("Error al buscar el turno guardado: %v", err)
	}

	if turnoResultado.ClienteID != turnoNuevo.ClienteID {
		t.Errorf("Se esperaba ClienteID %d, se obtuvo %d", turnoNuevo.ClienteID, turnoResultado.ClienteID)
	}
}
