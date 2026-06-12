package models

type MetodoPago struct {
	ID        int    `json:"id"`
	ClienteID int   `json:"cliente_id"`
	Tipo      string `json:"tipo"`
	Titular   string `json:"titular"`
	Estado    string `json:"estado"`
}