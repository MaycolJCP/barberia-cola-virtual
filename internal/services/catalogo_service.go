package services

import (
	"barberia-cola-virtual/internal/models"
	"barberia-cola-virtual/internal/repository"
)

// CatalogService representa la capa de lógica de negocio del módulo Catálogo.
//
// Esta capa se encuentra entre el Handler y el Repository:
//
// Handler -> Service -> Repository
//
// Su responsabilidad principal es validar reglas del negocio antes de guardar
// o consultar datos en la base de datos.
type CatalogService struct {
	// repo es la interfaz del repositorio del catálogo.
	// El Service depende de una interfaz y no de una implementación concreta,
	// lo que facilita cambiar la base de datos o usar mocks en tests.
	repo repository.CatalogRepository
}

// NewCatalogService crea una nueva instancia del Service de catálogo.
//
// Recibe el repository desde main.go mediante inyección de dependencias.
// Esto evita que el Service cree directamente su propia dependencia.
func NewCatalogService(repo repository.CatalogRepository) *CatalogService {
	return &CatalogService{repo: repo}
}

// ================= SERVICIOS =================

// CreateServicio valida y crea un nuevo servicio.
//
// Reglas de negocio:
// - El nombre no puede estar vacío.
// - El precio debe ser mayor a cero.
// - La duración debe ser mayor a cero.
//
// Si alguna regla falla, no se llama al Repository y se devuelve false.
func (s *CatalogService) CreateServicio(servicio models.Servicio) (models.Servicio, bool) {
	if servicio.Nombre == "" || servicio.Precio <= 0 || servicio.Duracion <= 0 {
		return models.Servicio{}, false
	}

	// Si las validaciones se cumplen, se llama al Repository para guardar en PostgreSQL.
	err := s.repo.CreateServicio(&servicio)
	return servicio, err == nil
}

// GetServicios obtiene todos los servicios registrados.
func (s *CatalogService) GetServicios() ([]models.Servicio, error) {
	return s.repo.GetServicios()
}

// GetServicioByID busca un servicio por su ID.
// Devuelve false cuando el Repository no encuentra el registro.
func (s *CatalogService) GetServicioByID(id uint) (models.Servicio, bool) {
	serv, err := s.repo.GetServicioByID(id)
	return serv, err == nil
}

// UpdateServicio actualiza los datos de un servicio.
// La validación principal de existencia queda delegada al Repository/GORM.
func (s *CatalogService) UpdateServicio(servicio models.Servicio) (models.Servicio, bool) {
	err := s.repo.UpdateServicio(&servicio)
	return servicio, err == nil
}

// DeleteServicio elimina un servicio por ID.
func (s *CatalogService) DeleteServicio(id uint) bool {
	return s.repo.DeleteServicio(id) == nil
}

// ================= CATEGORIAS =================

// CreateCategoriaServicio crea una categoría de servicios.
//
// Regla de negocio:
// - La categoría debe tener nombre.
func (s *CatalogService) CreateCategoriaServicio(categoria models.CategoriaServicio) (models.CategoriaServicio, bool) {
	if categoria.Nombre == "" {
		return models.CategoriaServicio{}, false
	}

	err := s.repo.CreateCategoria(&categoria)
	return categoria, err == nil
}

// GetCategoriasServicio obtiene todas las categorías.
func (s *CatalogService) GetCategoriasServicio() ([]models.CategoriaServicio, error) {
	return s.repo.GetCategorias()
}

// GetCategoriaServicioByID busca una categoría por ID.
func (s *CatalogService) GetCategoriaServicioByID(id uint) (models.CategoriaServicio, bool) {
	cat, err := s.repo.GetCategoriaByID(id)
	return cat, err == nil
}

// UpdateCategoriaServicio actualiza una categoría existente.
func (s *CatalogService) UpdateCategoriaServicio(categoria models.CategoriaServicio) (models.CategoriaServicio, bool) {
	err := s.repo.UpdateCategoria(&categoria)
	return categoria, err == nil
}

// DeleteCategoriaServicio elimina una categoría por ID.
func (s *CatalogService) DeleteCategoriaServicio(id uint) bool {
	return s.repo.DeleteCategoria(id) == nil
}

// ================= PROMOCIONES =================

// CreatePromocion crea una promoción.
//
// Reglas de negocio:
// - La promoción debe tener nombre.
// - El descuento debe ser mayor a cero.
func (s *CatalogService) CreatePromocion(promocion models.Promocion) (models.Promocion, bool) {
	if promocion.Nombre == "" || promocion.Descuento <= 0 {
		return models.Promocion{}, false
	}

	err := s.repo.CreatePromocion(&promocion)
	return promocion, err == nil
}

// GetPromociones obtiene todas las promociones.
func (s *CatalogService) GetPromociones() ([]models.Promocion, error) {
	return s.repo.GetPromociones()
}

// GetPromocionByID busca una promoción por ID.
func (s *CatalogService) GetPromocionByID(id uint) (models.Promocion, bool) {
	promo, err := s.repo.GetPromocionByID(id)
	return promo, err == nil
}

// UpdatePromocion actualiza una promoción existente.
func (s *CatalogService) UpdatePromocion(promocion models.Promocion) (models.Promocion, bool) {
	err := s.repo.UpdatePromocion(&promocion)
	return promocion, err == nil
}

// DeletePromocion elimina una promoción por ID.
func (s *CatalogService) DeletePromocion(id uint) bool {
	return s.repo.DeletePromocion(id) == nil
}
