package models

import "gorm.io/gorm"

type Turno struct {
	gorm.Model
	ClienteID  uint   `json:"cliente_id" gorm:"not null;index"`
	ServicioID uint   `json:"servicio_id" gorm:"not null;index"`
	Estado     string `json:"estado" gorm:"default:'ESPERANDO'"`
}

type SeguimientoTurno struct {
	gorm.Model
	TurnoID               uint `json:"turno_id" gorm:"not null;uniqueIndex"`
	Posicion              int  `json:"posicion" gorm:"not null"`
	PersonasDelante       int  `json:"personas_delante" gorm:"not null"`
	TiempoEstimadoMinutos int  `json:"tiempo_estimado_minutos" gorm:"not null"`
}

type Notificacion struct {
	gorm.Model
	TurnoID uint   `json:"turno_id" gorm:"not null;index"`
	Mensaje string `json:"mensaje" gorm:"not null"`
	Leida   bool   `json:"leida" gorm:"default:false"`
}
