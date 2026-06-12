package models

import "time"

type Turno struct {
	ID                    int       `json:"id"`
	ClienteID             int       `json:"cliente_id"`
	ServicioID            int       `json:"servicio_id"`
	Estado                string    `json:"estado"`
	Posicion              int       `json:"posicion"`
	PersonasDelante       int       `json:"personas_delante"`
	TiempoEstimadoMinutos int       `json:"tiempo_estimado_minutos"`
	HoraEstimada          time.Time `json:"hora_estimada"`
}