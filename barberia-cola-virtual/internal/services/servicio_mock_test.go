package services

import (
	"testing"

	"barberia-cola-virtual/internal/models"
)

// mockServicioRepository simula un repositorio.
// Solo registra si fue llamado o no.
type mockServicioRepository struct {
	fueLlamado bool
}

func (m *mockServicioRepository) CreateServicio(servicio models.Servicio) models.Servicio {
	m.fueLlamado = true
	servicio.ID = 1
	return servicio
}

func TestCreateServicioConRepo_InvalidoNoLlamaRepository(t *testing.T) {
	mockRepo := &mockServicioRepository{}

	servicio := models.Servicio{
		Nombre:   "",
		Precio:   0,
		Duracion: 0,
	}

	servicioCreado, ok := CreateServicioConRepo(servicio, mockRepo)

	if ok {
		t.Error("se esperaba que el servicio invalido falle")
	}

	if servicioCreado.ID != 0 {
		t.Error("no deberia crear servicio invalido")
	}

	if mockRepo.fueLlamado {
		t.Error("el repository no debio ser llamado")
	}
}