package models

type Servicio struct {
	ID          int     `json:"id"`
	Nombre      string  `json:"nombre"`
	Descripcion string  `json:"descripcion"`
	Precio      float64 `json:"precio"`
	Duracion    int     `json:"duracion"`
}