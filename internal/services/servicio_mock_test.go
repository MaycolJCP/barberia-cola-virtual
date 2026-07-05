package services

import (
	"testing"

	"barberia-cola-virtual/internal/models"
	"barberia-cola-virtual/internal/repository"
)

// Aseguramos que herede o use la estructura de repositorio
type mockServicioRepository struct {
	repository.CatalogRepository
	fueLlamado bool
}

func TestCreateServicioConRepo_InvalidoNoLlamaRepository_Unico(t *testing.T) {
	mockRepo := &mockServicioRepository{}
	service := NewCatalogService(mockRepo) // Usamos el constructor oficial

	servicioInvalido := models.Servicio{
		Nombre:   "",
		Precio:   0,
		Duracion: 0,
	}

	// Ejecuta directo desde el servicio
	servicioCreado, ok := service.CreateServicio(servicioInvalido)

	if ok {
		t.Error("ERROR: Se validó como correcto un servicio con parámetros vacíos inválidos.")
	}
	if servicioCreado.ID != 0 {
		t.Error("ERROR: Se generó un ID de entidad para un flujo que debió ser rechazado.")
	}
}
