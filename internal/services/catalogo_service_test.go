package services

import (
	"barberia-cola-virtual/internal/models"
	"errors"
	"testing"
)

// MockCatalogRepository implementa de forma simulada todos los métodos requeridos por repository.CatalogRepository
type MockCatalogRepository struct {
	SimulateError bool
}

// Métodos de Servicios
func (m *MockCatalogRepository) CreateServicio(s *models.Servicio) error {
	if m.SimulateError {
		return errors.New("db error")
	}
	s.ID = 1
	return nil
}
func (m *MockCatalogRepository) GetServicios() ([]models.Servicio, error) { return nil, nil }
func (m *MockCatalogRepository) GetServicioByID(id uint) (models.Servicio, error) {
	if m.SimulateError {
		return models.Servicio{}, errors.New("not found")
	}
	return models.Servicio{Nombre: "Corte Clasico"}, nil
}
func (m *MockCatalogRepository) UpdateServicio(s *models.Servicio) error { return nil }
func (m *MockCatalogRepository) DeleteServicio(id uint) error            { return nil }

// Métodos de Categorías
func (m *MockCatalogRepository) CreateCategoria(c *models.CategoriaServicio) error  { return nil }
func (m *MockCatalogRepository) GetCategorias() ([]models.CategoriaServicio, error) { return nil, nil }
func (m *MockCatalogRepository) GetCategoriaByID(id uint) (models.CategoriaServicio, error) {
	return models.CategoriaServicio{}, nil
}
func (m *MockCatalogRepository) UpdateCategoria(c *models.CategoriaServicio) error { return nil }
func (m *MockCatalogRepository) DeleteCategoria(id uint) error                     { return nil }

// Métodos de Promociones
func (m *MockCatalogRepository) CreatePromocion(p *models.Promocion) error   { return nil }
func (m *MockCatalogRepository) GetPromociones() ([]models.Promocion, error) { return nil, nil }
func (m *MockCatalogRepository) GetPromocionByID(id uint) (models.Promocion, error) {
	return models.Promocion{}, nil
}
func (m *MockCatalogRepository) UpdatePromocion(p *models.Promocion) error { return nil }
func (m *MockCatalogRepository) DeletePromocion(id uint) error             { return nil }

// ============================================================================
// SUITE DE 5 PRUEBAS REQUERIDAS PARA EL MÓDULO DE CATÁLOGO
// ============================================================================

// 1. Test de Camino Feliz: Creación de servicio con datos válidos
func TestCreateServicio_Valido(t *testing.T) {
	mockRepo := &MockCatalogRepository{SimulateError: false}
	service := NewCatalogService(mockRepo)

	nuevoServicio := models.Servicio{
		Nombre:   "Corte Degradado + Barba",
		Precio:   12.50,
		Duracion: 30,
	}

	_, ok := service.CreateServicio(nuevoServicio)
	if !ok {
		t.Error("Se esperaba que el servicio se creara correctamente con datos validos")
	}
}

// 2. Test de Validación de Regla de Negocio: Campos obligatorios vacíos o inválidos
func TestCreateServicio_Invalido_CamposVacios(t *testing.T) {
	mockRepo := &MockCatalogRepository{SimulateError: false}
	service := NewCatalogService(mockRepo)

	// Nombre vacío, precio y duración inválidos (cero o negativos)
	servicioInvalido := models.Servicio{
		Nombre:   "",
		Precio:   0,
		Duracion: -5,
	}

	_, ok := service.CreateServicio(servicioInvalido)
	if ok {
		t.Error("Se esperaba que fallara la creacion debido a campos obligatorios invalidos")
	}
}

// 3. Test de Error de Infraestructura: Simular fallo al guardar en la base de datos
func TestCreateServicio_ErrorBaseDatos(t *testing.T) {
	// Forzamos al repositorio simulado a retornar un error interno de base de datos
	mockRepo := &MockCatalogRepository{SimulateError: true}
	service := NewCatalogService(mockRepo)

	nuevoServicio := models.Servicio{
		Nombre:   "Tinte de Cabello",
		Precio:   20.00,
		Duracion: 45,
	}

	_, ok := service.CreateServicio(nuevoServicio)
	if ok {
		t.Error("Se esperaba que retornara false debido a un error de escritura en el repositorio")
	}
}

// 4. Test de Búsqueda Exitosa: Obtener un servicio existente por su ID único
func TestGetServicioByID_Exitoso(t *testing.T) {
	mockRepo := &MockCatalogRepository{SimulateError: false}
	service := NewCatalogService(mockRepo)

	servicio, ok := service.GetServicioByID(1)
	if !ok {
		t.Fatal("Se esperaba encontrar el servicio buscado")
	}

	if servicio.Nombre != "Corte Clasico" {
		t.Errorf("Se esperaba el servicio 'Corte Clasico', se obtuvo: %s", servicio.Nombre)
	}
}

// 5. Test de Búsqueda Fallida: Intentar obtener un servicio que no existe en el almacenamiento
func TestGetServicioByID_NoEncontrado(t *testing.T) {
	mockRepo := &MockCatalogRepository{SimulateError: true}
	service := NewCatalogService(mockRepo)

	_, ok := service.GetServicioByID(999)
	if ok {
		t.Error("Se esperaba que retornara false (no encontrado) al buscar un ID inexistente")
	}
}
