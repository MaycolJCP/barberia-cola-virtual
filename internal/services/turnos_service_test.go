package services

import (
	"testing"

	"barberia-cola-virtual/internal/models"
	"barberia-cola-virtual/internal/storage"
)

// =======================================
// TEST CREAR TURNO VALIDO
//
// PREPARAR:
// Crear un turno con cliente_id y servicio_id validos.
//
// EJECUTAR:
// Llamar a CreateTurno().
//
// VERIFICAR:
// Debe crear el turno correctamente.
// Tambien debe generar seguimiento y notificacion.
// =======================================
func TestCreateTurno_Valido(t *testing.T) {

	// Limpiar memoria antes del test
	storage.Turnos = []models.Turno{}
	storage.SeguimientosTurno = []models.SeguimientoTurno{}
	storage.Notificaciones = []models.Notificacion{}

	// Turno valido
	turno := models.Turno{
		ClienteID:  1,
		ServicioID: 1,
	}

	// Ejecutar
	turnoCreado, ok := CreateTurno(turno)

	// Verificar turno
	if !ok {
		t.Error("se esperaba crear el turno correctamente")
	}

	if turnoCreado.ID == 0 {
		t.Error("se esperaba que el turno tenga ID")
	}

	if turnoCreado.Estado != "ESPERANDO" {
		t.Error("se esperaba estado ESPERANDO")
	}

	// Verificar seguimiento automatico
	if len(storage.SeguimientosTurno) != 1 {
		t.Error("se esperaba crear un seguimiento automaticamente")
	}

	// Verificar notificacion automatica
	if len(storage.Notificaciones) != 1 {
		t.Error("se esperaba crear una notificacion automaticamente")
	}
}

// =======================================
// TEST CREAR TURNO INVALIDO
//
// PREPARAR:
// Crear un turno sin cliente_id ni servicio_id.
//
// EJECUTAR:
// Llamar a CreateTurno().
//
// VERIFICAR:
// Debe fallar porque los datos obligatorios
// no fueron enviados.
// =======================================
func TestCreateTurno_Invalido(t *testing.T) {

	// Limpiar memoria
	storage.Turnos = []models.Turno{}
	storage.SeguimientosTurno = []models.SeguimientoTurno{}
	storage.Notificaciones = []models.Notificacion{}

	// Turno invalido
	turno := models.Turno{
		ClienteID:  0,
		ServicioID: 0,
	}

	// Ejecutar
	turnoCreado, ok := CreateTurno(turno)

	// Verificar
	if ok {
		t.Error("se esperaba que falle")
	}

	if turnoCreado.ID != 0 {
		t.Error("no deberia crear turno invalido")
	}

	if len(storage.SeguimientosTurno) != 0 {
		t.Error("no deberia crear seguimiento")
	}

	if len(storage.Notificaciones) != 0 {
		t.Error("no deberia crear notificacion")
	}
}