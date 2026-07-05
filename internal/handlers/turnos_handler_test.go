package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"barberia-cola-virtual/internal/models"
	"barberia-cola-virtual/internal/services"
)

type MockTurnoRepository struct{}

func (m *MockTurnoRepository) Create(t *models.Turno) error {
	t.ID = 1
	return nil
}
func (m *MockTurnoRepository) GetAll() ([]models.Turno, error)       { return nil, nil }
func (m *MockTurnoRepository) GetByID(id uint) (models.Turno, error) { return models.Turno{}, nil }
func (m *MockTurnoRepository) Update(t *models.Turno) error          { return nil }
func (m *MockTurnoRepository) Delete(id uint) error                  { return nil }

func (m *MockTurnoRepository) CreateSeguimiento(s *models.SeguimientoTurno) error  { return nil }
func (m *MockTurnoRepository) CreateNotificacion(n *models.Notificacion) error     { return nil }
func (m *MockTurnoRepository) GetSeguimientos() ([]models.SeguimientoTurno, error) { return nil, nil }
func (m *MockTurnoRepository) GetNotificaciones() ([]models.Notificacion, error)   { return nil, nil }

func TestCreateTurnoHandler_Valido(t *testing.T) {
	body := []byte(`{
		"cliente_id": 1,
		"servicio_id": 1
	}`)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/turnos", bytes.NewBuffer(body))
	rec := httptest.NewRecorder()

	mockRepo := &MockTurnoRepository{}
	turnoService := services.NewTurnoService(mockRepo)
	handler := NewTurnosHandler(turnoService)

	handler.CreateTurno(rec, req)

	if rec.Code != http.StatusCreated && rec.Code != http.StatusOK {
		t.Errorf("se esperaba status exitoso pero llego %d", rec.Code)
	}
}

func TestCreateTurnoHandler_Invalido(t *testing.T) {
	body := []byte(`{
		"cliente_id": 0,
		"servicio_id": 0
	}`)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/turnos", bytes.NewBuffer(body))
	rec := httptest.NewRecorder()

	mockRepo := &MockTurnoRepository{}
	turnoService := services.NewTurnoService(mockRepo)
	handler := NewTurnosHandler(turnoService)

	handler.CreateTurno(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("se esperaba status %d pero llego %d", http.StatusBadRequest, rec.Code)
	}
}

func TestCreateTurno_Handler401(t *testing.T) {
	w := httptest.NewRecorder()
	w.WriteHeader(http.StatusUnauthorized)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Se esperaba 401 Unauthorized, se obtuvo %d", w.Code)
	}
}
