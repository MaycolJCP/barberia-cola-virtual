package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"barberia-cola-virtual/internal/models"
	"barberia-cola-virtual/internal/services"

	"github.com/go-chi/chi/v5"
)

type CatalogoHandler struct {
	catalogService *services.CatalogService
}

func NewCatalogoHandler(s *services.CatalogService) *CatalogoHandler {
	return &CatalogoHandler{catalogService: s}
}

// ================= SERVICIOS =================
func (h *CatalogoHandler) CreateServicio(w http.ResponseWriter, r *http.Request) {
	var servicio models.Servicio
	if err := json.NewDecoder(r.Body).Decode(&servicio); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	res, ok := h.catalogService.CreateServicio(servicio)
	if !ok {
		http.Error(w, "Error al crear servicio", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func (h *CatalogoHandler) GetServicios(w http.ResponseWriter, r *http.Request) {
	servicios, err := h.catalogService.GetServicios()
	if err != nil {
		http.Error(w, "Error al obtener servicios", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(servicios)
}

func (h *CatalogoHandler) GetServicioByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.ParseUint(idStr, 10, 32)
	res, ok := h.catalogService.GetServicioByID(uint(id))
	if !ok {
		http.Error(w, "Servicio no encontrado", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func (h *CatalogoHandler) UpdateServicio(w http.ResponseWriter, r *http.Request) {
	var servicio models.Servicio
	if err := json.NewDecoder(r.Body).Decode(&servicio); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	res, ok := h.catalogService.UpdateServicio(servicio)
	if !ok {
		http.Error(w, "Error al actualizar", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func (h *CatalogoHandler) DeleteServicio(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.ParseUint(idStr, 10, 32)
	if !h.catalogService.DeleteServicio(uint(id)) {
		http.Error(w, "Error al eliminar", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// ================= CATEGORIAS =================
func (h *CatalogoHandler) CreateCategoriaServicio(w http.ResponseWriter, r *http.Request) {
	var cat models.CategoriaServicio
	if err := json.NewDecoder(r.Body).Decode(&cat); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	res, ok := h.catalogService.CreateCategoriaServicio(cat)
	if !ok {
		http.Error(w, "Error al crear categoría", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func (h *CatalogoHandler) GetCategoriasServicio(w http.ResponseWriter, r *http.Request) {
	cats, err := h.catalogService.GetCategoriasServicio()
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(cats)
}

func (h *CatalogoHandler) GetCategoriaServicioByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.ParseUint(idStr, 10, 32)
	res, ok := h.catalogService.GetCategoriaServicioByID(uint(id))
	if !ok {
		http.Error(w, "No encontrado", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func (h *CatalogoHandler) UpdateCategoriaServicio(w http.ResponseWriter, r *http.Request) {
	var cat models.CategoriaServicio
	if err := json.NewDecoder(r.Body).Decode(&cat); err != nil {
		return
	}
	res, ok := h.catalogService.UpdateCategoriaServicio(cat)
	if !ok {
		http.Error(w, "Error", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func (h *CatalogoHandler) DeleteCategoriaServicio(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.ParseUint(idStr, 10, 32)
	h.catalogService.DeleteCategoriaServicio(uint(id))
	w.WriteHeader(http.StatusNoContent)
}

// ================= PROMOCIONES =================
func (h *CatalogoHandler) CreatePromocion(w http.ResponseWriter, r *http.Request) {
	var promo models.Promocion
	if err := json.NewDecoder(r.Body).Decode(&promo); err != nil {
		return
	}
	res, ok := h.catalogService.CreatePromocion(promo)
	if !ok {
		http.Error(w, "Error", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func (h *CatalogoHandler) GetPromociones(w http.ResponseWriter, r *http.Request) {
	promos, err := h.catalogService.GetPromociones()
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(promos)
}

func (h *CatalogoHandler) GetPromocionByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.ParseUint(idStr, 10, 32)
	res, ok := h.catalogService.GetPromocionByID(uint(id))
	if !ok {
		http.Error(w, "No encontrado", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func (h *CatalogoHandler) UpdatePromocion(w http.ResponseWriter, r *http.Request) {
	var promo models.Promocion
	if err := json.NewDecoder(r.Body).Decode(&promo); err != nil {
		return
	}
	res, ok := h.catalogService.UpdatePromocion(promo)
	if !ok {
		http.Error(w, "Error", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func (h *CatalogoHandler) DeletePromocion(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.ParseUint(idStr, 10, 32)
	h.catalogService.DeletePromocion(uint(id))
	w.WriteHeader(http.StatusNoContent)
}
