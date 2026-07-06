package services

import (
	"barberia-cola-virtual/internal/models"
	"barberia-cola-virtual/internal/repository"
)

// ServicioRepository define lo mínimo que necesita el service
// para crear un servicio. Esto permite usar un mock en los tests.
type ServicioRepository interface {
	CreateServicio(servicio models.Servicio) models.Servicio
}


// ================= SERVICIOS =================

func CreateServicio(servicio models.Servicio) (models.Servicio, bool) {
	if servicio.Nombre == "" || servicio.Precio <= 0 || servicio.Duracion <= 0 {
		return models.Servicio{}, false
	}

	return repository.CreateServicio(servicio), true
}


// CreateServicioConRepo permite probar la lógica de negocio
// usando un repositorio externo, como un mock en testing.
func CreateServicioConRepo(servicio models.Servicio, repo ServicioRepository) (models.Servicio, bool) {
	if servicio.Nombre == "" || servicio.Precio <= 0 || servicio.Duracion <= 0 {
		return models.Servicio{}, false
	}

	return repo.CreateServicio(servicio), true
}

func GetServicios() []models.Servicio {
	return repository.GetServicios()
}

func GetServicioByID(id int) (models.Servicio, bool) {
	return repository.GetServicioByID(id)
}

func UpdateServicio(id int, servicio models.Servicio) (models.Servicio, bool) {
	return repository.UpdateServicio(id, servicio)
}

func DeleteServicio(id int) bool {
	return repository.DeleteServicio(id)
}

// ================= CATEGORIAS =================

func CreateCategoriaServicio(categoria models.CategoriaServicio) (models.CategoriaServicio, bool) {
	if categoria.Nombre == "" {
		return models.CategoriaServicio{}, false
	}

	return repository.CreateCategoriaServicio(categoria), true
}

func GetCategoriasServicio() []models.CategoriaServicio {
	return repository.GetCategoriasServicio()
}

func GetCategoriaServicioByID(id int) (models.CategoriaServicio, bool) {
	return repository.GetCategoriaServicioByID(id)
}

func UpdateCategoriaServicio(id int, categoria models.CategoriaServicio) (models.CategoriaServicio, bool) {
	return repository.UpdateCategoriaServicio(id, categoria)
}

func DeleteCategoriaServicio(id int) bool {
	return repository.DeleteCategoriaServicio(id)
}

// ================= PROMOCIONES =================

func CreatePromocion(promocion models.Promocion) (models.Promocion, bool) {
	if promocion.Nombre == "" || promocion.Descuento <= 0 {
		return models.Promocion{}, false
	}

	return repository.CreatePromocion(promocion), true
}

func GetPromociones() []models.Promocion {
	return repository.GetPromociones()
}

func GetPromocionByID(id int) (models.Promocion, bool) {
	return repository.GetPromocionByID(id)
}

func UpdatePromocion(id int, promocion models.Promocion) (models.Promocion, bool) {
	return repository.UpdatePromocion(id, promocion)
}

func DeletePromocion(id int) bool {
	return repository.DeletePromocion(id)
}