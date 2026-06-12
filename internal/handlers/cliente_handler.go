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
			w.Header().Set("Content-Type", "application/json")
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

			if updatedCliente.Nombre != "" {
				storage.Clientes[i].Nombre = updatedCliente.Nombre
			}
			if updatedCliente.Correo != "" {
				storage.Clientes[i].Correo = updatedCliente.Correo
			}
			if updatedCliente.Telefono != "" {
				storage.Clientes[i].Telefono = updatedCliente.Telefono
			}
			if updatedCliente.Direccion != "" {
				storage.Clientes[i].Direccion = updatedCliente.Direccion
			}
			if updatedCliente.Genero != "" {
				storage.Clientes[i].Genero = updatedCliente.Genero
			}
			if updatedCliente.UltimaVisita != "" {
				storage.Clientes[i].UltimaVisita = updatedCliente.UltimaVisita
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(storage.Clientes[i])
			return
		}
	}

	http.Error(w, "Cliente no encontrado", http.StatusNotFound)
}
