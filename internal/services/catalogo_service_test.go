package services

import (
	"barberia-cola-virtual/internal/models"
	"errors"
	"testing"
)

//En esta capa pruebo las reglas de negocio del servicio. 
// Por ejemplo, que no se pueda crear un servicio con nombre vacío, precio cero o duración inválida. 
// También pruebo el caso exitoso y errores simulados del repositorio

type MockCatalogRepository struct {
	SimulateError bool
}

func (m *MockCatalogRepository) CreateServicio(s *models.Servicio) error {
	if m.SimulateError {
		return errors.New("db error")
	}
	s.ID = 1
	return nil
}

func (m *MockCatalogRepository) GetServicios() ([]models.Servicio, error) {
	if m.SimulateError {
		return nil, errors.New("db error")
	}
	return []models.Servicio{{Nombre: "Corte Clasico", Precio: 5, Duracion: 30}}, nil
}

func (m *MockCatalogRepository) GetServicioByID(id uint) (models.Servicio, error) {
	if m.SimulateError || id == 999 {
		return models.Servicio{}, errors.New("not found")
	}
	return models.Servicio{Nombre: "Corte Clasico", Precio: 5, Duracion: 30}, nil
}

func (m *MockCatalogRepository) UpdateServicio(s *models.Servicio) error {
	if m.SimulateError {
		return errors.New("db error")
	}
	return nil
}

func (m *MockCatalogRepository) DeleteServicio(id uint) error {
	if m.SimulateError {
		return errors.New("db error")
	}
	return nil
}

func (m *MockCatalogRepository) CreateCategoria(c *models.CategoriaServicio) error {
	if m.SimulateError {
		return errors.New("db error")
	}
	c.ID = 1
	return nil
}

func (m *MockCatalogRepository) GetCategorias() ([]models.CategoriaServicio, error) {
	if m.SimulateError {
		return nil, errors.New("db error")
	}
	return []models.CategoriaServicio{{Nombre: "Cortes"}}, nil
}

func (m *MockCatalogRepository) GetCategoriaByID(id uint) (models.CategoriaServicio, error) {
	if m.SimulateError || id == 999 {
		return models.CategoriaServicio{}, errors.New("not found")
	}
	return models.CategoriaServicio{Nombre: "Cortes"}, nil
}

func (m *MockCatalogRepository) UpdateCategoria(c *models.CategoriaServicio) error {
	if m.SimulateError {
		return errors.New("db error")
	}
	return nil
}

func (m *MockCatalogRepository) DeleteCategoria(id uint) error {
	if m.SimulateError {
		return errors.New("db error")
	}
	return nil
}

func (m *MockCatalogRepository) CreatePromocion(p *models.Promocion) error {
	if m.SimulateError {
		return errors.New("db error")
	}
	p.ID = 1
	return nil
}

func (m *MockCatalogRepository) GetPromociones() ([]models.Promocion, error) {
	if m.SimulateError {
		return nil, errors.New("db error")
	}
	return []models.Promocion{{Nombre: "Promo Corte", Descuento: 10}}, nil
}

func (m *MockCatalogRepository) GetPromocionByID(id uint) (models.Promocion, error) {
	if m.SimulateError || id == 999 {
		return models.Promocion{}, errors.New("not found")
	}
	return models.Promocion{Nombre: "Promo Corte", Descuento: 10}, nil
}

func (m *MockCatalogRepository) UpdatePromocion(p *models.Promocion) error {
	if m.SimulateError {
		return errors.New("db error")
	}
	return nil
}

func (m *MockCatalogRepository) DeletePromocion(id uint) error {
	if m.SimulateError {
		return errors.New("db error")
	}
	return nil
}

// ================= SERVICIOS =================

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

func TestGetServicios_Error(t *testing.T) {
	service := NewCatalogService(&MockCatalogRepository{SimulateError: true})

	_, err := service.GetServicios()
	if err == nil {
		t.Error("se esperaba error al listar servicios")
	}
}

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

func TestGetServicioByID_NoEncontrado(t *testing.T) {
	service := NewCatalogService(&MockCatalogRepository{})

	_, ok := service.GetServicioByID(999)
	if ok {
		t.Error("se esperaba que no encuentre el servicio")
	}
}

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

func TestUpdateServicio_Error(t *testing.T) {
	service := NewCatalogService(&MockCatalogRepository{SimulateError: true})

	_, ok := service.UpdateServicio(models.Servicio{})
	if ok {
		t.Error("se esperaba error al actualizar")
	}
}

func TestDeleteServicio_Exitoso(t *testing.T) {
	service := NewCatalogService(&MockCatalogRepository{})

	ok := service.DeleteServicio(1)
	if !ok {
		t.Error("se esperaba eliminar el servicio")
	}
}

func TestDeleteServicio_Error(t *testing.T) {
	service := NewCatalogService(&MockCatalogRepository{SimulateError: true})

	ok := service.DeleteServicio(1)
	if ok {
		t.Error("se esperaba error al eliminar")
	}
}

// ================= CATEGORÍAS =================

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

func TestCreateCategoria_Invalida(t *testing.T) {
	service := NewCatalogService(&MockCatalogRepository{})

	_, ok := service.CreateCategoriaServicio(models.CategoriaServicio{})
	if ok {
		t.Error("se esperaba error por categoria sin nombre")
	}
}

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

func TestGetCategoriaByID_Exitoso(t *testing.T) {
	service := NewCatalogService(&MockCatalogRepository{})

	_, ok := service.GetCategoriaServicioByID(1)
	if !ok {
		t.Error("se esperaba encontrar la categoria")
	}
}

func TestGetCategoriaByID_NoEncontrada(t *testing.T) {
	service := NewCatalogService(&MockCatalogRepository{})

	_, ok := service.GetCategoriaServicioByID(999)
	if ok {
		t.Error("se esperaba que no encuentre la categoria")
	}
}

func TestUpdateCategoria_Exitoso(t *testing.T) {
	service := NewCatalogService(&MockCatalogRepository{})

	_, ok := service.UpdateCategoriaServicio(models.CategoriaServicio{Nombre: "Cortes Premium"})
	if !ok {
		t.Error("se esperaba actualizar la categoria")
	}
}

func TestDeleteCategoria_Exitoso(t *testing.T) {
	service := NewCatalogService(&MockCatalogRepository{})

	ok := service.DeleteCategoriaServicio(1)
	if !ok {
		t.Error("se esperaba eliminar la categoria")
	}
}

// ================= PROMOCIONES =================

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

func TestGetPromocionByID_Exitoso(t *testing.T) {
	service := NewCatalogService(&MockCatalogRepository{})

	_, ok := service.GetPromocionByID(1)
	if !ok {
		t.Error("se esperaba encontrar la promocion")
	}
}

func TestGetPromocionByID_NoEncontrada(t *testing.T) {
	service := NewCatalogService(&MockCatalogRepository{})

	_, ok := service.GetPromocionByID(999)
	if ok {
		t.Error("se esperaba que no encuentre la promocion")
	}
}

func TestUpdatePromocion_Exitoso(t *testing.T) {
	service := NewCatalogService(&MockCatalogRepository{})

	_, ok := service.UpdatePromocion(models.Promocion{Nombre: "Promo Actualizada", Descuento: 15})
	if !ok {
		t.Error("se esperaba actualizar la promocion")
	}
}

func TestDeletePromocion_Exitoso(t *testing.T) {
	service := NewCatalogService(&MockCatalogRepository{})

	ok := service.DeletePromocion(1)
	if !ok {
		t.Error("se esperaba eliminar la promocion")
	}
}