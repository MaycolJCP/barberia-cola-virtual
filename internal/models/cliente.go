package models

type Cliente struct {
	ID           int    `json:"id"`
	Nombre       string `json:"nombre"`
	Correo       string `json:"correo"`
	Telefono     string `json:"telefono"`
	Direccion    string `json:"direccion"`
	Genero       string `json:"genero"`
	UltimaVisita string `json:"ultima_visita"`
}
