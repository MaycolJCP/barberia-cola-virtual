package models

import (
	"time"

	"gorm.io/gorm"
)

// Turno representa la entidad core del módulo de turnos con sus relaciones
type Turno struct {
	gorm.Model

	// Relación Belongs-To con el Cliente/Usuario
	ClienteID uint     `json:"cliente_id" gorm:"not null;index"`
	Cliente   *Usuario `json:"cliente,omitempty" gorm:"foreignKey:ClienteID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`

	// Relación Belongs-To con el Servicio
	ServicioID uint      `json:"servicio_id" gorm:"not null;index"`
	Servicio   *Servicio `json:"servicio,omitempty" gorm:"foreignKey:ServicioID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`

	Estado            string    `json:"estado" gorm:"type:varchar(20);default:'ESPERANDO';not null"`
	FechaHora         time.Time `json:"fecha_hora" gorm:"not null"`
	TiempoEstimadoMin int       `json:"tiempo_estimado_min" gorm:"default:0"`
}

// MANTENER: Entidad para el seguimiento de la cola virtual
type SeguimientoTurno struct {
	gorm.Model
	TurnoID               uint `json:"turno_id" gorm:"not null;uniqueIndex"`
	Posicion              int  `json:"posicion" gorm:"not null"`
	PersonasDelante       int  `json:"personas_delante" gorm:"not null"`
	TiempoEstimadoMinutos int  `json:"tiempo_estimado_minutos" gorm:"not null"`
}

// MANTENER: Entidad para el sistema de alertas al usuario
type Notificacion struct {
	gorm.Model
	TurnoID uint   `json:"turno_id" gorm:"not null;index"`
	Mensaje string `json:"mensaje" gorm:"not null"`
	Leida   bool   `json:"leida" gorm:"default:false"`
}
