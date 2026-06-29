	package repository

	import (
		"barberia-cola-virtual/internal/models"
		"barberia-cola-virtual/internal/storage"
	)

	// ================= SERVICIOS =================

	func CreateServicio(servicio models.Servicio) models.Servicio {
		servicio.ID = len(storage.Servicios) + 1
		storage.Servicios = append(storage.Servicios, servicio)
		return servicio
	}

	func GetServicios() []models.Servicio {
		return storage.Servicios
	}

	func GetServicioByID(id int) (models.Servicio, bool) {
		for _, servicio := range storage.Servicios {
			if servicio.ID == id {
				return servicio, true
			}
		}
		return models.Servicio{}, false
	}

	func UpdateServicio(id int, updatedServicio models.Servicio) (models.Servicio, bool) {
		for i, servicio := range storage.Servicios {
			if servicio.ID == id {
				updatedServicio.ID = servicio.ID
				storage.Servicios[i] = updatedServicio
				return updatedServicio, true
			}
		}
		return models.Servicio{}, false
	}

	func DeleteServicio(id int) bool {
		for i, servicio := range storage.Servicios {
			if servicio.ID == id {
				storage.Servicios = append(storage.Servicios[:i], storage.Servicios[i+1:]...)
				return true
			}
		}
		return false
	}

	// ================= CATEGORIAS =================

	func CreateCategoriaServicio(categoria models.CategoriaServicio) models.CategoriaServicio {
		categoria.ID = len(storage.CategoriasServicio) + 1
		storage.CategoriasServicio = append(storage.CategoriasServicio, categoria)
		return categoria
	}

	func GetCategoriasServicio() []models.CategoriaServicio {
		return storage.CategoriasServicio
	}

	func GetCategoriaServicioByID(id int) (models.CategoriaServicio, bool) {
		for _, categoria := range storage.CategoriasServicio {
			if categoria.ID == id {
				return categoria, true
			}
		}
		return models.CategoriaServicio{}, false
	}

	func UpdateCategoriaServicio(id int, updatedCategoria models.CategoriaServicio) (models.CategoriaServicio, bool) {
		for i, categoria := range storage.CategoriasServicio {
			if categoria.ID == id {
				updatedCategoria.ID = categoria.ID
				storage.CategoriasServicio[i] = updatedCategoria
				return updatedCategoria, true
			}
		}
		return models.CategoriaServicio{}, false
	}

	func DeleteCategoriaServicio(id int) bool {
		for i, categoria := range storage.CategoriasServicio {
			if categoria.ID == id {
				storage.CategoriasServicio = append(storage.CategoriasServicio[:i], storage.CategoriasServicio[i+1:]...)
				return true
			}
		}
		return false
	}

	// ================= PROMOCIONES =================

	func CreatePromocion(promocion models.Promocion) models.Promocion {
		promocion.ID = len(storage.Promociones) + 1
		storage.Promociones = append(storage.Promociones, promocion)
		return promocion
	}

	func GetPromociones() []models.Promocion {
		return storage.Promociones
	}

	func GetPromocionByID(id int) (models.Promocion, bool) {
		for _, promocion := range storage.Promociones {
			if promocion.ID == id {
				return promocion, true
			}
		}
		return models.Promocion{}, false
	}

	func UpdatePromocion(id int, updatedPromocion models.Promocion) (models.Promocion, bool) {
		for i, promocion := range storage.Promociones {
			if promocion.ID == id {
				updatedPromocion.ID = promocion.ID
				storage.Promociones[i] = updatedPromocion
				return updatedPromocion, true
			}
		}
		return models.Promocion{}, false
	}

	func DeletePromocion(id int) bool {
		for i, promocion := range storage.Promociones {
			if promocion.ID == id {
				storage.Promociones = append(storage.Promociones[:i], storage.Promociones[i+1:]...)
				return true
			}
		}
		return false
	}