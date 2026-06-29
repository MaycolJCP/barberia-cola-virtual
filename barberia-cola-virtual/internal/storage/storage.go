package storage

import "barberia-cola-virtual/internal/models"

//autenticación y autorización
var Usuarios []models.Usuario

// MÓDULO MI PERFIL
var Clientes []models.Cliente
var PreferenciasPago []models.PreferenciaPago
var PreferenciasCliente []models.PreferenciaCliente

// MÓDULO CATÁLOGO Y SERVICIOS
var Servicios []models.Servicio
var CategoriasServicio []models.CategoriaServicio
var Promociones []models.Promocion

// MÓDULO TURNOS
var Turnos []models.Turno
var SeguimientosTurno []models.SeguimientoTurno
var Notificaciones []models.Notificacion