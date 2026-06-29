package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"barberia-cola-virtual/internal/models"
	"barberia-cola-virtual/internal/storage"
	"barberia-cola-virtual/internal/services"
)

// =======================================
// TEST HANDLER REGISTER USUARIO VALIDO
//
// PREPARAR:
// Crear un JSON valido como si viniera desde Postman.
//
// EJECUTAR:
// Enviar la peticion falsa al handler Register().
//
// VERIFICAR:
// Debe responder HTTP 201 Created.
// =======================================
func TestRegisterHandler_UsuarioValido(t *testing.T) {

	// Limpiar memoria antes del test
	storage.Usuarios = []models.Usuario{}

	// JSON que simula el body enviado desde Postman
	body := []byte(`{
		"nombre": "Danny",
		"correo": "danny@gmail.com",
		"password": "123456",
		"rol": "CLIENTE"
	}`)

	// Crear request falso
	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/register", bytes.NewBuffer(body))

	// Recorder guarda la respuesta del handler
	rec := httptest.NewRecorder()

	// Ejecutar handler
	Register(rec, req)

	// Verificar codigo HTTP
	if rec.Code != http.StatusCreated {
		t.Errorf("se esperaba status %d pero llego %d", http.StatusCreated, rec.Code)
	}

}

// =======================================
// TEST HANDLER LOGIN USUARIO VALIDO
//
// PREPARAR:
// Registrar un usuario en memoria.
// Crear JSON de login valido.
//
// EJECUTAR:
// Enviar la peticion falsa al handler Login().
//
// VERIFICAR:
// Debe responder HTTP 200 OK.
// =======================================
func TestLoginHandler_UsuarioValido(t *testing.T) {

	storage.Usuarios = []models.Usuario{}

	usuario := models.Usuario{
		Nombre:   "Danny",
		Correo:   "danny@gmail.com",
		Password: "123456",
		Rol:      "CLIENTE",
	}

	services.Register(usuario)

	body := []byte(`{
		"correo": "danny@gmail.com",
		"password": "123456"
	}`)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBuffer(body))
	rec := httptest.NewRecorder()

	Login(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("se esperaba status %d pero llego %d", http.StatusOK, rec.Code)
	}
}

// =======================================
// TEST HANDLER LOGIN PASSWORD INCORRECTO
//
// PREPARAR:
// Registrar un usuario en memoria.
// Crear JSON de login con password incorrecto.
//
// EJECUTAR:
// Enviar la peticion falsa al handler Login().
//
// VERIFICAR:
// Debe responder HTTP 401 Unauthorized.
// =======================================
func TestLoginHandler_PasswordIncorrecto(t *testing.T) {

	storage.Usuarios = []models.Usuario{}

	usuario := models.Usuario{
		Nombre:   "Danny",
		Correo:   "danny@gmail.com",
		Password: "123456",
		Rol:      "CLIENTE",
	}

	services.Register(usuario)

	body := []byte(`{
		"correo": "danny@gmail.com",
		"password": "999999"
	}`)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBuffer(body))
	rec := httptest.NewRecorder()

	Login(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Errorf("se esperaba status %d pero llego %d", http.StatusUnauthorized, rec.Code)
	}
}

