package services

import (
	"barberia-cola-virtual/internal/models"
	"barberia-cola-virtual/internal/repository"
)

type TurnoService struct {
	repo repository.TurnosRepository
}

func NewTurnoService(repo repository.TurnosRepository) *TurnoService {
	return &TurnoService{repo: repo}
}

func (s *TurnoService) CreateTurno(turno models.Turno) (models.Turno, bool) {
	if turno.ClienteID <= 0 || turno.ServicioID <= 0 {
		return models.Turno{}, false
	}
	turno.Estado = "ESPERANDO"

	err := s.repo.Create(&turno)
	if err != nil {
		return models.Turno{}, false
	}

	// Lógica automática de seguimiento y notificaciones persistidas en base de datos
	seg := models.SeguimientoTurno{
		TurnoID:               turno.ID,
		Posicion:              1, // Ejemplo de inicialización
		PersonasDelante:       0,
		TiempoEstimadoMinutos: 15,
	}
	_ = s.repo.CreateSeguimiento(&seg)

	notif := models.Notificacion{
		TurnoID: turno.ID,
		Mensaje: "Tu turno ha sido registrado con éxito.",
	}
	_ = s.repo.CreateNotificacion(&notif)

	return turno, true
}

func (s *TurnoService) GetTurnos() ([]models.Turno, error) {
	return s.repo.GetAll()
}

func (s *TurnoService) GetTurnoByID(id uint) (models.Turno, bool) {
	turno, err := s.repo.GetByID(id)
	return turno, err == nil
}

func (s *TurnoService) UpdateTurno(turno models.Turno) (models.Turno, bool) {
	err := s.repo.Update(&turno)
	return turno, err == nil
}

func (s *TurnoService) DeleteTurno(id uint) bool {
	return s.repo.Delete(id) == nil
}

func (s *TurnoService) GetSeguimientosTurno() ([]models.SeguimientoTurno, error) {
	return s.repo.GetSeguimientos()
}

func (s *TurnoService) GetNotificaciones() ([]models.Notificacion, error) {
	return s.repo.GetNotificaciones()
}
