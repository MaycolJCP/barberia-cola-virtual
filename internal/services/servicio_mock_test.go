package services

import (
	"testing"

	"barberia-cola-virtual/internal/models"
)



// mockServicioRepository simula un repositorio.
// Solo registra si fue llamado o no.
//Mock. Es un objeto que reemplaza al Repository real durante la prueba
//es una simulación de un objeto real. Permite probar únicamente 
// la lógica del Service sin utilizar una base de datos

type mockServicioRepository struct {
	fueLlamado bool
}
//Aquí creo un Mock. 
// Es un objeto que reemplaza al Repository real durante la prueba.


func (m *mockServicioRepository) CreateServicio(servicio models.Servicio) models.Servicio {
	//Esta función simula el comportamiento del Repository.
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
	//Aquí preparo un Servicio inválido


	servicioCreado, ok := CreateServicioConRepo(servicio, mockRepo)
	//Aquí ejecuto el Service utilizando el Mock

	if ok {
		t.Error("se esperaba que el servicio invalido falle")
	}

	if servicioCreado.ID != 0 {
		t.Error("no deberia crear servicio invalido")
	}

	if mockRepo.fueLlamado {
		t.Error("el repository no debio ser llamado")
	//"Finalmente verifico que el Repository nunca haya sido llamado."
	//"Ese era justamente el objetivo de este test
	}
}