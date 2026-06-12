package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"barberia-cola-virtual/internal/models"
	"barberia-cola-virtual/internal/storage"

	"github.com/go-chi/chi/v5"
)

func GetClienteByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	for _, cliente := range storage.Clientes {
		if strconv.Itoa(cliente.ID) == id {
			json.NewEncoder(w).Encode(cliente)
			return
		}
	}

	http.Error(w, "Cliente no encontrado", http.StatusNotFound)
}

func UpdateCliente(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var updatedCliente models.Cliente

	err := json.NewDecoder(r.Body).Decode(&updatedCliente)
	if err != nil {
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}

	for i, cliente := range storage.Clientes {
		if strconv.Itoa(cliente.ID) == id {
			updatedCliente.ID = cliente.ID
			storage.Clientes[i] = updatedCliente

			json.NewEncoder(w).Encode(updatedCliente)
			return
		}
	}

	http.Error(w, "Cliente no encontrado", http.StatusNotFound)
}