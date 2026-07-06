package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"barberia-cola-virtual/internal/middleware"
	"barberia-cola-virtual/internal/services"

	"github.com/go-chi/chi/v5"
)

// ==================== TEST 2 DE HANDLER: PROTECCIÓN 401 UNAUTHORIZED ====================
func TestCreateServicioHandler_SinToken_401(t *testing.T) {
	// Reutiliza el Fake en memoria sin generar conflictos globales
	fakeRepo := NewFakeCatalogRepository()
	catalogService := services.NewCatalogService(fakeRepo)
	handler := NewCatalogoHandler(catalogService)

	router := chi.NewRouter()
	router.Group(func(admin chi.Router) {
		admin.Use(middleware.AuthMiddleware)
		admin.Use(middleware.AdminOnly)
		admin.Post("/api/v1/servicios", handler.CreateServicio)
	})

	body := []byte(`{
		"categoria_id": 1,
		"nombre": "Corte Clasico",
		"descripcion": "Intento de guardar sin login",
		"precio": 5,
		"duracion": 30
	}`)

	// Enviamos la petición SIN cabecera "Authorization"
	req := httptest.NewRequest(http.MethodPost, "/api/v1/servicios", bytes.NewBuffer(body))
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	// COMPROBACIÓN CRÍTICA: La ruta debe denegar el acceso devolviendo 401
	if rec.Code != http.StatusUnauthorized {
		t.Errorf("EXIGENCIA DE RÚBRICA: Se esperaba código 401 (Unauthorized), pero se obtuvo %d", rec.Code)
	}
}
