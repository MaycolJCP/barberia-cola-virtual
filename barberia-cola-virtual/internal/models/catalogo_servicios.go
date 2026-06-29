package models

type Servicio struct {
	ID          int     `json:"id"`
	CategoriaID int    `json:"categoria_id"`
	Nombre      string  `json:"nombre"`
	Descripcion string  `json:"descripcion"`
	Precio      float64 `json:"precio"`
	Duracion    int     `json:"duracion"`
}

type CategoriaServicio struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
}

type Promocion struct {
	ID         int     `json:"id"`
	ServicioID int    `json:"servicio_id"`
	Nombre     string  `json:"nombre"`
	Descuento  float64 `json:"descuento"`
}