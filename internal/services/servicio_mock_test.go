package services

import (
	"testing"

	"barberia-cola-virtual/internal/models"
	"barberia-cola-virtual/internal/repository"
)

// Mock del repositorio de catálogo.
// Se utiliza para simular el repositorio real durante las pruebas,
// evitando acceder al almacenamiento en memoria o a una base de datos.

type mockServicioRepository struct {
	repository.CatalogRepository 
	fueLlamado bool 
	//sirve para verificar si el método CreateServicio fue llamado durante la prueba.             
}

// Prueba unitaria que verifica que NO se pueda crear un servicio con datos inválidos.
//
// Caso probado:
// - Nombre vacío.
// - Precio igual a 0.
// - Duración igual a 0.
//
// Resultado esperado:
// - El servicio debe rechazar la creación.
// - No debe asignarse un ID.
// - La operación debe devolver false.
func TestCreateServicioConRepo_InvalidoNoLlamaRepository_Unico(t *testing.T) {

	// Crea un repositorio falso (Mock).
	mockRepo := &mockServicioRepository{}

	// Crea el servicio utilizando el constructor oficial
	// e inyectando el repositorio falso.
	service := NewCatalogService(mockRepo)

	// Se crea un servicio con datos inválidos para comprobar
	// que las validaciones funcionen correctamente.
	servicioInvalido := models.Servicio{
		Nombre:   "",
		Precio:   0,
		Duracion: 0,
	}

	// Ejecuta el método CreateServicio del servicio.
	// Debe validar primero los datos antes de intentar guardar.
	servicioCreado, ok := service.CreateServicio(servicioInvalido)

	// Si devuelve true significa que aceptó datos inválidos,
	// por lo tanto la prueba falla.
	if ok {
		t.Error("ERROR: Se validó como correcto un servicio con parámetros vacíos inválidos.")
	}

	// Si el ID es distinto de 0 significa que el servicio
	// fue creado cuando realmente debió ser rechazado.
	if servicioCreado.ID != 0 {
		t.Error("ERROR: Se generó un ID de entidad para un flujo que debió ser rechazado.")
	}
}