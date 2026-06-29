package main

import (
	"log"
	"net/http"

	"barberia-cola-virtual/internal/handlers"
	"barberia-cola-virtual/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("API Barberia Cola Virtual funcionando"))
		})

		// Módulo Autenticación y Registro
		r.Route("/auth", func(router chi.Router) {
	router.Post("/register", handlers.Register)
	router.Post("/login", handlers.Login)

	router.Group(func(auth chi.Router) {
		auth.Use(middleware.AuthMiddleware)
		auth.Get("/me", handlers.Me)
	})
})

		// Módulo Mi Perfil
		r.Route("/clientes", func(router chi.Router) {
		router.Use(middleware.AuthMiddleware)

			router.Post("/", handlers.CreateCliente)
			router.Get("/", handlers.GetClientes)
			router.Get("/{id}", handlers.GetClienteByID)
			router.Put("/{id}", handlers.UpdateCliente)
			router.Delete("/{id}", handlers.DeleteCliente)
})

		r.Route("/preferencias-pago", func(router chi.Router) {
		router.Use(middleware.AuthMiddleware)

			router.Post("/", handlers.CreatePreferenciaPago)
			router.Get("/", handlers.GetPreferenciasPago)
})
		r.Route("/preferencias-cliente", func(router chi.Router) {
		router.Use(middleware.AuthMiddleware)

			router.Post("/", handlers.CreatePreferenciaCliente)
			router.Get("/", handlers.GetPreferenciasCliente)
})

		// Módulo Catálogo y Selección de Servicios
		r.Route("/servicios", func(router chi.Router) {
			router.Get("/", handlers.GetServicios)
			router.Get("/{id}", handlers.GetServicioByID)

			router.Group(func(admin chi.Router) {
			admin.Use(middleware.AuthMiddleware)
			admin.Use(middleware.AdminOnly)

			admin.Post("/", handlers.CreateServicio)
			admin.Put("/{id}", handlers.UpdateServicio)
			admin.Delete("/{id}", handlers.DeleteServicio)
	})
})
		// Módulo Gestión de Promociones y Descuentos
		r.Route("/categorias-servicio", func(router chi.Router) {
			router.Get("/", handlers.GetCategoriasServicio)
			router.Get("/{id}", handlers.GetCategoriaServicioByID)

			router.Group(func(admin chi.Router) {
			admin.Use(middleware.AuthMiddleware)
			admin.Use(middleware.AdminOnly)

			admin.Post("/", handlers.CreateCategoriaServicio)
			admin.Put("/{id}", handlers.UpdateCategoriaServicio)
			admin.Delete("/{id}", handlers.DeleteCategoriaServicio)
	})
})

			r.Route("/promociones", func(router chi.Router) {
			router.Get("/", handlers.GetPromociones)
			router.Get("/{id}", handlers.GetPromocionByID)

			router.Group(func(admin chi.Router) {
			admin.Use(middleware.AuthMiddleware)
			admin.Use(middleware.AdminOnly)

			admin.Post("/", handlers.CreatePromocion)
			admin.Put("/{id}", handlers.UpdatePromocion)
			admin.Delete("/{id}", handlers.DeletePromocion)
	})
})

		// Módulo Mis Turnos y Seguimiento
		r.Route("/turnos", func(router chi.Router) {
		router.Use(middleware.AuthMiddleware)

		router.Post("/", handlers.CreateTurno)
		router.Get("/", handlers.GetTurnos)
		router.Get("/{id}", handlers.GetTurnoByID)
		router.Put("/{id}", handlers.UpdateTurno)
		router.Delete("/{id}", handlers.DeleteTurno)
		})

		r.Route("/seguimientos-turno", func(router chi.Router) {
		router.Use(middleware.AuthMiddleware)

		router.Get("/", handlers.GetSeguimientosTurno)
		})

		r.Route("/notificaciones", func(router chi.Router) {
		router.Use(middleware.AuthMiddleware)

		router.Get("/", handlers.GetNotificaciones)
		})
	})

	log.Println("Servidor escuchando en http://localhost:8080")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}