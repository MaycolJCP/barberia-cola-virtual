package services

import (
	"barberia-cola-virtual/internal/models"
	"errors"
	"testing"

	"gorm.io/gorm" // Necesario para instanciar gorm.Model
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

// SOLUCIONADO: Inicialización correcta de structs embebidos de GORM usando gorm.Model
func (m *MockTurnosRepositoryCompleto) GetAll() ([]models.Turno, error) {
	if m.SimulateError {
		return nil, errors.New("db error")
	}
	return []models.Turno{
		{
			Model:      gorm.Model{ID: 1}, // Solución al error ID heredado
			ClienteID:  1,
			ServicioID: 2,
			Estado:     "ESPERANDO",
			Servicio: &models.Servicio{
				Model:    gorm.Model{ID: 2}, // Solución al error ID heredado
				Nombre:   "Corte Clasico",
				Precio:   10.00,
				Duracion: 20,
			},
		},
	}, nil
}

func (m *MockTurnosRepositoryCompleto) GetByID(id uint) (models.Turno, error) {
	if m.SimulateError || id == 999 {
		return models.Turno{}, errors.New("not found")
	}
	return models.Turno{
		ClienteID:  1,
		ServicioID: 2,
		Estado:     "ESPERANDO",
		Servicio: &models.Servicio{
			Model:    gorm.Model{ID: 2}, // Solución al error ID heredado
			Nombre:   "Corte Clasico",
			Duracion: 20,
		},
	}, nil
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
	if m.SimulateError {
		return errors.New("db error")
	}
	return nil
}

func (m *MockTurnosRepositoryCompleto) CreateNotificacion(n *models.Notificacion) error {
	if m.SimulateError {
		return errors.New("db error")
	}
	return nil
}

func (m *MockTurnosRepositoryCompleto) GetSeguimientos() ([]models.SeguimientoTurno, error) {
	if m.SimulateError {
		return nil, errors.New("db error")
	}
	return []models.SeguimientoTurno{{TurnoID: 1, Posicion: 1}}, nil
}

func (m *MockTurnosRepositoryCompleto) GetNotificaciones() ([]models.Notificacion, error) {
	if m.SimulateError {
		return nil, errors.New("db error")
	}
	return []models.Notificacion{{TurnoID: 1, Mensaje: "Tu turno se acerca"}}, nil
}

// ============================================================================
// SUITE DE 10 PRUEBAS UNITARIAS REQUERIDAS (RÚBRICA HITO 3 - COBERTURA >= 50%)
// ============================================================================

// 1. Camino Feliz: Registro correcto de un turno con IDs válidos
func TestCreateTurno_Valido(t *testing.T) {
	mockRepo := &MockTurnosRepositoryCompleto{SimulateError: false}
	service := NewTurnoService(mockRepo)

	turnoNuevo := models.Turno{
		Model:      gorm.Model{ID: 2}, // Solución al error ID heredado
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

// 2. Validación de Regla de Negocio: Rechazo de turnos con IDs en cero
func TestCreateTurno_Invalido(t *testing.T) {
	mockRepo := &MockTurnosRepositoryCompleto{SimulateError: false}
	service := NewTurnoService(mockRepo)

	turnoInvalido := models.Turno{
		ClienteID:  0,
		ServicioID: 0,
	}

	_, ok := service.CreateTurno(turnoInvalido)
	if ok {
		t.Error("Se esperaba que fallara la creacion del turno debido a IDs invalidos")
	}
}

// 3. Error de Infraestructura: Fallo de escritura en la persistencia del Repositorio
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

// 5. Flujo Alterno de Actualización: Modificar un turno existente de forma correcta
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

// 6. Búsqueda Fallida (Edge Case): Buscar un turno que no existe en el sistema
func TestGetTurnoByID_NoEncontrado(t *testing.T) {
	mockRepo := &MockTurnosRepositoryCompleto{SimulateError: false}
	service := NewTurnoService(mockRepo)

	_, ok := service.GetTurnoByID(999)
	if ok {
		t.Error("Se esperaba un fallo (false) al buscar un turno inexistente")
	}
}

// 7. Listar todos los Turnos de la Cola Virtual
func TestGetTurnos_Exitoso(t *testing.T) {
	mockRepo := &MockTurnosRepositoryCompleto{SimulateError: false}
	service := NewTurnoService(mockRepo)

	lista, err := service.GetTurnos()
	if err != nil {
		t.Errorf("No se esperaba un error al listar turnos: %v", err)
	}
	if len(lista) == 0 {
		t.Error("La lista devuelta no debería estar vacía")
	}
}

// 8. Eliminación Correcta de un Turno por su identificador
func TestDeleteTurno_Exitoso(t *testing.T) {
	mockRepo := &MockTurnosRepositoryCompleto{SimulateError: false}
	service := NewTurnoService(mockRepo)

	ok := service.DeleteTurno(1)
	if !ok {
		t.Error("Se esperaba una eliminación exitosa del turno")
	}
}

// 9. Obtención del estado de los Seguimientos
func TestGetSeguimientos_Exitoso(t *testing.T) {
	mockRepo := &MockTurnosRepositoryCompleto{SimulateError: false}
	service := NewTurnoService(mockRepo)

	lista, err := service.GetSeguimientosTurno()
	if err != nil {
		t.Errorf("No se esperaba un error al obtener seguimientos: %v", err)
	}
	if len(lista) == 0 {
		t.Error("La lista de seguimientos no debería estar vacía")
	}
}

// 10. Consulta de las Alertas o Notificaciones registradas
func TestGetNotificaciones_Exitoso(t *testing.T) {
	mockRepo := &MockTurnosRepositoryCompleto{SimulateError: false}
	service := NewTurnoService(mockRepo)

	lista, err := service.GetNotificaciones()
	if err != nil {
		t.Errorf("No se esperaba un error al obtener notificaciones: %v", err)
	}
	if len(lista) == 0 {
		t.Error("La lista de notificaciones no debería estar vacía")
	}
}
