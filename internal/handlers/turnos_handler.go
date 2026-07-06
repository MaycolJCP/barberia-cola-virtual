package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"barberia-cola-virtual/internal/models"
	"barberia-cola-virtual/internal/services"

	"github.com/go-chi/chi/v5"
)

type TurnosHandler struct {
	turnoService *services.TurnoService
}

func NewTurnosHandler(s *services.TurnoService) *TurnosHandler {
	return &TurnosHandler{turnoService: s}
}

func (h *TurnosHandler) CreateTurno(w http.ResponseWriter, r *http.Request) {
	var turno models.Turno
	if err := json.NewDecoder(r.Body).Decode(&turno); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	turnoCreado, ok := h.turnoService.CreateTurno(turno)
	if !ok {
		http.Error(w, "Error al crear el turno. Verifique los datos.", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(turnoCreado)
}

func (h *TurnosHandler) GetTurnos(w http.ResponseWriter, r *http.Request) {
	turnos, err := h.turnoService.GetTurnos()
	if err != nil {
		http.Error(w, "Error al obtener turnos", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(turnos)
}

func (h *TurnosHandler) GetTurnoByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	turno, ok := h.turnoService.GetTurnoByID(uint(id))
	if !ok {
		http.Error(w, "Turno no encontrado", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(turno)
}

func (h *TurnosHandler) UpdateTurno(w http.ResponseWriter, r *http.Request) {
	var turno models.Turno
	if err := json.NewDecoder(r.Body).Decode(&turno); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	turnoActualizado, ok := h.turnoService.UpdateTurno(turno)
	if !ok {
		http.Error(w, "Error al actualizar el turno", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(turnoActualizado)
}

func (h *TurnosHandler) DeleteTurno(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if !h.turnoService.DeleteTurno(uint(id)) {
		http.Error(w, "No se pudo eliminar el turno", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *TurnosHandler) GetSeguimientosTurno(w http.ResponseWriter, r *http.Request) {
	segs, err := h.turnoService.GetSeguimientosTurno()
	if err != nil {
		http.Error(w, "Error al obtener seguimientos", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(segs)
}

func (h *TurnosHandler) GetNotificaciones(w http.ResponseWriter, r *http.Request) {
	notifs, err := h.turnoService.GetNotificaciones()
	if err != nil {
		http.Error(w, "Error al obtener notificaciones", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notifs)
}
