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

	// 3. Inicializar el repositorio real de producción
	repo := NewTurnosRepository(db)

	// 4. Crear registro usando tu método Create (Quitamos el .Error del final)
	turnoNuevo := models.Turno{
		ClienteID:  1,
		ServicioID: 2,
		Estado:     "ESPERANDO",
	}

	if err := repo.Create(&turnoNuevo); err != nil {
		t.Fatalf("Error al guardar el turno usando el repositorio: %v", err)
	}

	// 5. Buscar el registro usando tu método GetByID (que incluye Preload)
	turnoResultado, err := repo.GetByID(turnoNuevo.ID)
	if err != nil {
		t.Fatalf("Error al buscar el turno guardado usando el repositorio: %v", err)
	}

	if turnoResultado.ClienteID != turnoNuevo.ClienteID {
		t.Errorf("Se esperaba ClienteID %d, se obtuvo %d", turnoNuevo.ClienteID, turnoResultado.ClienteID)
	}
}
