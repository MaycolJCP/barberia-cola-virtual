package repository

import (
	"barberia-cola-virtual/internal/models"

	"gorm.io/gorm"
)

// ServicioGormRepository maneja Servicio usando GORM.
type ServicioGormRepository struct {
	DB *gorm.DB
}

// CrearServicio guarda un servicio en SQLite.
func (r *ServicioGormRepository) CrearServicio(servicio models.Servicio) error {
	return r.DB.Create(&servicio).Error
}

// ObtenerServicios devuelve todos los servicios.
func (r *ServicioGormRepository) ObtenerServicios() ([]models.Servicio, error) {
	var servicios []models.Servicio

	err := r.DB.Find(&servicios).Error

	return servicios, err
}