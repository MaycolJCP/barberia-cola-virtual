package repository

import (
	"barberia-cola-virtual/internal/models"
	"barberia-cola-virtual/internal/storage"
)

// ================= TURNOS =================

func CreateTurno(turno models.Turno) models.Turno {
	turno.ID = len(storage.Turnos) + 1
	turno.Estado = "ESPERANDO"

	storage.Turnos = append(storage.Turnos, turno)

	seguimiento := models.SeguimientoTurno{
		ID:                    len(storage.SeguimientosTurno) + 1,
		TurnoID:               turno.ID,
		Posicion:              len(storage.Turnos),
		PersonasDelante:       len(storage.Turnos) - 1,
		TiempoEstimadoMinutos: (len(storage.Turnos) - 1) * 15,
	}

	storage.SeguimientosTurno = append(storage.SeguimientosTurno, seguimiento)

	notificacion := models.Notificacion{
		ID:      len(storage.Notificaciones) + 1,
		TurnoID: turno.ID,
		Mensaje: "Tu turno fue registrado correctamente",
		Leida:   false,
	}

	storage.Notificaciones = append(storage.Notificaciones, notificacion)

	return turno
}

func GetTurnos() []models.Turno {
	return storage.Turnos
}

func GetTurnoByID(id int) (models.Turno, bool) {
	for _, turno := range storage.Turnos {
		if turno.ID == id {
			return turno, true
		}
	}

	return models.Turno{}, false
}

func UpdateTurno(id int, updatedTurno models.Turno) (models.Turno, bool) {
	for i, turno := range storage.Turnos {
		if turno.ID == id {
			updatedTurno.ID = turno.ID
			storage.Turnos[i] = updatedTurno
			return updatedTurno, true
		}
	}

	return models.Turno{}, false
}

func DeleteTurno(id int) bool {
	for i, turno := range storage.Turnos {
		if turno.ID == id {
			storage.Turnos = append(storage.Turnos[:i], storage.Turnos[i+1:]...)
			return true
		}
	}

	return false
}

// ================= SEGUIMIENTO =================

func GetSeguimientosTurno() []models.SeguimientoTurno {
	return storage.SeguimientosTurno
}

// ================= NOTIFICACIONES =================

func GetNotificaciones() []models.Notificacion {
	return storage.Notificaciones
}