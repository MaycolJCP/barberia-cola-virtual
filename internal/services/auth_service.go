package services

import (
	"barberia-cola-virtual/internal/models"
	"barberia-cola-virtual/internal/repository"
	"barberia-cola-virtual/internal/utils"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(usuario models.Usuario) (models.Usuario, string, bool) {
	if usuario.Nombre == "" || usuario.Correo == "" || usuario.Password == "" {
		return models.Usuario{}, "nombre, correo y password son obligatorios", false
	}

	if usuario.Rol == "" {
		usuario.Rol = "CLIENTE"
	}

	_, err := s.repo.GetUsuarioByCorreo(usuario.Correo)
	if err == nil {
		return models.Usuario{}, "el correo ya esta registrado", false
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(usuario.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.Usuario{}, "error al encriptar password", false
	}
	usuario.Password = string(passwordHash)

	err = s.repo.CreateUsuario(&usuario)
	if err != nil {
		return models.Usuario{}, "error interno al guardar el usuario", false
	}

	return usuario, "usuario registrado correctamente", true
}

func (s *AuthService) Login(correo string, password string) (models.Usuario, string, string, bool) {
	if correo == "" || password == "" {
		return models.Usuario{}, "", "correo y password son obligatorios", false
	}

	usuario, err := s.repo.GetUsuarioByCorreo(correo)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Usuario{}, "", "usuario no encontrado", false
		}
		return models.Usuario{}, "", "error interno de autenticación", false
	}

	err = bcrypt.CompareHashAndPassword([]byte(usuario.Password), []byte(password))
	if err != nil {
		return models.Usuario{}, "", "password incorrecto", false
	}

	token, err := utils.GenerarToken(int(usuario.ID), usuario.Correo, usuario.Rol)
	if err != nil {
		return models.Usuario{}, "", "error al generar token", false
	}

	return usuario, token, "login correcto", true
}
