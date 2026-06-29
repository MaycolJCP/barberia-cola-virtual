package handlers

import (
	"encoding/json"
	"net/http"

	"barberia-cola-virtual/internal/models"
	"barberia-cola-virtual/internal/services"
	"barberia-cola-virtual/internal/middleware"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var usuario models.Usuario

	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}

	usuarioCreado, mensaje, ok := services.Register(usuario)
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

// LoginHandler maneja las solicitudes de inicio de sesión, decodificando el JSON de entrada, llamando al servicio de autenticación y devolviendo una respuesta JSON con el resultado del login.

func Login(w http.ResponseWriter, r *http.Request) {
	var datosLogin struct {
		Correo   string `json:"correo"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&datosLogin)
	if err != nil {
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}

	usuario, token, mensaje, ok := services.Login(datosLogin.Correo, datosLogin.Password)
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

func Me(w http.ResponseWriter, r *http.Request) {
	usuarioID := r.Context().Value(middleware.UsuarioIDKey)
	correo := r.Context().Value(middleware.CorreoKey)
	rol := r.Context().Value(middleware.RolKey)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"usuario_id": usuarioID,
		"correo":     correo,
		"rol":        rol,
	})
}
