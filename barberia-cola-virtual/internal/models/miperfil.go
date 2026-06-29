package models

type Cliente struct {
	ID       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Correo   string `json:"correo"`
	Telefono string `json:"telefono"`
}

type PreferenciaPago struct {
	ID        int    `json:"id"`
	ClienteID int    `json:"cliente_id"`
	TipoPago  string `json:"tipo_pago"`
}

type PreferenciaCliente struct {
	ID                int    `json:"id"`
	ClienteID         int    `json:"cliente_id"`
	BarberoPreferido  string `json:"barbero_preferido"`
	ServicioFrecuente string `json:"servicio_frecuente"`
}