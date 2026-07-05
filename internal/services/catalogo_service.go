package services

import (
	"barberia-cola-virtual/internal/models"
	"barberia-cola-virtual/internal/repository"
)

type CatalogService struct {
	repo repository.CatalogRepository
}

func NewCatalogService(repo repository.CatalogRepository) *CatalogService {
	return &CatalogService{repo: repo}
}

// ================= SERVICIOS =================
func (s *CatalogService) CreateServicio(servicio models.Servicio) (models.Servicio, bool) {
	if servicio.Nombre == "" || servicio.Precio <= 0 || servicio.Duracion <= 0 {
		return models.Servicio{}, false
	}
	err := s.repo.CreateServicio(&servicio)
	return servicio, err == nil
}

func (s *CatalogService) GetServicios() ([]models.Servicio, error) {
	return s.repo.GetServicios()
}

func (s *CatalogService) GetServicioByID(id uint) (models.Servicio, bool) {
	serv, err := s.repo.GetServicioByID(id)
	return serv, err == nil
}

func (s *CatalogService) UpdateServicio(servicio models.Servicio) (models.Servicio, bool) {
	err := s.repo.UpdateServicio(&servicio)
	return servicio, err == nil
}

func (s *CatalogService) DeleteServicio(id uint) bool {
	return s.repo.DeleteServicio(id) == nil
}

// ================= CATEGORIAS =================
func (s *CatalogService) CreateCategoriaServicio(categoria models.CategoriaServicio) (models.CategoriaServicio, bool) {
	if categoria.Nombre == "" {
		return models.CategoriaServicio{}, false
	}
	err := s.repo.CreateCategoria(&categoria)
	return categoria, err == nil
}

func (s *CatalogService) GetCategoriasServicio() ([]models.CategoriaServicio, error) {
	return s.repo.GetCategorias()
}

func (s *CatalogService) GetCategoriaServicioByID(id uint) (models.CategoriaServicio, bool) {
	cat, err := s.repo.GetCategoriaByID(id)
	return cat, err == nil
}

func (s *CatalogService) UpdateCategoriaServicio(categoria models.CategoriaServicio) (models.CategoriaServicio, bool) {
	err := s.repo.UpdateCategoria(&categoria)
	return categoria, err == nil
}

func (s *CatalogService) DeleteCategoriaServicio(id uint) bool {
	return s.repo.DeleteCategoria(id) == nil
}

// ================= PROMOCIONES =================
func (s *CatalogService) CreatePromocion(promocion models.Promocion) (models.Promocion, bool) {
	if promocion.Nombre == "" || promocion.Descuento <= 0 {
		return models.Promocion{}, false
	}
	err := s.repo.CreatePromocion(&promocion)
	return promocion, err == nil
}

func (s *CatalogService) GetPromociones() ([]models.Promocion, error) {
	return s.repo.GetPromociones()
}

func (s *CatalogService) GetPromocionByID(id uint) (models.Promocion, bool) {
	promo, err := s.repo.GetPromocionByID(id)
	return promo, err == nil
}

func (s *CatalogService) UpdatePromocion(promocion models.Promocion) (models.Promocion, bool) {
	err := s.repo.UpdatePromocion(&promocion)
	return promocion, err == nil
}

func (s *CatalogService) DeletePromocion(id uint) bool {
	return s.repo.DeletePromocion(id) == nil
}
