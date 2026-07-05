package models

import "gorm.io/gorm"

type Cliente struct {
	gorm.Model
	Nombre   string `json:"nombre" gorm:"not null"`
	Correo   string `json:"correo" gorm:"unique;not null"`
	Telefono string `json:"telefono"`
}

type PreferenciaPago struct {
	gorm.Model
	ClienteID uint   `json:"cliente_id" gorm:"not null;index"`
	TipoPago  string `json:"tipo_pago" gorm:"not null"`
}

type PreferenciaCliente struct {
	gorm.Model
	ClienteID         uint   `json:"cliente_id" gorm:"not null;index"`
	BarberoPreferido  string `json:"barbero_preferido"`
	ServicioFrecuente string `json:"servicio_frecuente"`
}
