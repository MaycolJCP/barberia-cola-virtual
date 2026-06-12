package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"barberia-cola-virtual/internal/storage"

	"github.com/go-chi/chi/v5"
)

func GetServicios(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(storage.Servicios)
}

func GetServicioByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	for _, servicio := range storage.Servicios {
		if strconv.Itoa(servicio.ID) == id {
			json.NewEncoder(w).Encode(servicio)
			return
		}
	}

	http.Error(w, "Servicio no encontrado", http.StatusNotFound)
}

func BuscarServicios(w http.ResponseWriter, r *http.Request) {
	query := strings.ToLower(r.URL.Query().Get("q"))

	var resultados []any

	for _, servicio := range storage.Servicios {
		if strings.Contains(strings.ToLower(servicio.Nombre), query) {
			resultados = append(resultados, servicio)
		}
	}

	json.NewEncoder(w).Encode(resultados)
}