package services

import (
	"testing"

	"barberia-cola-virtual/internal/models"
	"barberia-cola-virtual/internal/storage"
)

// =======================================
// TEST CREAR SERVICIO VALIDO
//
// PREPARAR:
// Crear un servicio con nombre, precio y duracion validos.
//
// EJECUTAR:
// Llamar a CreateServicio().
//
// VERIFICAR:
// Debe crear el servicio correctamente.
// =======================================
func TestCreateServicio_Valido(t *testing.T) {

	storage.Servicios = []models.Servicio{}

	servicio := models.Servicio{
		CategoriaID: 1,
		Nombre:      "Corte Clasico",
		Descripcion: "Corte tradicional para caballero",
		Precio:      5,
		Duracion:    30,
	}

	servicioCreado, ok := CreateServicio(servicio)

	if !ok {
		t.Error("se esperaba que el servicio se cree correctamente")
	}

	if servicioCreado.ID == 0 {
		t.Error("se esperaba que el servicio tenga ID")
	}

	if servicioCreado.Nombre != "Corte Clasico" {
		t.Error("el nombre del servicio no coincide")
	}
}

// =======================================
// TEST CREAR SERVICIO INVALIDO
//
// PREPARAR:
// Crear un servicio sin nombre, precio y duracion validos.
//
// EJECUTAR:
// Llamar a CreateServicio().
//
// VERIFICAR:
// Debe fallar porque los campos obligatorios son invalidos.
// =======================================
func TestCreateServicio_Invalido(t *testing.T) {

	storage.Servicios = []models.Servicio{}

	servicio := models.Servicio{
		CategoriaID: 1,
		Nombre:      "",
		Descripcion: "Servicio sin nombre",
		Precio:      0,
		Duracion:    0,
	}

	servicioCreado, ok := CreateServicio(servicio)

	if ok {
		t.Error("se esperaba que el servicio invalido falle")
	}

	if servicioCreado.ID != 0 {
		t.Error("no deberia crear servicio invalido")
	}
}

// =======================================
// TEST CREAR CATEGORIA VALIDA
//
// PREPARAR:
// Crear una categoria con nombre valido.
//
// EJECUTAR:
// Llamar a CreateCategoriaServicio().
//
// VERIFICAR:
// Debe crear la categoria correctamente.
// =======================================
func TestCreateCategoriaServicio_Valida(t *testing.T) {

	storage.CategoriasServicio = []models.CategoriaServicio{}

	categoria := models.CategoriaServicio{
		Nombre:      "Cabello",
		Descripcion: "Servicios relacionados al cabello",
	}

	categoriaCreada, ok := CreateCategoriaServicio(categoria)

	if !ok {
		t.Error("se esperaba crear la categoria")
	}

	if categoriaCreada.ID == 0 {
		t.Error("se esperaba ID generado")
	}
}

// =======================================
// TEST CREAR CATEGORIA INVALIDA
//
// PREPARAR:
// Crear una categoria sin nombre.
//
// EJECUTAR:
// Llamar a CreateCategoriaServicio().
//
// VERIFICAR:
// Debe fallar porque el nombre es obligatorio.
// =======================================
func TestCreateCategoriaServicio_Invalida(t *testing.T) {

	storage.CategoriasServicio = []models.CategoriaServicio{}

	categoria := models.CategoriaServicio{
		Nombre:      "",
		Descripcion: "Categoria invalida",
	}

	categoriaCreada, ok := CreateCategoriaServicio(categoria)

	if ok {
		t.Error("se esperaba que falle")
	}

	if categoriaCreada.ID != 0 {
		t.Error("no deberia crear categoria invalida")
	}
}

// =======================================
// TEST CREAR PROMOCION VALIDA
//
// PREPARAR:
// Crear una promocion valida.
//
// EJECUTAR:
// Llamar a CreatePromocion().
//
// VERIFICAR:
// Debe crear la promocion correctamente.
// =======================================
func TestCreatePromocion_Valida(t *testing.T) {

	storage.Promociones = []models.Promocion{}

	promocion := models.Promocion{
		ServicioID: 1,
		Nombre:     "Promo Lunes",
		Descuento:  10,
	}

	promocionCreada, ok := CreatePromocion(promocion)

	if !ok {
		t.Error("se esperaba crear la promocion")
	}

	if promocionCreada.ID == 0 {
		t.Error("se esperaba ID generado")
	}
}

// =======================================
// TEST CREAR PROMOCION INVALIDA
//
// PREPARAR:
// Crear una promocion sin nombre.
//
// EJECUTAR:
// Llamar a CreatePromocion().
//
// VERIFICAR:
// Debe fallar.
// =======================================
func TestCreatePromocion_Invalida(t *testing.T) {

	storage.Promociones = []models.Promocion{}

	promocion := models.Promocion{
		ServicioID: 1,
		Nombre:     "",
		Descuento:  0,
	}

	promocionCreada, ok := CreatePromocion(promocion)

	if ok {
		t.Error("se esperaba que falle")
	}

	if promocionCreada.ID != 0 {
		t.Error("no deberia crear promocion")
	}
}