package models

type Turno struct {
	ID         int    `json:"id"`
	ClienteID  int    `json:"cliente_id"`
	ServicioID int    `json:"servicio_id"`
	Estado     string `json:"estado"`
}

type SeguimientoTurno struct {
	ID                    int `json:"id"`
	TurnoID               int `json:"turno_id"`
	Posicion              int `json:"posicion"`
	PersonasDelante       int `json:"personas_delante"`
	TiempoEstimadoMinutos int `json:"tiempo_estimado_minutos"`
}

type Notificacion struct {
	ID      int    `json:"id"`
	TurnoID int    `json:"turno_id"`
	Mensaje string `json:"mensaje"`
	Leida   bool   `json:"leida"`
}