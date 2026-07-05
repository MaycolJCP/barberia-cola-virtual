package repository

import (
	"barberia-cola-virtual/internal/models"

	"gorm.io/gorm"
)

type ServicioGormRepository struct {
	db *gorm.DB
}

var globalDB *gorm.DB

// SetGormDB permite asignar la base de datos desde el main
func SetGormDB(db *gorm.DB) {
	globalDB = db
}

// NewServicioGormRepository inicializa el repositorio con su instancia de BD
func NewServicioGormRepository(db *gorm.DB) *ServicioGormRepository {
	return &ServicioGormRepository{db: db}
}

func (r *ServicioGormRepository) CreateServicio(servicio models.Servicio) models.Servicio {
	r.db.Create(&servicio)
	return servicio
}

func (r *ServicioGormRepository) GetServicios() []models.Servicio {
	var servicios []models.Servicio
	r.db.Find(&servicios)
	return servicios
}
