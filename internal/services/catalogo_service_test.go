package services

import (
	"barberia-cola-virtual/internal/models"
	"errors"
	"testing"
)

// MockCatalogRepository simula la capa Repository para probar el Service sin usar una base de datos real.

type MockCatalogRepository struct {
	SimulateError bool
}

// CreateServicio Simula la creación de un servicio. Si SimulateError está activo, devuelve un error; de lo contrario asigna un ID.
func (m *MockCatalogRepository) CreateServicio(s *models.Servicio) error {
	if m.SimulateError {
		return errors.New("db error")
	}
	s.ID = 1
	return nil
}

// GetServicios Simula la consulta de todos los servicios. Devuelve una lista de ejemplo o un error controlado.
func (m *MockCatalogRepository) GetServicios() ([]models.Servicio, error) {
	if m.SimulateError {
		return nil, errors.New("db error")
	}
	return []models.Servicio{{Nombre: "Corte Clasico", Precio: 5, Duracion: 30}}, nil
}

// GetServicioByID Simula la búsqueda de un servicio por ID. El ID 999 representa un registro inexistente.
func (m *MockCatalogRepository) GetServicioByID(id uint) (models.Servicio, error) {
	if m.SimulateError || id == 999 {
		return models.Servicio{}, errors.New("not found")
	}
	return models.Servicio{Nombre: "Corte Clasico", Precio: 5, Duracion: 30}, nil
}

// UpdateServicio Simula la actualización de un servicio y permite provocar un error del repositorio.
func (m *MockCatalogRepository) UpdateServicio(s *models.Servicio) error {
	if m.SimulateError {
		return errors.New("db error")
	}
	return nil
}

// DeleteServicio Simula la eliminación de un servicio por ID.
func (m *MockCatalogRepository) DeleteServicio(id uint) error {
	if m.SimulateError {
		return errors.New("db error")
	}
	return nil
}

// CreateCategoria Simula la creación de una categoría y le asigna un ID.
func (m *MockCatalogRepository) CreateCategoria(c *models.CategoriaServicio) error {
	if m.SimulateError {
		return errors.New("db error")
	}
	c.ID = 1
	return nil
}

// GetCategorias Simula la consulta de todas las categorías.
func (m *MockCatalogRepository) GetCategorias() ([]models.CategoriaServicio, error) {
	if m.SimulateError {
		return nil, errors.New("db error")
	}
	return []models.CategoriaServicio{{Nombre: "Cortes"}}, nil
}

// GetCategoriaByID Simula la búsqueda de una categoría por ID; el ID 999 representa que no existe.
func (m *MockCatalogRepository) GetCategoriaByID(id uint) (models.CategoriaServicio, error) {
	if m.SimulateError || id == 999 {
		return models.CategoriaServicio{}, errors.New("not found")
	}
	return models.CategoriaServicio{Nombre: "Cortes"}, nil
}

// UpdateCategoria Simula la actualización de una categoría.
func (m *MockCatalogRepository) UpdateCategoria(c *models.CategoriaServicio) error {
	if m.SimulateError {
		return errors.New("db error")
	}
	return nil
}

// DeleteCategoria Simula la eliminación de una categoría.
func (m *MockCatalogRepository) DeleteCategoria(id uint) error {
	if m.SimulateError {
		return errors.New("db error")
	}
	return nil
}

// CreatePromocion Simula la creación de una promoción y le asigna un ID.
func (m *MockCatalogRepository) CreatePromocion(p *models.Promocion) error {
	if m.SimulateError {
		return errors.New("db error")
	}
	p.ID = 1
	return nil
}

// GetPromociones Simula la consulta de todas las promociones.
func (m *MockCatalogRepository) GetPromociones() ([]models.Promocion, error) {
	if m.SimulateError {
		return nil, errors.New("db error")
	}
	return []models.Promocion{{Nombre: "Promo Corte", Descuento: 10}}, nil
}

// GetPromocionByID Simula la búsqueda de una promoción por ID; el ID 999 representa que no existe.
func (m *MockCatalogRepository) GetPromocionByID(id uint) (models.Promocion, error) {
	if m.SimulateError || id == 999 {
		return models.Promocion{}, errors.New("not found")
	}
	return models.Promocion{Nombre: "Promo Corte", Descuento: 10}, nil
}

// UpdatePromocion Simula la actualización de una promoción.
func (m *MockCatalogRepository) UpdatePromocion(p *models.Promocion) error {
	if m.SimulateError {
		return errors.New("db error")
	}
	return nil
}

// DeletePromocion Simula la eliminación de una promoción.
func (m *MockCatalogRepository) DeletePromocion(id uint) error {
	if m.SimulateError {
		return errors.New("db error")
	}
	return nil
}

// ================= SERVICIOS =================

// TestCreateServicio_Valido Comprueba que un servicio con nombre, precio y duración válidos pueda crearse.
func TestCreateServicio_Valido(t *testing.T) {
	service := NewCatalogService(&MockCatalogRepository{})

	servicio := models.Servicio{
		Nombre:   "Corte Clasico",
		Precio:   5,
		Duracion: 30,
	}

	_, ok := service.CreateServicio(servicio)
	if !ok {
		t.Error("se esperaba crear el servicio correctamente")
	}
}

