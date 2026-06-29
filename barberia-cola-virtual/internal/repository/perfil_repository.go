package repository

import (
	"barberia-cola-virtual/internal/models"
	"barberia-cola-virtual/internal/storage"
)

func CreateCliente(cliente models.Cliente) models.Cliente {
	cliente.ID = len(storage.Clientes) + 1
	storage.Clientes = append(storage.Clientes, cliente)
	return cliente
}

func GetClientes() []models.Cliente {
	return storage.Clientes
}

func GetClienteByID(id int) (models.Cliente, bool) {
	for _, cliente := range storage.Clientes {
		if cliente.ID == id {
			return cliente, true
		}
	}
	return models.Cliente{}, false
}

func UpdateCliente(id int, updatedCliente models.Cliente) (models.Cliente, bool) {
	for i, cliente := range storage.Clientes {
		if cliente.ID == id {
			updatedCliente.ID = cliente.ID
			storage.Clientes[i] = updatedCliente
			return updatedCliente, true
		}
	}
	return models.Cliente{}, false
}

func DeleteCliente(id int) bool {
	for i, cliente := range storage.Clientes {
		if cliente.ID == id {
			storage.Clientes = append(storage.Clientes[:i], storage.Clientes[i+1:]...)
			return true
		}
	}
	return false
}

func CreatePreferenciaPago(preferencia models.PreferenciaPago) models.PreferenciaPago {
	preferencia.ID = len(storage.PreferenciasPago) + 1
	storage.PreferenciasPago = append(storage.PreferenciasPago, preferencia)
	return preferencia
}

func GetPreferenciasPago() []models.PreferenciaPago {
	return storage.PreferenciasPago
}

func CreatePreferenciaCliente(preferencia models.PreferenciaCliente) models.PreferenciaCliente {
	preferencia.ID = len(storage.PreferenciasCliente) + 1
	storage.PreferenciasCliente = append(storage.PreferenciasCliente, preferencia)
	return preferencia
}

func GetPreferenciasCliente() []models.PreferenciaCliente {
	return storage.PreferenciasCliente
}