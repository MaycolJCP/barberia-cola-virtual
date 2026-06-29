package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"barberia-cola-virtual/internal/middleware"
	"barberia-cola-virtual/internal/models"
	"barberia-cola-virtual/internal/storage"
	"barberia-cola-virtual/internal/utils"

	"github.com/go-chi/chi/v5"
)

func TestCreateServicioHandler_SinToken_401(t *testing.T) {
	router := chi.NewRouter()

	router.Group(func(admin chi.Router) {
		admin.Use(middleware.AuthMiddleware)
		admin.Use(middleware.AdminOnly)
		admin.Post("/api/v1/servicios", CreateServicio)
	})

	body := []byte(`{
		"categoria_id": 1,
		"nombre": "Corte Clasico",
		"descripcion": "Corte tradicional",
		"precio": 5,
		"duracion": 30
	}`)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/servicios", bytes.NewBuffer(body))
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Errorf("se esperaba status %d pero llego %d", http.StatusUnauthorized, rec.Code)
	}
}

func TestCreateServicioHandler_AdminConToken_201(t *testing.T) {
	storage.Servicios = []models.Servicio{}

	router := chi.NewRouter()

	router.Group(func(admin chi.Router) {
		admin.Use(middleware.AuthMiddleware)
		admin.Use(middleware.AdminOnly)
		admin.Post("/api/v1/servicios", CreateServicio)
	})

	token, err := utils.GenerarToken(1, "admin@gmail.com", "ADMIN")
	if err != nil {
		t.Fatal("no se pudo generar token")
	}

	body := []byte(`{
		"categoria_id": 1,
		"nombre": "Corte Clasico",
		"descripcion": "Corte tradicional",
		"precio": 5,
		"duracion": 30
	}`)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/servicios", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+token)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("se esperaba status %d pero llego %d", http.StatusCreated, rec.Code)
	}

	if len(storage.Servicios) != 1 {
		t.Error("se esperaba que el servicio se guarde en memoria")
	}
}