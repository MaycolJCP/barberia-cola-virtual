package services

import (
	"barberia-cola-virtual/internal/models"
	"barberia-cola-virtual/internal/repository"
	"barberia-cola-virtual/internal/utils"


	"golang.org/x/crypto/bcrypt"

	
)

func Register(usuario models.Usuario) (models.Usuario, string, bool) {
	if usuario.Nombre == "" || usuario.Correo == "" || usuario.Password == "" {
		return models.Usuario{}, "nombre, correo y password son obligatorios", false
	}

	if usuario.Rol == "" {
		usuario.Rol = "CLIENTE"
	}

	_, existe := repository.GetUsuarioByCorreo(usuario.Correo)
	if existe {
		return models.Usuario{}, "el correo ya esta registrado", false
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(usuario.Password), bcrypt.DefaultCost)
if err != nil {
	return models.Usuario{}, "error al encriptar password", false
}

usuario.Password = string(passwordHash)

usuarioCreado := repository.CreateUsuario(usuario)

	return usuarioCreado, "usuario registrado correctamente", true
}


// Login verifica las credenciales del usuario y devuelve un mensaje de éxito o error, junto con un booleano que indica si el login fue exitoso.

func Login(correo string, password string) (models.Usuario, string, string, bool) {
	if correo == "" || password == "" {
		return models.Usuario{}, "", "correo y password son obligatorios", false
	}

	usuario, existe := repository.GetUsuarioByCorreo(correo)
	if !existe {
		return models.Usuario{}, "", "usuario no encontrado", false
	}

	err := bcrypt.CompareHashAndPassword([]byte(usuario.Password), []byte(password))
	if err != nil {
		return models.Usuario{}, "", "password incorrecto", false
	}

	token, err := utils.GenerarToken(usuario.ID, usuario.Correo, usuario.Rol)
	if err != nil {
		return models.Usuario{}, "", "error al generar token", false
	}

	return usuario, token, "login correcto", true
}