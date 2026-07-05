package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"barberia-cola-virtual/internal/middleware"
	"barberia-cola-virtual/internal/models"
	"barberia-cola-virtual/internal/services"
	"barberia-cola-virtual/internal/utils"

	"github.com/go-chi/chi/v5"
)

// FakeCatalogRepository actúa como el doble en memoria exigido por la rúbrica
type FakeCatalogRepository struct {
	servicios []models.Servicio
	currentID uint
}

func NewFakeCatalogRepository() *FakeCatalogRepository {
	return &FakeCatalogRepository{
		servicios: make([]models.Servicio, 0),
		currentID: 1,
	}
}

func (f *FakeCatalogRepository) CreateServicio(s *models.Servicio) error {
	s.ID = f.currentID
	f.currentID++
	f.servicios = append(f.servicios, *s) // Sí guarda en memoria transitoria
	return nil
}

func (f *FakeCatalogRepository) GetServicios() ([]models.Servicio, error) { return f.servicios, nil }
func (f *FakeCatalogRepository) GetServicioByID(id uint) (models.Servicio, error) {
	for _, s := range f.servicios {
		if s.ID == id {
			return s, nil
		}
	}
	return models.Servicio{}, nil
}
func (f *FakeCatalogRepository) UpdateServicio(s *models.Servicio) error { return nil }
func (f *FakeCatalogRepository) DeleteServicio(id uint) error            { return nil }

// Métodos satélites para cumplir con la interfaz del repositorio
func (f *FakeCatalogRepository) CreateCategoria(c *models.CategoriaServicio) error  { return nil }
func (f *FakeCatalogRepository) GetCategorias() ([]models.CategoriaServicio, error) { return nil, nil }
func (f *FakeCatalogRepository) GetCategoriaByID(id uint) (models.CategoriaServicio, error) {
	return models.CategoriaServicio{}, nil
}
func (f *FakeCatalogRepository) UpdateCategoria(c *models.CategoriaServicio) error { return nil }
func (f *FakeCatalogRepository) DeleteCategoria(id uint) error                     { return nil }
func (f *FakeCatalogRepository) CreatePromocion(p *models.Promocion) error         { return nil }
func (f *FakeCatalogRepository) GetPromociones() ([]models.Promocion, error)       { return nil, nil }
func (f *FakeCatalogRepository) GetPromocionByID(id uint) (models.Promocion, error) {
	return models.Promocion{}, nil
}
func (f *FakeCatalogRepository) UpdatePromocion(p *models.Promocion) error { return nil }
func (f *FakeCatalogRepository) DeletePromocion(id uint) error             { return nil }

// ==================== TEST 1 DE HANDLER: RUTA EXITOSA CON TOKEN ====================
func TestCreateServicioHandler_AdminConToken_201(t *testing.T) {
	fakeRepo := NewFakeCatalogRepository()
	catalogService := services.NewCatalogService(fakeRepo)
	handler := NewCatalogoHandler(catalogService)

	router := chi.NewRouter()
	router.Group(func(admin chi.Router) {
		admin.Use(middleware.AuthMiddleware)
		admin.Use(middleware.AdminOnly)
		admin.Post("/api/v1/servicios", handler.CreateServicio)
	})

	token, err := utils.GenerarToken(1, "admin@gmail.com", "ADMIN")
	if err != nil {
		t.Fatal("no se pudo generar el token")
	}

	body := []byte(`{
		"categoria_id": 1,
		"nombre": "Corte Premium",
		"descripcion": "Corte de cabello estilizado",
		"precio": 10,
		"duracion": 40
	}`)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/servicios", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated && rec.Code != http.StatusOK {
		t.Errorf("se esperaba retorno 200/201 pero llego %d", rec.Code)
	}

	// COMPROBACIÓN FAKE: Validamos que el elemento se guardó en el slice transitorio
	if len(fakeRepo.servicios) != 1 {
		t.Errorf("se esperaba 1 servicio registrado en el Fake, se encontraron: %d", len(fakeRepo.servicios))
	}
}
