package repository

import (
	"barberia-cola-virtual/internal/models"

	"gorm.io/gorm"
)

// 🟢 Cambiado a ITurnosRepository para evitar choques de nombres (redeclared) en el paquete
type ITurnosRepository interface {
	Create(turno *models.Turno) error
	GetAll() ([]models.Turno, error)
	GetByID(id uint) (models.Turno, error)
	Update(turno *models.Turno) error
	Delete(id uint) error
	CreateSeguimiento(seg *models.SeguimientoTurno) error
	GetSeguimientos() ([]models.SeguimientoTurno, error)
	CreateNotificacion(notif *models.Notificacion) error
	GetNotificaciones() ([]models.Notificacion, error)
}

type SqliteTurnosRepository struct {
	db *gorm.DB
}

// El constructor ahora retorna la interfaz renombrada de forma limpia
func NewTurnosRepository(db *gorm.DB) ITurnosRepository {
	return &SqliteTurnosRepository{db: db}
}

func (r *SqliteTurnosRepository) Create(turno *models.Turno) error {
	return r.db.Create(turno).Error
}

// GetAll carga de forma relacional estricta al Cliente y al Servicio asociados
func (r *SqliteTurnosRepository) GetAll() ([]models.Turno, error) {
	var turnos []models.Turno
	// Preload lee de forma automática los datos correspondientes usando las llaves foráneas
	err := r.db.Preload("Cliente").Preload("Servicio").Find(&turnos).Error
	return turnos, err
}

// GetByID carga de forma relacional estricta al Cliente y al Servicio asociados
func (r *SqliteTurnosRepository) GetByID(id uint) (models.Turno, error) {
	var turno models.Turno
	// Preload evita que los objetos anidados retornen null en las consultas individuales
	err := r.db.Preload("Cliente").Preload("Servicio").First(&turno, id).Error
	return turno, err
}

// 🟢 Solución al PUT: Forzamos un UPDATE SQL real basándonos en el ID inyectado por tu Handler
func (r *SqliteTurnosRepository) Update(turno *models.Turno) error {
	return r.db.Model(turno).Updates(turno).Error
}

func (r *SqliteTurnosRepository) Delete(id uint) error {
	return r.db.Delete(&models.Turno{}, id).Error
}

func (r *SqliteTurnosRepository) CreateSeguimiento(seg *models.SeguimientoTurno) error {
	return r.db.Create(seg).Error
}

// 🟢 Optimizado con Preload condicional para alimentar las métricas de tu cola virtual
func (r *SqliteTurnosRepository) GetSeguimientos() ([]models.SeguimientoTurno, error) {
	var seguimientos []models.SeguimientoTurno
	// Cargamos de forma anidada el turno completo con su cliente y el servicio requerido
	err := r.db.Preload("Turno").Preload("Turno.Cliente").Preload("Turno.Servicio").Find(&seguimientos).Error
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
