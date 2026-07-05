package services

import (
	"barberia-cola-virtual/internal/models"
	"barberia-cola-virtual/internal/repository"
)

type PerfilService struct {
	repo repository.PerfilRepository
}

func NewPerfilService(repo repository.PerfilRepository) *PerfilService {
	return &PerfilService{repo: repo}
}

func (s *PerfilService) CreateCliente(cliente models.Cliente) (models.Cliente, bool) {
	if cliente.Nombre == "" || cliente.Correo == "" {
		return models.Cliente{}, false
	}
	err := s.repo.CreateCliente(&cliente)
	return cliente, err == nil
}

func (s *PerfilService) GetClientes() ([]models.Cliente, error) {
	return s.repo.GetClientes()
}

func (s *PerfilService) GetClienteByID(id uint) (models.Cliente, bool) {
	cli, err := s.repo.GetClienteByID(id)
	return cli, err == nil
}

func (s *PerfilService) UpdateCliente(cliente models.Cliente) (models.Cliente, bool) {
	err := s.repo.UpdateCliente(&cliente)
	return cliente, err == nil
}

func (s *PerfilService) DeleteCliente(id uint) bool {
	return s.repo.DeleteCliente(id) == nil
}

func (s *PerfilService) CreatePreferenciaPago(pref models.PreferenciaPago) (models.PreferenciaPago, bool) {
	err := s.repo.CreatePreferenciaPago(&pref)
	return pref, err == nil
}

func (s *PerfilService) GetPreferenciasPago(clienteID uint) ([]models.PreferenciaPago, error) {
	return s.repo.GetPreferenciasPagoByCliente(clienteID)
}

func (s *PerfilService) CreatePreferenciaCliente(pref models.PreferenciaCliente) (models.PreferenciaCliente, bool) {
	err := s.repo.CreatePreferenciaCliente(&pref)
	return pref, err == nil
}

func (s *PerfilService) GetPreferenciasCliente(clienteID uint) (models.PreferenciaCliente, bool) {
	pref, err := s.repo.GetPreferenciasClienteByCliente(clienteID)
	return pref, err == nil
}
