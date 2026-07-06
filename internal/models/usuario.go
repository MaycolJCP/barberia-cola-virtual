package models

import "gorm.io/gorm"

// Usuario representa la entidad del módulo de Perfil e IAM (Integrante C)
type Usuario struct {
	gorm.Model

	Username string `json:"username" gorm:"type:varchar(50);not null;uniqueIndex"`
	Password string `json:"password" gorm:"type:varchar(255);not null"` // <-- CORREGIDO: quitado el "-" para permitir el Decode
	Nombre   string `json:"nombre" gorm:"type:varchar(100);not null"`
	Correo   string `json:"email" gorm:"type:varchar(100);not null;uniqueIndex"`
	Rol      string `json:"role" gorm:"type:varchar(20);default:'CLIENTE';not null"`

	// Relación Has-Many: Un cliente puede solicitar muchos turnos en la barbería
	Turnos []Turno `json:"turnos,omitempty" gorm:"foreignKey:ClienteID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}

type UsuarioResponse struct {
	ID     uint   `json:"id"`
	Nombre string `json:"nombre"`
	Correo string `json:"correo"`
	Rol    string `json:"rol"`
}
