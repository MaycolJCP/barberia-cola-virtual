package services

import (
	"barberia-cola-virtual/internal/models"
	"barberia-cola-virtual/internal/repository"
	"strconv"
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

	// 1. Guardar el turno en la base de datos para obtener su ID único
	err := s.repo.Create(&turno)
	if err != nil {
		return models.Turno{}, false
	}

	// 2. REGLA DE NEGOCIO NO-CRUD: Calcular métricas en tiempo real de la cola virtual
	turnosActivos, err := s.repo.GetAll()
	if err != nil {
		return models.Turno{}, false
	}

	personasDelante := 0
	tiempoEsperaAcumulado := 0

	for _, t := range turnosActivos {
		// Contamos solo los turnos anteriores al nuestro que siguen esperando o están en atención
		if t.ID < turno.ID && (t.Estado == "ESPERANDO" || t.Estado == "EN_PROCESO") {
			personasDelante++
			if t.Servicio != nil {
				// Usa el campo 'Duracion' tal como lo definió tu compañero en su catálogo
				tiempoEsperaAcumulado += t.Servicio.Duracion
			} else {
				tiempoEsperaAcumulado += 20 // Tiempo estándar por defecto por si falta precargar la relación
			}
		}
	}

	// 3. Persistir de forma automática el Seguimiento de la Cola Virtual con datos reales
	seg := models.SeguimientoTurno{
		TurnoID:               turno.ID,
		Posicion:              personasDelante + 1,
		PersonasDelante:       personasDelante,
		TiempoEstimadoMinutos: tiempoEsperaAcumulado,
	}
	_ = s.repo.CreateSeguimiento(&seg)

	// 4. Generar la Alerta/Notificación con el estado inicial del usuario
	mensajeNotif := "Tu turno ha sido registrado con éxito. Tienes " + strconv.Itoa(personasDelante) + " personas delante. Tiempo estimado: " + strconv.Itoa(tiempoEsperaAcumulado) + " min."
	notif := models.Notificacion{
		TurnoID: turno.ID,
		Mensaje: mensajeNotif,
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
