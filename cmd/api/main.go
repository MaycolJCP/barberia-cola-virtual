package main

import (
	"log"
	"net/http"

	"barberia-cola-virtual/internal/db"
	"barberia-cola-virtual/internal/handlers"
	"barberia-cola-virtual/internal/middleware"
	"barberia-cola-virtual/internal/repository"
	"barberia-cola-virtual/internal/services"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func main() {
	// 1. Inicializar Base de Datos en SQLite
	database := db.InitDB("barberia.db")
	log.Println("¡Persistencia SQLite con GORM inicializada con éxito!")

	// 2. Instanciar Repositorios
	authRepo := repository.NewAuthRepository(database)
	catalogRepo := repository.NewCatalogRepository(database)
	perfilRepo := repository.NewPerfilRepository(database)
	turnosRepo := repository.NewTurnosRepository(database)

	// 3. Instanciar Servicios
	authService := services.NewAuthService(authRepo)
	catalogService := services.NewCatalogService(catalogRepo)
	perfilService := services.NewPerfilService(perfilRepo)
	turnoService := services.NewTurnoService(turnosRepo)

	// 4. Instanciar Handlers inyectando los servicios correspondientes
	authHandler := handlers.NewAuthHandler(authService)
	catalogoHandler := handlers.NewCatalogoHandler(catalogService)
	perfilHandler := handlers.NewPerfilHandler(perfilService)
	turnosHandler := handlers.NewTurnosHandler(turnoService)

	// 5. Configuración del Router
	r := chi.NewRouter()
	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Recoverer)

	r.Route("/api/v1", func(r chi.Router) {

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("API Barberia Cola Virtual funcionando con GORM y SQLite"))
		})

		// Módulo Autenticación y Registro
		r.Route("/auth", func(router chi.Router) {
			router.Post("/register", authHandler.Register)
			router.Post("/login", authHandler.Login)

			router.Group(func(auth chi.Router) {
				auth.Use(middleware.AuthMiddleware)
				auth.Get("/me", authHandler.Me)
			})
		})

		// Módulo Mi Perfil
		r.Route("/clientes", func(router chi.Router) {
			router.Use(middleware.AuthMiddleware)

			router.Post("/", perfilHandler.CreateCliente)
			router.Get("/", perfilHandler.GetClientes)
			router.Get("/{id}", perfilHandler.GetClienteByID)
			router.Put("/{id}", perfilHandler.UpdateCliente)
			router.Delete("/{id}", perfilHandler.DeleteCliente)
		})

		r.Route("/preferencias-pago", func(router chi.Router) {
			router.Use(middleware.AuthMiddleware)

			router.Post("/", perfilHandler.CreatePreferenciaPago)
			router.Get("/", perfilHandler.GetPreferenciasPago)
		})

		r.Route("/preferencias-cliente", func(router chi.Router) {
			router.Use(middleware.AuthMiddleware)

			router.Post("/", perfilHandler.CreatePreferenciaCliente)
			router.Get("/", perfilHandler.GetPreferenciasCliente)
		})

		// Módulo Catálogo y Selección de Servicios
		r.Route("/servicios", func(router chi.Router) {
			router.Get("/", catalogoHandler.GetServicios)        // CORREGIDO: router. en vez de r.
			router.Get("/{id}", catalogoHandler.GetServicioByID) // CORREGIDO: router. en vez de r.

			router.Group(func(admin chi.Router) {
				admin.Use(middleware.AuthMiddleware)
				admin.Use(middleware.AdminOnly)

				admin.Post("/", catalogoHandler.CreateServicio)
				admin.Put("/{id}", catalogoHandler.UpdateServicio)
				admin.Delete("/{id}", catalogoHandler.DeleteServicio)
			})
		})

		// Módulo Categorías de Servicio
		r.Route("/categorias-servicio", func(router chi.Router) {
			router.Get("/", catalogoHandler.GetCategoriasServicio)        // CORREGIDO: router. en vez de r.
			router.Get("/{id}", catalogoHandler.GetCategoriaServicioByID) // CORREGIDO: router. en vez de r.

			router.Group(func(admin chi.Router) {
				admin.Use(middleware.AuthMiddleware)
				admin.Use(middleware.AdminOnly)

				admin.Post("/", catalogoHandler.CreateCategoriaServicio)
				admin.Put("/{id}", catalogoHandler.UpdateCategoriaServicio)
				admin.Delete("/{id}", catalogoHandler.DeleteCategoriaServicio)
			})
		})

		// Módulo Gestión de Promociones y Descuentos
		r.Route("/promociones", func(router chi.Router) {
			router.Get("/", catalogoHandler.GetPromociones)       // CORREGIDO: router. en vez de r.
			router.Get("/{id}", catalogoHandler.GetPromocionByID) // CORREGIDO: router. en vez de r.

			router.Group(func(admin chi.Router) {
				admin.Use(middleware.AuthMiddleware)
				admin.Use(middleware.AdminOnly)

				admin.Post("/", catalogoHandler.CreatePromocion)
				admin.Put("/{id}", catalogoHandler.UpdatePromocion)
				admin.Delete("/{id}", catalogoHandler.DeletePromocion)
			})
		})

		// Módulo Mis Turnos y Seguimiento
		r.Route("/turnos", func(router chi.Router) {
			router.Use(middleware.AuthMiddleware)

			router.Post("/", turnosHandler.CreateTurno)
			router.Get("/", turnosHandler.GetTurnos)
			router.Get("/{id}", turnosHandler.GetTurnoByID)
			router.Put("/{id}", turnosHandler.UpdateTurno)
			router.Delete("/{id}", turnosHandler.DeleteTurno)
		})

		r.Route("/seguimientos-turno", func(router chi.Router) {
			router.Use(middleware.AuthMiddleware)

			router.Get("/", turnosHandler.GetSeguimientosTurno)
		})

		r.Route("/notificaciones", func(router chi.Router) {
			router.Use(middleware.AuthMiddleware)

			router.Get("/", turnosHandler.GetNotificaciones)
		})
	})

	log.Println("Servidor escuchando en http://localhost:8080")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("Error crítico al levantar el servidor HTTP:", err)
	}
}
