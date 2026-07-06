package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"barberia-cola-virtual/internal/models"
	"barberia-cola-virtual/internal/services"

	"github.com/go-chi/chi/v5"
)

// Handlers para el módulo Mi Perfil - Cliente

func CreateCliente(w http.ResponseWriter, r *http.Request) {
	var cliente models.Cliente

	err := json.NewDecoder(r.Body).Decode(&cliente)
	if err != nil {
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}

	clienteCreado := services.CreateCliente(cliente)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(clienteCreado)
}


func GetClientes(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(services.GetClientes())
}

func GetClienteByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	for _, cliente := range services.GetClientes() {
		if strconv.Itoa(cliente.ID) == id {
			json.NewEncoder(w).Encode(cliente)
			return
		}
	}

	http.Error(w, "Cliente no encontrado", http.StatusNotFound)
}

func UpdateCliente(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	var updatedCliente models.Cliente

	err = json.NewDecoder(r.Body).Decode(&updatedCliente)
	if err != nil {
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}

	clienteActualizado, encontrado := services.UpdateCliente(id, updatedCliente)
	if !encontrado {
		http.Error(w, "Cliente no encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(clienteActualizado)
}	

func DeleteCliente(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	eliminado := services.DeleteCliente(id)
	if !eliminado {
		http.Error(w, "Cliente no encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"mensaje": "Cliente eliminado",
	})
}


// Handlers para Preferencias de Pago
func CreatePreferenciaPago(w http.ResponseWriter, r *http.Request) {
	var preferencia models.PreferenciaPago

	err := json.NewDecoder(r.Body).Decode(&preferencia)
	if err != nil {
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}

	preferenciaCreada := services.CreatePreferenciaPago(preferencia)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(preferenciaCreada)
}

func GetPreferenciasPago(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(services.GetPreferenciasPago())
}

// Handlers para Preferencias del Cliente
func CreatePreferenciaCliente(w http.ResponseWriter, r *http.Request) {
	var preferencia models.PreferenciaCliente

	err := json.NewDecoder(r.Body).Decode(&preferencia)
	if err != nil {
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}

	preferenciaCreada := services.CreatePreferenciaCliente(preferencia)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(preferenciaCreada)
}

func GetPreferenciasCliente(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(services.GetPreferenciasCliente())
}