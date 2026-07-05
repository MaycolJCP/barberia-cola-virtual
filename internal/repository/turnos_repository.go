package repository

import (
	"barberia-cola-virtual/internal/models"

	"gorm.io/gorm"
)

type SqliteTurnosRepository struct {
	db *gorm.DB
}

func NewTurnosRepository(db *gorm.DB) TurnosRepository {
	return &SqliteTurnosRepository{db: db}
}

func (r *SqliteTurnosRepository) Create(turno *models.Turno) error {
	return r.db.Create(turno).Error
}

// MODIFICAR: Ahora carga de forma relacional al Cliente y al Servicio asociados
func (r *SqliteTurnosRepository) GetAll() ([]models.Turno, error) {
	var turnos []models.Turno
	// Preload lee de forma automática los datos correspondientes usando las llaves foráneas
	err := r.db.Preload("Cliente").Preload("Servicio").Find(&turnos).Error
	return turnos, err
}

// MODIFICAR: Ahora carga de forma relacional al Cliente y al Servicio asociados
func (r *SqliteTurnosRepository) GetByID(id uint) (models.Turno, error) {
	var turno models.Turno
	// Preload evita que los objetos anidados retornen null en las consultas individuales
	err := r.db.Preload("Cliente").Preload("Servicio").First(&turno, id).Error
	return turno, err
}

func (r *SqliteTurnosRepository) Update(turno *models.Turno) error {
	return r.db.Save(turno).Error
}

func (r *SqliteTurnosRepository) Delete(id uint) error {
	return r.db.Delete(&models.Turno{}, id).Error
}

func (r *SqliteTurnosRepository) CreateSeguimiento(seg *models.SeguimientoTurno) error {
	return r.db.Create(seg).Error
}

func (r *SqliteTurnosRepository) GetSeguimientos() ([]models.SeguimientoTurno, error) {
	var seguimientos []models.SeguimientoTurno
	err := r.db.Find(&seguimientos).Error
	return seguimientos, err
}

func (r *SqliteTurnosRepository) CreateNotificacion(notif *models.Notificacion) error {
	return r.db.Create(notif).Error
}

func (r *SqliteTurnosRepository) GetNotificaciones() ([]models.Notificacion, error) {
	var notificaciones []models.Notificacion
	err := r.db.Find(&notificaciones).Error
	return notificaciones, err
}
