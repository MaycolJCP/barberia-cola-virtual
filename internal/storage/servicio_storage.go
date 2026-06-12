package storage

import "barberia-cola-virtual/internal/models"

var Servicios = []models.Servicio{
	{
		ID:          1,
		Nombre:      "Corte Clasico",
		Descripcion: "Corte tradicional de cabello",
		Precio:      5,
		Duracion:    30,
	},
	{
		ID:          2,
		Nombre:      "Barba",
		Descripcion: "Perfilado y arreglo de barba",
		Precio:      3,
		Duracion:    15,
	},
	{
		ID:          3,
		Nombre:      "Corte + Barba",
		Descripcion: "Combo completo de corte y barba",
		Precio:      8,
		Duracion:    45,
	},
}