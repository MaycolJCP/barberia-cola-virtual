package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"barberia-cola-virtual/internal/models"
	"barberia-cola-virtual/internal/services"

	"github.com/go-chi/chi/v5"
)

// CatalogoHandler es la capa HTTP del módulo Catálogo de Servicios.
//
// Esta estructura recibe las peticiones que llegan desde Postman, navegador
// o cualquier cliente HTTP. No debe guardar datos directamente en la base,
// sino comunicarse con el Service para mantener la arquitectura en capas:

type CatalogoHandler struct {
	catalogService *services.CatalogService
}

// NewCatalogoHandler recibe el Service del catálogo y devuelve un Handler listo para usar.
// el Handler no crea el Service, sino que lo recibe desde main.go. 
func NewCatalogoHandler(s *services.CatalogService) *CatalogoHandler {
	return &CatalogoHandler{catalogService: s}
}

// ================= SERVICIOS =================

// CreateServicio maneja la petición HTTP para crear un nuevo servicio.
//
// Flujo:
// 1. Lee el JSON enviado por el cliente.
// 2. Convierte ese JSON en un modelo Servicio.
// 3. Envía el modelo al Service para aplicar reglas de negocio.
// 4. Si todo sale bien, responde 201 Created con el servicio creado.

func (h *CatalogoHandler) CreateServicio(w http.ResponseWriter, r *http.Request) {
	var servicio models.Servicio

	// Decodifica el cuerpo JSON de la petición HTTP.
	// Si el JSON está mal escrito, responde 400 Bad Request.
	if err := json.NewDecoder(r.Body).Decode(&servicio); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	// Llama al Service, que es donde se validan las reglas de negocio.
	// Por ejemplo: nombre obligatorio, precio mayor a cero y duración válida.
	res, ok := h.catalogService.CreateServicio(servicio)
	if !ok {
		http.Error(w, "Error al crear servicio", http.StatusBadRequest)
		return
	}

	// Si el servicio se creó correctamente, se responde con 201 Created.
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

// GetServicios maneja la petición HTTP para listar todos los servicios.
//
// Esta función permite que el cliente consulte el catálogo disponible.
// Normalmente puede ser pública, porque los clientes necesitan ver los servicios
// antes de solicitar un turno.
func (h *CatalogoHandler) GetServicios(w http.ResponseWriter, r *http.Request) {
	// Pide al Service la lista de servicios.
	servicios, err := h.catalogService.GetServicios()
	if err != nil {
		http.Error(w, "Error al obtener servicios", http.StatusInternalServerError)
		return
	}

	// Devuelve la lista en formato JSON.
	json.NewEncoder(w).Encode(servicios)
}

// GetServicioByID busca un servicio específico usando el ID que viene en la URL.
//
// Ejemplo de ruta:
// GET /api/v1/servicios/1
func (h *CatalogoHandler) GetServicioByID(w http.ResponseWriter, r *http.Request) {
	// chi.URLParam obtiene el parámetro {id} declarado en la ruta.
	idStr := chi.URLParam(r, "id")

	// Convierte el ID de texto a número.
	// Se usa uint porque los IDs de GORM son enteros positivos.
	id, _ := strconv.ParseUint(idStr, 10, 32)

	// Pide al Service que busque el servicio.
	res, ok := h.catalogService.GetServicioByID(uint(id))
	if !ok {
		// Si no existe, responde 404 Not Found.
		http.Error(w, "Servicio no encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(res)
}

// UpdateServicio actualiza un servicio existente.
//
// El JSON enviado debe incluir los datos del servicio, incluyendo el ID,
// para que GORM sepa qué registro actualizar.
func (h *CatalogoHandler) UpdateServicio(w http.ResponseWriter, r *http.Request) {
	var servicio models.Servicio

	// Lee los datos enviados por el cliente.
	if err := json.NewDecoder(r.Body).Decode(&servicio); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	// Envía el servicio al Service para actualizarlo.
	res, ok := h.catalogService.UpdateServicio(servicio)
	if !ok {
		http.Error(w, "Error al actualizar", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res)
}

// DeleteServicio elimina un servicio usando el ID de la URL.
//
// Esta operación normalmente debe ser solo para ADMIN, porque modifica
// el catálogo de la barbería.
func (h *CatalogoHandler) DeleteServicio(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.ParseUint(idStr, 10, 32)

	// El Service devuelve false si no pudo eliminar.
	if !h.catalogService.DeleteServicio(uint(id)) {
		http.Error(w, "Error al eliminar", http.StatusBadRequest)
		return
	}

	// 204 No Content significa que la operación fue exitosa
	// pero no se devuelve cuerpo en la respuesta.
	w.WriteHeader(http.StatusNoContent)
}

// ================= CATEGORIAS =================

// CreateCategoriaServicio crea una categoría para clasificar servicios.
//
// Ejemplo: "Cortes", "Barba", "Tratamientos".
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

// GetCategoriasServicio lista todas las categorías de servicios.
func (h *CatalogoHandler) GetCategoriasServicio(w http.ResponseWriter, r *http.Request) {
	cats, err := h.catalogService.GetCategoriasServicio()
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(cats)
}

// GetCategoriaServicioByID busca una categoría por ID.
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

// UpdateCategoriaServicio actualiza los datos de una categoría.
func (h *CatalogoHandler) UpdateCategoriaServicio(w http.ResponseWriter, r *http.Request) {
	var cat models.CategoriaServicio

	// Si el JSON no se puede leer, se termina la función.
	// En una mejora futura se podría responder explícitamente con 400.
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

// DeleteCategoriaServicio elimina una categoría por ID.
func (h *CatalogoHandler) DeleteCategoriaServicio(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.ParseUint(idStr, 10, 32)

	h.catalogService.DeleteCategoriaServicio(uint(id))
	w.WriteHeader(http.StatusNoContent)
}

// ================= PROMOCIONES =================

// CreatePromocion crea una promoción asociada al catálogo.
//
// Una promoción puede representar descuentos o beneficios aplicados
// sobre un servicio de la barbería.
func (h *CatalogoHandler) CreatePromocion(w http.ResponseWriter, r *http.Request) {
	var promo models.Promocion

	// Decodifica el JSON recibido.
	// En una mejora futura se podría devolver 400 si ocurre error.
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

// GetPromociones lista todas las promociones registradas.
func (h *CatalogoHandler) GetPromociones(w http.ResponseWriter, r *http.Request) {
	promos, err := h.catalogService.GetPromociones()
	if err != nil {
		return
	}

	json.NewEncoder(w).Encode(promos)
}

// GetPromocionByID busca una promoción por ID.
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

// UpdatePromocion actualiza una promoción existente.
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

// DeletePromocion elimina una promoción por ID.
func (h *CatalogoHandler) DeletePromocion(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.ParseUint(idStr, 10, 32)

	h.catalogService.DeletePromocion(uint(id))
	w.WriteHeader(http.StatusNoContent)
}
