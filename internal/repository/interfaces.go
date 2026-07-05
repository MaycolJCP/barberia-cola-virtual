package repository

import "barberia-cola-virtual/internal/models"

type AuthRepository interface {
	CreateUsuario(usuario *models.Usuario) error
	GetUsuarioByCorreo(correo string) (models.Usuario, error)
}

type CatalogRepository interface {
	CreateServicio(servicio *models.Servicio) error
	GetServicios() ([]models.Servicio, error)
	GetServicioByID(id uint) (models.Servicio, error)
	UpdateServicio(servicio *models.Servicio) error
	DeleteServicio(id uint) error

	CreateCategoria(categoria *models.CategoriaServicio) error
	GetCategorias() ([]models.CategoriaServicio, error)
	GetCategoriaByID(id uint) (models.CategoriaServicio, error)
	UpdateCategoria(categoria *models.CategoriaServicio) error
	DeleteCategoria(id uint) error

	CreatePromocion(promocion *models.Promocion) error
	GetPromociones() ([]models.Promocion, error)
	GetPromocionByID(id uint) (models.Promocion, error)
	UpdatePromocion(promocion *models.Promocion) error
	DeletePromocion(id uint) error
}

type PerfilRepository interface {
	CreateCliente(cliente *models.Cliente) error
	GetClientes() ([]models.Cliente, error)
	GetClienteByID(id uint) (models.Cliente, error)
	UpdateCliente(cliente *models.Cliente) error
	DeleteCliente(id uint) error

	CreatePreferenciaPago(pref *models.PreferenciaPago) error
	GetPreferenciasPagoByCliente(clienteID uint) ([]models.PreferenciaPago, error)
	CreatePreferenciaCliente(pref *models.PreferenciaCliente) error
	GetPreferenciasClienteByCliente(clienteID uint) (models.PreferenciaCliente, error)
}

type TurnosRepository interface {
	Create(turno *models.Turno) error
	GetAll() ([]models.Turno, error)
	GetByID(id uint) (models.Turno, error)
	Update(turno *models.Turno) error
	Delete(id uint) error

	CreateSeguimiento(seg *models.SeguimientoTurno) error
	GetSeguimientos() ([]models.SeguimientoTurno, error)
	CreateNotificacion(notif *models.Notificacion) error
	GetNotificaciones() ([]models.Notificacion, error)
}
