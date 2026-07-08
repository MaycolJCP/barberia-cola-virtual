package repository

import (
	"barberia-cola-virtual/internal/models"

	"gorm.io/gorm"
)

// SqliteCatalogRepository es la implementación concreta del repositorio de catálogo.
//
// Aunque el nombre histórico dice "Sqlite", actualmente usa *gorm.DB.
// Eso significa que puede trabajar con PostgreSQL, SQLite u otro motor,
// dependiendo de la conexión que se inyecte desde main.go.
type SqliteCatalogRepository struct {
	// db es la conexión de GORM hacia la base de datos.
	// En el Hito 3 esta conexión apunta a PostgreSQL dentro de Docker.
	db *gorm.DB
}

// NewCatalogRepository crea el repositorio del catálogo.
//
// Devuelve la interfaz CatalogRepository, no la estructura concreta.
// Esto ayuda a mantener desacopladas las capas y permite usar mocks en tests.
func NewCatalogRepository(db *gorm.DB) CatalogRepository {
	return &SqliteCatalogRepository{db: db}
}

// ================= SERVICIOS =================

// CreateServicio guarda un nuevo servicio en la base de datos.
//
// GORM genera el INSERT correspondiente y asigna automáticamente el ID.
func (r *SqliteCatalogRepository) CreateServicio(servicio *models.Servicio) error {
	return r.db.Create(servicio).Error
}

// GetServicios obtiene todos los servicios registrados.
//
// GORM ejecuta una consulta equivalente a SELECT * FROM servicios.
func (r *SqliteCatalogRepository) GetServicios() ([]models.Servicio, error) {
	var servicios []models.Servicio
	err := r.db.Find(&servicios).Error
	return servicios, err
}

// GetServicioByID busca un servicio por su clave primaria.
func (r *SqliteCatalogRepository) GetServicioByID(id uint) (models.Servicio, error) {
	var servicio models.Servicio
	err := r.db.First(&servicio, id).Error
	return servicio, err
}

// UpdateServicio actualiza un servicio.
//
// Save de GORM crea o actualiza dependiendo de si el modelo tiene ID.
// En este caso se usa para actualizar un servicio existente.
func (r *SqliteCatalogRepository) UpdateServicio(servicio *models.Servicio) error {
	return r.db.Save(servicio).Error
}

// DeleteServicio elimina un servicio por ID.
//
// GORM ejecuta el DELETE sobre la tabla correspondiente.
func (r *SqliteCatalogRepository) DeleteServicio(id uint) error {
	return r.db.Delete(&models.Servicio{}, id).Error
}

// ================= CATEGORIAS =================

// CreateCategoria guarda una categoría de servicios.
func (r *SqliteCatalogRepository) CreateCategoria(categoria *models.CategoriaServicio) error {
	return r.db.Create(categoria).Error
}

// GetCategorias lista todas las categorías.
func (r *SqliteCatalogRepository) GetCategorias() ([]models.CategoriaServicio, error) {
	var categorias []models.CategoriaServicio
	err := r.db.Find(&categorias).Error
	return categorias, err
}

// GetCategoriaByID busca una categoría por ID.
func (r *SqliteCatalogRepository) GetCategoriaByID(id uint) (models.CategoriaServicio, error) {
	var categoria models.CategoriaServicio
	err := r.db.First(&categoria, id).Error
	return categoria, err
}

// UpdateCategoria actualiza una categoría.
func (r *SqliteCatalogRepository) UpdateCategoria(categoria *models.CategoriaServicio) error {
	return r.db.Save(categoria).Error
}

// DeleteCategoria elimina una categoría por ID.
func (r *SqliteCatalogRepository) DeleteCategoria(id uint) error {
	return r.db.Delete(&models.CategoriaServicio{}, id).Error
}

// ================= PROMOCIONES =================

// CreatePromocion guarda una promoción en la base de datos.
func (r *SqliteCatalogRepository) CreatePromocion(promocion *models.Promocion) error {
	return r.db.Create(promocion).Error
}

// GetPromociones lista todas las promociones.
func (r *SqliteCatalogRepository) GetPromociones() ([]models.Promocion, error) {
	var promociones []models.Promocion
	err := r.db.Find(&promociones).Error
	return promociones, err
}

// GetPromocionByID busca una promoción por ID.
func (r *SqliteCatalogRepository) GetPromocionByID(id uint) (models.Promocion, error) {
	var promocion models.Promocion
	err := r.db.First(&promocion, id).Error
	return promocion, err
}

// UpdatePromocion actualiza una promoción.
func (r *SqliteCatalogRepository) UpdatePromocion(promocion *models.Promocion) error {
	return r.db.Save(promocion).Error
}

// DeletePromocion elimina una promoción por ID.
func (r *SqliteCatalogRepository) DeletePromocion(id uint) error {
	return r.db.Delete(&models.Promocion{}, id).Error
}
