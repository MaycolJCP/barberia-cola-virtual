package services

import (
	"barberia-cola-virtual/internal/models"
	"errors"
	"testing"
)

// MockTurnosRepositoryCompleto simula el repositorio con todas las funciones para TurnoService
type MockTurnosRepositoryCompleto struct {
	SimulateError bool
}

func (m *MockTurnosRepositoryCompleto) Create(t *models.Turno) error {
	if m.SimulateError {
		return errors.New("db error")
	}
	t.ID = 1
	return nil
}

func (m *MockTurnosRepositoryCompleto) GetAll() ([]models.Turno, error) {
	if m.SimulateError {
		return nil, errors.New("db error")
	}
	return []models.Turno{{ClienteID: 1, ServicioID: 2, Estado: "ESPERANDO"}}, nil
}

func (m *MockTurnosRepositoryCompleto) GetByID(id uint) (models.Turno, error) {
	if m.SimulateError || id == 999 {
		return models.Turno{}, errors.New("not found")
	}
	return models.Turno{ClienteID: 1, ServicioID: 2, Estado: "ESPERANDO"}, nil
}

func (m *MockTurnosRepositoryCompleto) Update(t *models.Turno) error {
	if m.SimulateError {
		return errors.New("db error")
	}
	return nil
}

func (m *MockTurnosRepositoryCompleto) Delete(id uint) error {
	if m.SimulateError {
		return errors.New("db error")
	}
	return nil
}

func (m *MockTurnosRepositoryCompleto) CreateSeguimiento(s *models.SeguimientoTurno) error {
	return nil
}
func (m *MockTurnosRepositoryCompleto) CreateNotificacion(n *models.Notificacion) error { return nil }
func (m *MockTurnosRepositoryCompleto) GetSeguimientos() ([]models.SeguimientoTurno, error) {
	return nil, nil
}
func (m *MockTurnosRepositoryCompleto) GetNotificaciones() ([]models.Notificacion, error) {
	return nil, nil
}

// ============================================================================
// SUITE DE 5 PRUEBAS UNITARIAS REQUERIDAS PARA EL MÓDULO DE TURNOS
// ============================================================================

// 1. Camino Feliz: Registro correcto de un turno con ID de cliente y servicio válidos
func TestCreateTurno_Valido(t *testing.T) {
	mockRepo := &MockTurnosRepositoryCompleto{SimulateError: false}
	service := NewTurnoService(mockRepo)

	turnoNuevo := models.Turno{
		ClienteID:  1,
		ServicioID: 2,
	}

	resultado, ok := service.CreateTurno(turnoNuevo)
	if !ok {
		t.Fatal("Se esperaba que el turno se creara correctamente")
	}

	if resultado.Estado != "ESPERANDO" {
		t.Errorf("Se esperaba el estado 'ESPERANDO', se obtuvo: %s", resultado.Estado)
	}
}

// 2. Validación de Regla de Negocio: Rechazo de turnos con IDs en cero o negativos
func TestCreateTurno_Invalido(t *testing.T) {
	mockRepo := &MockTurnosRepositoryCompleto{SimulateError: false}
	service := NewTurnoService(mockRepo)

	turnoInvalido := models.Turno{
		ClienteID:  0,
		ServicioID: 0, // CAMBIADO: 0 es un uint válido para Go, pero inválido para tu regla de negocio
	}

	_, ok := service.CreateTurno(turnoInvalido)
	if ok {
		t.Error("Se esperaba que fallara la creacion del turno debido a IDs invalidos")
	}
}

// 3. Error de Infraestructura: Error al intentar persistir el turno en la BD
func TestCreateTurno_ErrorBaseDatos(t *testing.T) {
	mockRepo := &MockTurnosRepositoryCompleto{SimulateError: true}
	service := NewTurnoService(mockRepo)

	turnoNuevo := models.Turno{
		ClienteID:  1,
		ServicioID: 2,
	}

	_, ok := service.CreateTurno(turnoNuevo)
	if ok {
		t.Error("Se esperaba un fallo (false) debido a problemas de escritura en el repositorio")
	}
}

// 4. Búsqueda Exitosa: Obtener un turno por su ID único
func TestGetTurnoByID_Exitoso(t *testing.T) {
	mockRepo := &MockTurnosRepositoryCompleto{SimulateError: false}
	service := NewTurnoService(mockRepo)

	_, ok := service.GetTurnoByID(1)
	if !ok {
		t.Error("Se esperaba encontrar el turno con un ID existente")
	}
}

// 5. Flujo Alterno de Actualización: Modificar de forma correcta un turno existente
func TestUpdateTurno_Exitoso(t *testing.T) {
	mockRepo := &MockTurnosRepositoryCompleto{SimulateError: false}
	service := NewTurnoService(mockRepo)

	turnoAEditar := models.Turno{
		ClienteID:  1,
		ServicioID: 2,
		Estado:     "ATENDIDO",
	}

	_, ok := service.UpdateTurno(turnoAEditar)
	if !ok {
		t.Error("Se esperaba que la actualizacion del turno fuera exitosa")
	}
}
