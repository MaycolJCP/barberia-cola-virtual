package models

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model
	Nombre   string `json:"nombre" gorm:"not null"`
	Correo   string `json:"correo" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	Rol      string `json:"rol" gorm:"default:'CLIENTE'"`
}

type UsuarioResponse struct {
	ID     uint   `json:"id"`
	Nombre string `json:"nombre"`
	Correo string `json:"correo"`
	Rol    string `json:"rol"`
}
