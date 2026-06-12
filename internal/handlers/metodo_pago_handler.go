package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"barberia-cola-virtual/internal/models"
	"barberia-cola-virtual/internal/storage"

	"github.com/go-chi/chi/v5"
)

func CreateMetodoPago(w http.ResponseWriter, r *http.Request) {
	var metodo models.MetodoPago

	err := json.NewDecoder(r.Body).Decode(&metodo)
	if err != nil {
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}

	metodo.ID = len(storage.MetodosPago) + 1
	metodo.Estado = "ACTIVO"

	storage.MetodosPago = append(storage.MetodosPago, metodo)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(metodo)
}

func GetMetodosPago(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(storage.MetodosPago)
}

func GetMetodoPagoByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	for _, metodo := range storage.MetodosPago {
		if strconv.Itoa(metodo.ID) == id {
			json.NewEncoder(w).Encode(metodo)
			return
		}
	}

	http.Error(w, "Metodo de pago no encontrado", http.StatusNotFound)
}

func UpdateMetodoPago(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var updatedMetodo models.MetodoPago

	err := json.NewDecoder(r.Body).Decode(&updatedMetodo)
	if err != nil {
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}

	for i, metodo := range storage.MetodosPago {
		if strconv.Itoa(metodo.ID) == id {
			updatedMetodo.ID = metodo.ID
			storage.MetodosPago[i] = updatedMetodo
			json.NewEncoder(w).Encode(updatedMetodo)
			return
		}
	}

	http.Error(w, "Metodo de pago no encontrado", http.StatusNotFound)
}

func DeleteMetodoPago(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	for i, metodo := range storage.MetodosPago {
		if strconv.Itoa(metodo.ID) == id {
			storage.MetodosPago = append(
				storage.MetodosPago[:i],
				storage.MetodosPago[i+1:]...,
			)

			json.NewEncoder(w).Encode(map[string]string{
				"message": "Metodo de pago eliminado correctamente",
			})
			return
		}
	}

	http.Error(w, "Metodo de pago no encontrado", http.StatusNotFound)
}