package services

import (
	"barberia-cola-virtual/internal/models"
	"barberia-cola-virtual/internal/repository"
)

// ================= TURNOS =================

func CreateTurno(turno models.Turno) (models.Turno, bool) {
	if turno.ClienteID <= 0 || turno.ServicioID <= 0 {
		return models.Turno{}, false
	}

	return repository.CreateTurno(turno), true
}

func GetTurnos() []models.Turno {
	return repository.GetTurnos()
}

func GetTurnoByID(id int) (models.Turno, bool) {
	return repository.GetTurnoByID(id)
}

func UpdateTurno(id int, turno models.Turno) (models.Turno, bool) {
	return repository.UpdateTurno(id, turno)
}

func DeleteTurno(id int) bool {
	return repository.DeleteTurno(id)
}

// ================= SEGUIMIENTO =================

func GetSeguimientosTurno() []models.SeguimientoTurno {
	return repository.GetSeguimientosTurno()
}

// ================= NOTIFICACIONES =================

func GetNotificaciones() []models.Notificacion {
	return repository.GetNotificaciones()
}