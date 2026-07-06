package models

import "gorm.io/gorm"

type CategoriaServicio struct {
	gorm.Model
	Nombre      string `json:"nombre" gorm:"unique;not null"`
	Descripcion string `json:"descripcion"`
}

type Servicio struct {
	gorm.Model
	CategoriaID uint    `json:"categoria_id" gorm:"not null;index"`
	Nombre      string  `json:"nombre" gorm:"not null"`
	Descripcion string  `json:"descripcion"`
	Precio      float64 `json:"precio" gorm:"not null"`
	Duracion    int     `json:"duracion" gorm:"not null"` // En minutos (Se usa en tu regla de negocio)

	// RELACIÓN AVANZADA (Has-Many): Un servicio está vinculado a muchos turnos de la cola
	Turnos []Turno `json:"turnos,omitempty" gorm:"foreignKey:ServicioID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}

type Promocion struct {
	gorm.Model
	ServicioID uint    `json:"servicio_id" gorm:"not null;index"`
	Nombre     string  `json:"nombre" gorm:"not null"`
	Descuento  float64 `json:"descuento" gorm:"not null"`
}
