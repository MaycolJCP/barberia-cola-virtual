package storage

import "barberia-cola-virtual/internal/models"

var Clientes = []models.Cliente{
	{
		ID:           1,
		Nombre:       "Danny Zambrano",
		Correo:       "danny@gmail.com",
		Telefono:     "0999999999",
		Direccion:    "Manta, Centro",
		Genero:       "Masculino",
		UltimaVisita: "2026-06-01",
	},

	{
		ID:           2,
		Nombre:       "Michael Cedeño",
		Correo:       "michael@gmail.com",
		Telefono:     "0988888888",
		Direccion:    "Manta, Tarqui",
		Genero:       "Masculino",
		UltimaVisita: "2026-05-20",
	},
}
