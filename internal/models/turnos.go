package models

import (
	"time"

	"gorm.io/gorm"
)

// Turno representa la entidad principal del módulo de turnos (Michael)
type Turno struct {
	gorm.Model
	ClienteID  uint      `json:"cliente_id" gorm:"not null"`
	ServicioID uint      `json:"servicio_id" gorm:"not null"`
	FechaHora  time.Time `json:"fecha_hora" gorm:"not null"`
	Estado     string    `json:"estado" gorm:"type:varchar(20);default:'ESPERANDO';not null"` // ESPERANDO, EN_PROCESO, ATENDIDO, CANCELADO

	// Relaciones Avanzadas GORM (Belongs-To)
	// Vincula tus llaves foráneas con los modelos de tus compañeros
	Cliente  *Usuario  `json:"cliente,omitempty" gorm:"foreignKey:ClienteID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Servicio *Servicio `json:"servicio,omitempty" gorm:"foreignKey:ServicioID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}

// SeguimientoTurno representa la métrica viva de la cola virtual
type SeguimientoTurno struct {
	gorm.Model
	TurnoID               uint `json:"turno_id" gorm:"not null;uniqueIndex"`
	Posicion              int  `json:"posicion" gorm:"not null"`
	PersonasDelante       int  `json:"personas_delante" gorm:"not null"`
	TiempoEstimadoMinutos int  `json:"tiempo_estimado_minutos" gorm:"not null"`

	Turno *Turno `json:"-" gorm:"foreignKey:TurnoID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// Notificacion representa las alertas generadas por el sistema de turnos
type Notificacion struct {
	gorm.Model
	TurnoID uint   `json:"turno_id" gorm:"not null"`
	Mensaje string `json:"mensaje" gorm:"type:text;not null"`
	Leido   bool   `json:"leido" gorm:"default:false;not null"`

	Turno *Turno `json:"-" gorm:"foreignKey:TurnoID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
