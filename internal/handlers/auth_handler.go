package handlers

import (
	"encoding/json"
	"net/http"

	"barberia-cola-virtual/internal/middleware"
	"barberia-cola-virtual/internal/models"
	"barberia-cola-virtual/internal/services"
)

// AuthHandler agrupa los métodos de autenticación usando el servicio inyectado
type AuthHandler struct {
	authService *services.AuthService
}

// NewAuthHandler es el constructor que necesita tu main.go
func NewAuthHandler(s *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: s}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var usuario models.Usuario

	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}

	// CAMBIO: Ahora llama al método del servicio persistente instanciado
	usuarioCreado, mensaje, ok := h.authService.Register(usuario)
	if !ok {
		http.Error(w, mensaje, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"mensaje": mensaje,
		"usuario": models.UsuarioResponse{
			ID:     usuarioCreado.ID,
			Nombre: usuarioCreado.Nombre,
			Correo: usuarioCreado.Correo,
			Rol:    usuarioCreado.Rol,
		},
	})
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var datosLogin struct {
		Correo   string `json:"correo"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&datosLogin)
	if err != nil {
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}

	// CAMBIO: Ahora llama al método del servicio persistente instanciado
	usuario, token, mensaje, ok := h.authService.Login(datosLogin.Correo, datosLogin.Password)
	if !ok {
		http.Error(w, mensaje, http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"mensaje": mensaje,
		"usuario": models.UsuarioResponse{
			ID:     usuario.ID,
			Nombre: usuario.Nombre,
			Correo: usuario.Correo,
			Rol:    usuario.Rol,
		},
		"token": token,
	})
}

func (h *AuthHandler) Me(w http.ResponseWriter, r *http.Request) {
	usuarioID := r.Context().Value(middleware.UsuarioIDKey)
	correo := r.Context().Value(middleware.CorreoKey)
	rol := r.Context().Value(middleware.RolKey)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"usuario_id": usuarioID,
		"correo":     correo,
		"rol":        rol,
	})
}
