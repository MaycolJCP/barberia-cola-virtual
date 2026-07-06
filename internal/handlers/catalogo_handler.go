package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"barberia-cola-virtual/internal/models"
	"barberia-cola-virtual/internal/services"

	"github.com/go-chi/chi/v5"
)

// ================= SERVICIOS =================

func CreateServicio(w http.ResponseWriter, r *http.Request) {
	var servicio models.Servicio

	err := json.NewDecoder(r.Body).Decode(&servicio)
	if err != nil {
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}

	servicioCreado, ok := services.CreateServicio(servicio)
	if !ok {
		http.Error(w, "Campos obligatorios invalidos", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(servicioCreado)
}

func GetServicios(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(services.GetServicios())
}

func GetServicioByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	servicio, encontrado := services.GetServicioByID(id)
	if !encontrado {
		http.Error(w, "Servicio no encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(servicio)
}

func UpdateServicio(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	var updatedServicio models.Servicio

	err = json.NewDecoder(r.Body).Decode(&updatedServicio)
	if err != nil {
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}

	servicioActualizado, encontrado := services.UpdateServicio(id, updatedServicio)
	if !encontrado {
		http.Error(w, "Servicio no encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(servicioActualizado)
}

func DeleteServicio(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	eliminado := services.DeleteServicio(id)
	if !eliminado {
		http.Error(w, "Servicio no encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"mensaje": "Servicio eliminado correctamente",
	})
}

// ================= CATEGORIAS =================

func CreateCategoriaServicio(w http.ResponseWriter, r *http.Request) {
	var categoria models.CategoriaServicio

	err := json.NewDecoder(r.Body).Decode(&categoria)
	if err != nil {
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}

	categoriaCreada, ok := services.CreateCategoriaServicio(categoria)
	if !ok {
		http.Error(w, "El nombre es obligatorio", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(categoriaCreada)
}

func GetCategoriasServicio(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(services.GetCategoriasServicio())
}

func GetCategoriaServicioByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	categoria, encontrado := services.GetCategoriaServicioByID(id)
	if !encontrado {
		http.Error(w, "Categoria no encontrada", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(categoria)
}

func UpdateCategoriaServicio(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	var updatedCategoria models.CategoriaServicio

	err = json.NewDecoder(r.Body).Decode(&updatedCategoria)
	if err != nil {
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}

	categoriaActualizada, encontrado := services.UpdateCategoriaServicio(id, updatedCategoria)
	if !encontrado {
		http.Error(w, "Categoria no encontrada", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(categoriaActualizada)
}

func DeleteCategoriaServicio(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	eliminado := services.DeleteCategoriaServicio(id)
	if !eliminado {
		http.Error(w, "Categoria no encontrada", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"mensaje": "Categoria eliminada correctamente",
	})
}

// ================= PROMOCIONES =================

func CreatePromocion(w http.ResponseWriter, r *http.Request) {
	var promocion models.Promocion

	err := json.NewDecoder(r.Body).Decode(&promocion)
	if err != nil {
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}

	promocionCreada, ok := services.CreatePromocion(promocion)
	if !ok {
		http.Error(w, "Campos obligatorios invalidos", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(promocionCreada)
}

func GetPromociones(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(services.GetPromociones())
}

func GetPromocionByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	promocion, encontrado := services.GetPromocionByID(id)
	if !encontrado {
		http.Error(w, "Promocion no encontrada", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(promocion)
}

func UpdatePromocion(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	var updatedPromocion models.Promocion

	err = json.NewDecoder(r.Body).Decode(&updatedPromocion)
	if err != nil {
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}

	promocionActualizada, encontrado := services.UpdatePromocion(id, updatedPromocion)
	if !encontrado {
		http.Error(w, "Promocion no encontrada", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(promocionActualizada)
}

func DeletePromocion(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	eliminado := services.DeletePromocion(id)
	if !eliminado {
		http.Error(w, "Promocion no encontrada", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"mensaje": "Promocion eliminada correctamente",
	})
}