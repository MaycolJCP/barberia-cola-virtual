package repository

import (
	"barberia-cola-virtual/internal/models"

	"gorm.io/gorm"
)

type SqliteAuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &SqliteAuthRepository{db: db}
}

func (r *SqliteAuthRepository) CreateUsuario(usuario *models.Usuario) error {
	return r.db.Create(usuario).Error
}

func (r *SqliteAuthRepository) GetUsuarioByCorreo(correo string) (models.Usuario, error) {
	var usuario models.Usuario
	err := r.db.Where("correo = ?", correo).First(&usuario).Error
	return usuario, err
}
