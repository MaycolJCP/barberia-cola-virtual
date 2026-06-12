package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"barberia-cola-virtual/internal/models"
	"barberia-cola-virtual/internal/storage"

	"github.com/go-chi/chi/v5"
)

func CreateTurno(w http.ResponseWriter, r *http.Request) {
	var turno models.Turno

	err := json.NewDecoder(r.Body).Decode(&turno)
	if err != nil {
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}

	if turno.ClienteID <= 0 || turno.ServicioID <= 0 {
		http.Error(w, "cliente_id y servicio_id son obligatorios", http.StatusBadRequest)
		return
	}

	turno.ID = len(storage.Turnos) + 1
	turno.Estado = "ESPERANDO"
	turno.Posicion = len(storage.Turnos) + 1
	turno.PersonasDelante = len(storage.Turnos)
	turno.TiempoEstimadoMinutos = turno.PersonasDelante * 15
	turno.HoraEstimada = time.Now().Add(time.Duration(turno.TiempoEstimadoMinutos) * time.Minute)

	storage.Turnos = append(storage.Turnos, turno)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(turno)
}

func GetTurnos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(storage.Turnos)
}

func GetTurnoByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	for _, turno := range storage.Turnos {
		if strconv.Itoa(turno.ID) == id {
			json.NewEncoder(w).Encode(turno)
			return
		}
	}

	http.Error(w, "Turno no encontrado", http.StatusNotFound)
}

func UpdateTurno(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var updatedTurno models.Turno

	err := json.NewDecoder(r.Body).Decode(&updatedTurno)
	if err != nil {
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}

	for i, turno := range storage.Turnos {
		if strconv.Itoa(turno.ID) == id {
			updatedTurno.ID = turno.ID
			storage.Turnos[i] = updatedTurno
			json.NewEncoder(w).Encode(updatedTurno)
			return
		}
	}

	http.Error(w, "Turno no encontrado", http.StatusNotFound)
}

func CancelarTurno(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	for i, turno := range storage.Turnos {
		if strconv.Itoa(turno.ID) == id {
			storage.Turnos[i].Estado = "CANCELADO"
			json.NewEncoder(w).Encode(storage.Turnos[i])
			return
		}
	}

	http.Error(w, "Turno no encontrado", http.StatusNotFound)
}