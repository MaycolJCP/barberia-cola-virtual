package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"barberia-cola-virtual/internal/models"
	"barberia-cola-virtual/internal/services"

	"gorm.io/gorm"

	"github.com/go-chi/chi/v5"
)

type MockAuthRepositoryForHandler struct{}

func (m *MockAuthRepositoryForHandler) GetUsuarioByCorreo(correo string) (models.Usuario, error) {
	return models.Usuario{}, gorm.ErrRecordNotFound
}

func (m *MockAuthRepositoryForHandler) CreateUsuario(u *models.Usuario) error {
	if u != nil {
		u.ID = 1
	}
	return nil
}

func TestRegisterHandler_UsuarioValido(t *testing.T) {
	mockRepo := &MockAuthRepositoryForHandler{}
	authService := services.NewAuthService(mockRepo)
	authHandler := NewAuthHandler(authService)

	router := chi.NewRouter()
	router.Post("/api/v1/register", authHandler.Register)

	body := []byte(`{
		"nombre": "Carlos Mendoza",
		"correo": "carlos@gmail.com",
		"password": "password123",
		"rol": "CLIENTE"
	}`)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	t.Log("RESPUESTA:", rec.Body.String())

	if rec.Code != http.StatusCreated && rec.Code != http.StatusOK {
		t.Errorf("se esperaba un status de éxito (200 o 201), pero llegó %d", rec.Code)
	}
}
