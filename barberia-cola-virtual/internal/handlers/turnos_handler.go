package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"barberia-cola-virtual/internal/models"
	"barberia-cola-virtual/internal/services"

	"github.com/go-chi/chi/v5"
)

// ================= TURNOS =================

func CreateTurno(w http.ResponseWriter, r *http.Request) {
	var turno models.Turno

	err := json.NewDecoder(r.Body).Decode(&turno)
	if err != nil {
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}

	turnoCreado, ok := services.CreateTurno(turno)
	if !ok {
		http.Error(w, "cliente_id y servicio_id son obligatorios", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(turnoCreado)
}

func GetTurnos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(services.GetTurnos())
}

func GetTurnoByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	turno, encontrado := services.GetTurnoByID(id)
	if !encontrado {
		http.Error(w, "Turno no encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(turno)
}

func UpdateTurno(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	var updatedTurno models.Turno

	err = json.NewDecoder(r.Body).Decode(&updatedTurno)
	if err != nil {
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}

	turnoActualizado, encontrado := services.UpdateTurno(id, updatedTurno)
	if !encontrado {
		http.Error(w, "Turno no encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(turnoActualizado)
}

func DeleteTurno(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	eliminado := services.DeleteTurno(id)
	if !eliminado {
		http.Error(w, "Turno no encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"mensaje": "Turno eliminado correctamente",
	})
}

// ================= SEGUIMIENTO =================

func GetSeguimientosTurno(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(services.GetSeguimientosTurno())
}

// ================= NOTIFICACIONES =================

func GetNotificaciones(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(services.GetNotificaciones())
}