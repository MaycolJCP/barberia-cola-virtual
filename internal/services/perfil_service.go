package services

import (
	"barberia-cola-virtual/internal/models"
	"barberia-cola-virtual/internal/repository"
)

func CreateCliente(cliente models.Cliente) models.Cliente {
	return repository.CreateCliente(cliente)
}

func GetClientes() []models.Cliente {
	return repository.GetClientes()
}

func GetClienteByID(id int) (models.Cliente, bool) {
	return repository.GetClienteByID(id)
}

func UpdateCliente(id int, cliente models.Cliente) (models.Cliente, bool) {
	return repository.UpdateCliente(id, cliente)
}

func DeleteCliente(id int) bool {
	return repository.DeleteCliente(id)
}

// Servicios para Preferencias de Pago y Cliente
func CreatePreferenciaPago(preferencia models.PreferenciaPago) models.PreferenciaPago {
	return repository.CreatePreferenciaPago(preferencia)
}

func GetPreferenciasPago() []models.PreferenciaPago {
	return repository.GetPreferenciasPago()
}

func CreatePreferenciaCliente(preferencia models.PreferenciaCliente) models.PreferenciaCliente {
	return repository.CreatePreferenciaCliente(preferencia)
}

func GetPreferenciasCliente() []models.PreferenciaCliente {
	return repository.GetPreferenciasCliente()
}