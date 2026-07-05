package repository

import (
	"barberia-cola-virtual/internal/models"

	"gorm.io/gorm"
)

type SqlitePerfilRepository struct {
	db *gorm.DB
}

func NewPerfilRepository(db *gorm.DB) PerfilRepository {
	return &SqlitePerfilRepository{db: db}
}

func (r *SqlitePerfilRepository) CreateCliente(cliente *models.Cliente) error {
	return r.db.Create(cliente).Error
}

func (r *SqlitePerfilRepository) GetClientes() ([]models.Cliente, error) {
	var clientes []models.Cliente
	err := r.db.Find(&clientes).Error
	return clientes, err
}

func (r *SqlitePerfilRepository) GetClienteByID(id uint) (models.Cliente, error) {
	var cliente models.Cliente
	err := r.db.First(&cliente, id).Error
	return cliente, err
}

func (r *SqlitePerfilRepository) UpdateCliente(cliente *models.Cliente) error {
	return r.db.Save(cliente).Error
}

func (r *SqlitePerfilRepository) DeleteCliente(id uint) error {
	return r.db.Delete(&models.Cliente{}, id).Error
}

func (r *SqlitePerfilRepository) CreatePreferenciaPago(pref *models.PreferenciaPago) error {
	return r.db.Create(pref).Error
}

func (r *SqlitePerfilRepository) GetPreferenciasPagoByCliente(clienteID uint) ([]models.PreferenciaPago, error) {
	var prefs []models.PreferenciaPago
	err := r.db.Where("cliente_id = ?", clienteID).Find(&prefs).Error
	return prefs, err
}

func (r *SqlitePerfilRepository) CreatePreferenciaCliente(pref *models.PreferenciaCliente) error {
	return r.db.Create(pref).Error
}

func (r *SqlitePerfilRepository) GetPreferenciasClienteByCliente(clienteID uint) (models.PreferenciaCliente, error) {
	var pref models.PreferenciaCliente
	err := r.db.Where("cliente_id = ?", clienteID).First(&pref).Error
	return pref, err
}
