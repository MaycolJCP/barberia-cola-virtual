package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"barberia-cola-virtual/internal/models"
	"barberia-cola-virtual/internal/services"

	"github.com/go-chi/chi/v5"
)

type PerfilHandler struct {
	perfilService *services.PerfilService
}

func NewPerfilHandler(s *services.PerfilService) *PerfilHandler {
	return &PerfilHandler{perfilService: s}
}

func (h *PerfilHandler) CreateCliente(w http.ResponseWriter, r *http.Request) {
	var cliente models.Cliente
	if err := json.NewDecoder(r.Body).Decode(&cliente); err != nil {
		return
	}
	res, ok := h.perfilService.CreateCliente(cliente)
	if !ok {
		http.Error(w, "Error", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func (h *PerfilHandler) GetClientes(w http.ResponseWriter, r *http.Request) {
	clientes, err := h.perfilService.GetClientes()
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(clientes)
}

func (h *PerfilHandler) GetClienteByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.ParseUint(idStr, 10, 32)
	res, ok := h.perfilService.GetClienteByID(uint(id))
	if !ok {
		http.Error(w, "No encontrado", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func (h *PerfilHandler) UpdateCliente(w http.ResponseWriter, r *http.Request) {
	var cliente models.Cliente
	if err := json.NewDecoder(r.Body).Decode(&cliente); err != nil {
		return
	}
	res, ok := h.perfilService.UpdateCliente(cliente)
	if !ok {
		http.Error(w, "Error", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func (h *PerfilHandler) DeleteCliente(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.ParseUint(idStr, 10, 32)
	h.perfilService.DeleteCliente(uint(id))
	w.WriteHeader(http.StatusNoContent)
}

func (h *PerfilHandler) CreatePreferenciaPago(w http.ResponseWriter, r *http.Request) {
	var pref models.PreferenciaPago
	if err := json.NewDecoder(r.Body).Decode(&pref); err != nil {
		return
	}
	res, ok := h.perfilService.CreatePreferenciaPago(pref)
	if !ok {
		http.Error(w, "Error", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func (h *PerfilHandler) GetPreferenciasPago(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("cliente_id")
	id, _ := strconv.ParseUint(idStr, 10, 32)
	prefs, err := h.perfilService.GetPreferenciasPago(uint(id))
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(prefs)
}

func (h *PerfilHandler) CreatePreferenciaCliente(w http.ResponseWriter, r *http.Request) {
	var pref models.PreferenciaCliente
	if err := json.NewDecoder(r.Body).Decode(&pref); err != nil {
		return
	}
	res, ok := h.perfilService.CreatePreferenciaCliente(pref)
	if !ok {
		http.Error(w, "Error", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func (h *PerfilHandler) GetPreferenciasCliente(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("cliente_id")
	id, _ := strconv.ParseUint(idStr, 10, 32)
	res, ok := h.perfilService.GetPreferenciasCliente(uint(id))
	if !ok {
		http.Error(w, "No encontrado", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(res)
}