// TestCreateServicio_Invalido Comprueba que el Service rechace un servicio con datos inválidos.
func TestCreateServicio_Invalido(t *testing.T) {
	service := NewCatalogService(&MockCatalogRepository{})

	servicio := models.Servicio{
		Nombre:   "",
		Precio:   0,
		Duracion: 0,
	}

	_, ok := service.CreateServicio(servicio)
	if ok {
		t.Error("se esperaba error por datos invalidos")
	}
}

// TestCreateServicio_ErrorRepositorio Comprueba que el Service maneje correctamente un error del repositorio al crear.
func TestCreateServicio_ErrorRepositorio(t *testing.T) {
	service := NewCatalogService(&MockCatalogRepository{SimulateError: true})

	servicio := models.Servicio{
		Nombre:   "Corte",
		Precio:   5,
		Duracion: 30,
	}

	_, ok := service.CreateServicio(servicio)
	if ok {
		t.Error("se esperaba error del repositorio")
	}
}

// TestGetServicios_Exitoso Comprueba que se pueda obtener una lista de servicios.
func TestGetServicios_Exitoso(t *testing.T) {
	service := NewCatalogService(&MockCatalogRepository{})

	servicios, err := service.GetServicios()
	if err != nil {
		t.Errorf("no se esperaba error: %v", err)
	}
	if len(servicios) == 0 {
		t.Error("se esperaba al menos un servicio")
	}
}

// TestGetServicios_Error Comprueba que se propague el error cuando falla la consulta de servicios.
func TestGetServicios_Error(t *testing.T) {
	service := NewCatalogService(&MockCatalogRepository{SimulateError: true})

	_, err := service.GetServicios()
	if err == nil {
		t.Error("se esperaba error al listar servicios")
	}
}

// TestGetServicioByID_Exitoso Comprueba que un servicio existente pueda encontrarse por su ID.
func TestGetServicioByID_Exitoso(t *testing.T) {
	service := NewCatalogService(&MockCatalogRepository{})

	servicio, ok := service.GetServicioByID(1)
	if !ok {
		t.Fatal("se esperaba encontrar el servicio")
	}
	if servicio.Nombre != "Corte Clasico" {
		t.Errorf("servicio incorrecto: %s", servicio.Nombre)
	}
}

// TestGetServicioByID_NoEncontrado Comprueba el comportamiento cuando el servicio solicitado no existe.
func TestGetServicioByID_NoEncontrado(t *testing.T) {
	service := NewCatalogService(&MockCatalogRepository{})

	_, ok := service.GetServicioByID(999)
	if ok {
		t.Error("se esperaba que no encuentre el servicio")
	}
}

// TestUpdateServicio_Exitoso Comprueba que un servicio válido pueda actualizarse.
func TestUpdateServicio_Exitoso(t *testing.T) {
	service := NewCatalogService(&MockCatalogRepository{})

	servicio := models.Servicio{
		Nombre:   "Corte Actualizado",
		Precio:   6,
		Duracion: 35,
	}

	_, ok := service.UpdateServicio(servicio)
	if !ok {
		t.Error("se esperaba actualizar el servicio")
	}
}

// TestUpdateServicio_Error Comprueba que se maneje un error del repositorio al actualizar.
func TestUpdateServicio_Error(t *testing.T) {
	service := NewCatalogService(&MockCatalogRepository{SimulateError: true})

	_, ok := service.UpdateServicio(models.Servicio{})
	if ok {
		t.Error("se esperaba error al actualizar")
	}
}

// TestDeleteServicio_Exitoso Comprueba que un servicio pueda eliminarse correctamente.
func TestDeleteServicio_Exitoso(t *testing.T) {
	service := NewCatalogService(&MockCatalogRepository{})

	ok := service.DeleteServicio(1)
	if !ok {
		t.Error("se esperaba eliminar el servicio")
	}
}

// TestDeleteServicio_Error Comprueba que se maneje un error del repositorio al eliminar.
func TestDeleteServicio_Error(t *testing.T) {
	service := NewCatalogService(&MockCatalogRepository{SimulateError: true})

	ok := service.DeleteServicio(1)
	if ok {
		t.Error("se esperaba error al eliminar")
	}
}

// ================= CATEGORÍAS =================

// TestCreateCategoria_Valida Comprueba que una categoría con nombre válido pueda crearse.
func TestCreateCategoria_Valida(t *testing.T) {
	service := NewCatalogService(&MockCatalogRepository{})

	categoria := models.CategoriaServicio{
		Nombre: "Cortes",
	}

	_, ok := service.CreateCategoriaServicio(categoria)
	if !ok {
		t.Error("se esperaba crear la categoria")
	}
}

// TestCreateCategoria_Invalida Comprueba que una categoría sin nombre sea rechazada.
func TestCreateCategoria_Invalida(t *testing.T) {
	service := NewCatalogService(&MockCatalogRepository{})

	_, ok := service.CreateCategoriaServicio(models.CategoriaServicio{})
	if ok {
		t.Error("se esperaba error por categoria sin nombre")
	}
}

