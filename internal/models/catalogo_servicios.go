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
	Duracion    int     `json:"duracion" gorm:"not null"` // En minutos
}

type Promocion struct {
	gorm.Model
	ServicioID uint    `json:"servicio_id" gorm:"not null;index"`
	Nombre     string  `json:"nombre" gorm:"not null"`
	Descuento  float64 `json:"descuento" gorm:"not null"`
}
