package services

import (
	"barberia-cola-virtual/internal/models"
	"errors"
	"testing"
)

// MockAuthRepository cumple con la interfaz repository.AuthRepository requerida por tu servicio real
type MockAuthRepository struct {
	WasCreatedCalled bool
}

func (m *MockAuthRepository) GetUsuarioByCorreo(correo string) (models.Usuario, error) {
	// Simulamos que el usuario NO existe para que el flujo de registro continúe sin problemas
	return models.Usuario{}, errors.New("record not found")
}

func (m *MockAuthRepository) CreateUsuario(u *models.Usuario) error {
	m.WasCreatedCalled = true
	return nil
}

func TestRegister_UsuarioValido(t *testing.T) {
	// 1. Inicializamos el mock simulando el repositorio correcto
	mockRepo := &MockAuthRepository{WasCreatedCalled: false}
	authService := NewAuthService(mockRepo)

	// 2. Usamos los nombres exactos de tus campos en español definidos en tu struct Usuario
	usuarioPrueba := models.Usuario{
		Nombre:   "Juan Perez",
		Correo:   "test@gmail.com",
		Password: "password123",
		Rol:      "CLIENTE",
	}

	// 3. Capturamos los tres valores de retorno exactos que tiene tu función real: (models.Usuario, string, bool)
	_, mensaje, ok := authService.Register(usuarioPrueba)

	if !ok {
		t.Fatalf("se esperaba un registro valido pero fallo: %s", mensaje)
	}

	// 4. Validamos que la regla de negocio llamó al almacenamiento
	if !mockRepo.WasCreatedCalled {
		t.Error("se esperaba que el servicio llamara a CreateUsuario en el repositorio")
	}
}