// TestGetCategorias_Exitoso Comprueba que se pueda obtener una lista de categorías.
func TestGetCategorias_Exitoso(t *testing.T) {
	service := NewCatalogService(&MockCatalogRepository{})

	categorias, err := service.GetCategoriasServicio()
	if err != nil {
		t.Errorf("no se esperaba error: %v", err)
	}
	if len(categorias) == 0 {
		t.Error("se esperaba al menos una categoria")
	}
}

// TestGetCategoriaByID_Exitoso Comprueba que una categoría existente pueda encontrarse por ID.
func TestGetCategoriaByID_Exitoso(t *testing.T) {
	service := NewCatalogService(&MockCatalogRepository{})

	_, ok := service.GetCategoriaServicioByID(1)
	if !ok {
		t.Error("se esperaba encontrar la categoria")
	}
}

// TestGetCategoriaByID_NoEncontrada Comprueba el comportamiento cuando la categoría no existe.
func TestGetCategoriaByID_NoEncontrada(t *testing.T) {
	service := NewCatalogService(&MockCatalogRepository{})

	_, ok := service.GetCategoriaServicioByID(999)
	if ok {
		t.Error("se esperaba que no encuentre la categoria")
	}
}

// TestUpdateCategoria_Exitoso Comprueba que una categoría pueda actualizarse.
func TestUpdateCategoria_Exitoso(t *testing.T) {
	service := NewCatalogService(&MockCatalogRepository{})

	_, ok := service.UpdateCategoriaServicio(models.CategoriaServicio{Nombre: "Cortes Premium"})
	if !ok {
		t.Error("se esperaba actualizar la categoria")
	}
}

// TestDeleteCategoria_Exitoso Comprueba que una categoría pueda eliminarse.
func TestDeleteCategoria_Exitoso(t *testing.T) {
	service := NewCatalogService(&MockCatalogRepository{})

	ok := service.DeleteCategoriaServicio(1)
	if !ok {
		t.Error("se esperaba eliminar la categoria")
	}
}

// ================= PROMOCIONES =================

// TestCreatePromocion_Valida Comprueba que una promoción válida pueda crearse.
func TestCreatePromocion_Valida(t *testing.T) {
	service := NewCatalogService(&MockCatalogRepository{})

	promo := models.Promocion{
		Nombre:    "Promo Corte",
		Descuento: 10,
	}

	_, ok := service.CreatePromocion(promo)
	if !ok {
		t.Error("se esperaba crear la promocion")
	}
}

// TestCreatePromocion_Invalida Comprueba que una promoción sin nombre o descuento válido sea rechazada.
func TestCreatePromocion_Invalida(t *testing.T) {
	service := NewCatalogService(&MockCatalogRepository{})

	promo := models.Promocion{
		Nombre:    "",
		Descuento: 0,
	}

	_, ok := service.CreatePromocion(promo)
	if ok {
		t.Error("se esperaba error por promocion invalida")
	}
}

// TestGetPromociones_Exitoso Comprueba que se pueda obtener una lista de promociones.
func TestGetPromociones_Exitoso(t *testing.T) {
	service := NewCatalogService(&MockCatalogRepository{})

	promos, err := service.GetPromociones()
	if err != nil {
		t.Errorf("no se esperaba error: %v", err)
	}
	if len(promos) == 0 {
		t.Error("se esperaba al menos una promocion")
	}
}

// TestGetPromocionByID_Exitoso Comprueba que una promoción existente pueda encontrarse por ID.
func TestGetPromocionByID_Exitoso(t *testing.T) {
	service := NewCatalogService(&MockCatalogRepository{})

	_, ok := service.GetPromocionByID(1)
	if !ok {
		t.Error("se esperaba encontrar la promocion")
	}
}

// TestGetPromocionByID_NoEncontrada Comprueba el comportamiento cuando la promoción no existe.
func TestGetPromocionByID_NoEncontrada(t *testing.T) {
	service := NewCatalogService(&MockCatalogRepository{})

	_, ok := service.GetPromocionByID(999)
	if ok {
		t.Error("se esperaba que no encuentre la promocion")
	}
}

// TestUpdatePromocion_Exitoso Comprueba que una promoción pueda actualizarse.
func TestUpdatePromocion_Exitoso(t *testing.T) {
	service := NewCatalogService(&MockCatalogRepository{})

	_, ok := service.UpdatePromocion(models.Promocion{Nombre: "Promo Actualizada", Descuento: 15})
	if !ok {
		t.Error("se esperaba actualizar la promocion")
	}
}

// TestDeletePromocion_Exitoso Comprueba que una promoción pueda eliminarse.
func TestDeletePromocion_Exitoso(t *testing.T) {
	service := NewCatalogService(&MockCatalogRepository{})

	ok := service.DeletePromocion(1)
	if !ok {
		t.Error("se esperaba eliminar la promocion")
	}
}