package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"barberia-cola-virtual/internal/models"
	"barberia-cola-virtual/internal/storage"
)

// =======================================
// TEST HANDLER CREAR TURNO VALIDO
//
// PREPARAR:
// Crear un JSON valido con cliente_id y servicio_id.
//
// EJECUTAR:
// Enviar la peticion falsa al handler CreateTurno().
//
// VERIFICAR:
// Debe responder HTTP 201 Created.
// Tambien debe guardar un turno en memoria.
// =======================================
func TestCreateTurnoHandler_Valido(t *testing.T) {

	// Limpiar memoria antes del test
	storage.Turnos = []models.Turno{}
	storage.SeguimientosTurno = []models.SeguimientoTurno{}
	storage.Notificaciones = []models.Notificacion{}

	// JSON que simula el body enviado desde Postman
	body := []byte(`{
		"cliente_id": 1,
		"servicio_id": 1
	}`)

	// Crear request falso
	req := httptest.NewRequest(http.MethodPost, "/api/v1/turnos", bytes.NewBuffer(body))

	// Recorder guarda la respuesta del handler
	rec := httptest.NewRecorder()

	// Ejecutar handler
	CreateTurno(rec, req)

	// Verificar codigo HTTP
	if rec.Code != http.StatusCreated {
		t.Errorf("se esperaba status %d pero llego %d", http.StatusCreated, rec.Code)
	}

	// Verificar que se guardo el turno
	if len(storage.Turnos) != 1 {
		t.Error("se esperaba un turno guardado en memoria")
	}

	// Verificar que se creo seguimiento
	if len(storage.SeguimientosTurno) != 1 {
		t.Error("se esperaba un seguimiento creado")
	}

	// Verificar que se creo notificacion
	if len(storage.Notificaciones) != 1 {
		t.Error("se esperaba una notificacion creada")
	}
}

// =======================================
// TEST HANDLER CREAR TURNO INVALIDO
//
// PREPARAR:
// Crear un JSON sin cliente_id ni servicio_id.
//
// EJECUTAR:
// Enviar la peticion falsa al handler CreateTurno().
//
// VERIFICAR:
// Debe responder HTTP 400 Bad Request.
// No debe guardar nada en memoria.
// =======================================
func TestCreateTurnoHandler_Invalido(t *testing.T) {

	// Limpiar memoria
	storage.Turnos = []models.Turno{}
	storage.SeguimientosTurno = []models.SeguimientoTurno{}
	storage.Notificaciones = []models.Notificacion{}

	// JSON invalido
	body := []byte(`{
		"cliente_id": 0,
		"servicio_id": 0
	}`)

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/turnos",
		bytes.NewBuffer(body),
	)

	rec := httptest.NewRecorder()

	// Ejecutar handler
	CreateTurno(rec, req)

	// Verificar status
	if rec.Code != http.StatusBadRequest {
		t.Errorf(
			"se esperaba status %d pero llego %d",
			http.StatusBadRequest,
			rec.Code,
		)
	}

	// Verificar que no guardo nada
	if len(storage.Turnos) != 0 {
		t.Error("no deberia guardar turnos")
	}

	if len(storage.SeguimientosTurno) != 0 {
		t.Error("no deberia crear seguimiento")
	}

	if len(storage.Notificaciones) != 0 {
		t.Error("no deberia crear notificacion")
	}
}