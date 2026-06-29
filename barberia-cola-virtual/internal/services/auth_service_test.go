package services

import (
	"testing"

	"barberia-cola-virtual/internal/models"
	"barberia-cola-virtual/internal/storage"
)

// =======================================
// TEST 1 REGISTER USUARIO VALIDO
//
// PREPARAR:
// Crear un usuario valido.
//
// EJECUTAR:
// Llamar a Register().
//
// VERIFICAR:
// Debe registrarse correctamente.
// =======================================
func TestRegister_UsuarioValido(t *testing.T) {

	// Limpiar memoria antes del test
	storage.Usuarios = []models.Usuario{}

	// Usuario de prueba
	usuario := models.Usuario{
		Nombre:   "Danny",
		Correo:   "danny@gmail.com",
		Password: "123456",
		Rol:      "CLIENTE",
	}

	// Ejecutar
	usuarioCreado, mensaje, ok := Register(usuario)

	// Verificar
	if !ok {
		t.Error("se esperaba registro exitoso")
	}

	if usuarioCreado.ID == 0 {
		t.Error("se esperaba que el usuario tenga ID")
	}

	if mensaje == "" {
		t.Error("se esperaba mensaje de exito")
	}
}

// =======================================
// TEST 3	 REGISTER CAMPOS VACIOS
//
// PREPARAR:
// Crear un usuario sin nombre, correo ni password.
//
// EJECUTAR:
// Llamar a Register().
//
// VERIFICAR:
// Debe fallar porque los campos obligatorios estan vacios.
// =======================================
func TestRegister_CamposVacios(t *testing.T) {

	// Limpiar memoria antes del test
	storage.Usuarios = []models.Usuario{}

	// Usuario invalido
	usuario := models.Usuario{
		Nombre:   "",
		Correo:   "",
		Password: "",
		Rol:      "CLIENTE",
	}

	// Ejecutar
	usuarioCreado, mensaje, ok := Register(usuario)

	// Verificar
	if ok {
		t.Error("se esperaba que el registro falle")
	}

	if usuarioCreado.ID != 0 {
		t.Error("no deberia crear usuario con campos vacios")
	}

	if mensaje == "" {
		t.Error("se esperaba mensaje de error")
	}
}

// =======================================
// TEST 2 REGISTER CORREO DUPLICADO
//
// PREPARAR:
// Registrar un usuario.
// Intentar registrar otro con el mismo correo.
//
// EJECUTAR:
// Llamar a Register() dos veces.
//
// VERIFICAR:
// El segundo registro debe fallar.
// =======================================
func TestRegister_CorreoDuplicado(t *testing.T) {

	// Limpiar memoria antes del test
	storage.Usuarios = []models.Usuario{}

	// Primer usuario
	usuario1 := models.Usuario{
		Nombre:   "Danny",
		Correo:   "danny@gmail.com",
		Password: "123456",
		Rol:      "CLIENTE",
	}

	// Segundo usuario con mismo correo
	usuario2 := models.Usuario{
		Nombre:   "Pedro",
		Correo:   "danny@gmail.com",
		Password: "654321",
		Rol:      "CLIENTE",
	}

	// Registrar primero
	Register(usuario1)

	// Intentar registrar segundo
	usuarioCreado, mensaje, ok := Register(usuario2)

	// Verificar
	if ok {
		t.Error("no deberia permitir correos duplicados")
	}

	if usuarioCreado.ID != 0 {
		t.Error("no deberia crear usuario duplicado")
	}

	if mensaje == "" {
		t.Error("se esperaba mensaje de error")
	}
}

// =======================================
// TEST 4 LOGIN USUARIO VALIDO
//
// PREPARAR:
// Limpiar memoria.
// Registrar un usuario valido.
//
// EJECUTAR:
// Llamar a Login() con correo y password correctos.
//
// VERIFICAR:
// Debe iniciar sesion correctamente.
// Debe devolver un token JWT.
// =======================================
func TestLogin_UsuarioValido(t *testing.T) {

	// Limpiar memoria antes del test
	storage.Usuarios = []models.Usuario{}

	// Usuario que primero se registra
	usuario := models.Usuario{
		Nombre:   "Danny",
		Correo:   "danny@gmail.com",
		Password: "123456",
		Rol:      "CLIENTE",
	}

	// Primero registramos para que exista en memoria
	Register(usuario)

	// Ejecutar login
	usuarioLogueado, token, mensaje, ok := Login("danny@gmail.com", "123456")

	// Verificar
	if !ok {
		t.Error("se esperaba login exitoso")
	}

	if usuarioLogueado.ID == 0 {
		t.Error("se esperaba usuario logueado con ID")
	}

	if token == "" {
		t.Error("se esperaba token JWT")
	}

	if mensaje == "" {
		t.Error("se esperaba mensaje de exito")
	}
}

// =======================================
// TEST 5 LOGIN PASSWORD INCORRECTO
//
// PREPARAR:
// Registrar un usuario valido.
//
// EJECUTAR:
// Intentar login con password incorrecto.
//
// VERIFICAR:
// Debe fallar el login.
// No debe devolver token.
// =======================================
func TestLogin_PasswordIncorrecto(t *testing.T) {

	// Limpiar memoria antes del test
	storage.Usuarios = []models.Usuario{}

	// Usuario de prueba
	usuario := models.Usuario{
		Nombre:   "Danny",
		Correo:   "danny@gmail.com",
		Password: "123456",
		Rol:      "CLIENTE",
	}

	// Registrar usuario
	Register(usuario)

	// Intentar login con password incorrecto
	usuarioLogueado, token, mensaje, ok := Login("danny@gmail.com", "999999")

	// Verificar
	if ok {
		t.Error("el login deberia fallar")
	}

	if usuarioLogueado.ID != 0 {
		t.Error("no deberia devolver usuario")
	}

	if token != "" {
		t.Error("no deberia devolver token")
	}

	if mensaje == "" {
		t.Error("se esperaba mensaje de error")
	}
}
