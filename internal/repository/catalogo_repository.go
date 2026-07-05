package repository

import (
	"barberia-cola-virtual/internal/models"

	"gorm.io/gorm"
)

type SqliteCatalogRepository struct {
	db *gorm.DB
}

func NewCatalogRepository(db *gorm.DB) CatalogRepository {
	return &SqliteCatalogRepository{db: db}
}

// ================= SERVICIOS =================
func (r *SqliteCatalogRepository) CreateServicio(servicio *models.Servicio) error {
	return r.db.Create(servicio).Error
}

func (r *SqliteCatalogRepository) GetServicios() ([]models.Servicio, error) {
	var servicios []models.Servicio
	err := r.db.Find(&servicios).Error
	return servicios, err
}

func (r *SqliteCatalogRepository) GetServicioByID(id uint) (models.Servicio, error) {
	var servicio models.Servicio
	err := r.db.First(&servicio, id).Error
	return servicio, err
}

func (r *SqliteCatalogRepository) UpdateServicio(servicio *models.Servicio) error {
	return r.db.Save(servicio).Error
}

func (r *SqliteCatalogRepository) DeleteServicio(id uint) error {
	return r.db.Delete(&models.Servicio{}, id).Error
}

// ================= CATEGORIAS =================
func (r *SqliteCatalogRepository) CreateCategoria(categoria *models.CategoriaServicio) error {
	return r.db.Create(categoria).Error
}

func (r *SqliteCatalogRepository) GetCategorias() ([]models.CategoriaServicio, error) {
	var categorias []models.CategoriaServicio
	err := r.db.Find(&categorias).Error
	return categorias, err
}

func (r *SqliteCatalogRepository) GetCategoriaByID(id uint) (models.CategoriaServicio, error) {
	var categoria models.CategoriaServicio
	err := r.db.First(&categoria, id).Error
	return categoria, err
}

func (r *SqliteCatalogRepository) UpdateCategoria(categoria *models.CategoriaServicio) error {
	return r.db.Save(categoria).Error
}

func (r *SqliteCatalogRepository) DeleteCategoria(id uint) error {
	return r.db.Delete(&models.CategoriaServicio{}, id).Error
}

// ================= PROMOCIONES =================
func (r *SqliteCatalogRepository) CreatePromocion(promocion *models.Promocion) error {
	return r.db.Create(promocion).Error
}

func (r *SqliteCatalogRepository) GetPromociones() ([]models.Promocion, error) {
	var promociones []models.Promocion
	err := r.db.Find(&promociones).Error
	return promociones, err
}

func (r *SqliteCatalogRepository) GetPromocionByID(id uint) (models.Promocion, error) {
	var promocion models.Promocion
	err := r.db.First(&promocion, id).Error
	return promocion, err
}

func (r *SqliteCatalogRepository) UpdatePromocion(promocion *models.Promocion) error {
	return r.db.Save(promocion).Error
}

func (r *SqliteCatalogRepository) DeletePromocion(id uint) error {
	return r.db.Delete(&models.Promocion{}, id).Error
}
