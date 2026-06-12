package main

import (
	"log"
	"net/http"

	"barberia-cola-virtual/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {

		r.Route("/clientes", func(router chi.Router) {
			router.Get("/{id}", handlers.GetClienteByID)
			router.Put("/{id}", handlers.UpdateCliente)
		})

		r.Route("/servicios", func(router chi.Router) {
			router.Get("/", handlers.GetServicios)
			router.Get("/buscar", handlers.BuscarServicios)
			router.Get("/{id}", handlers.GetServicioByID)
		})

		r.Route("/turnos", func(router chi.Router) {
			router.Post("/", handlers.CreateTurno)
			router.Get("/", handlers.GetTurnos)
			router.Get("/{id}", handlers.GetTurnoByID)
			router.Put("/{id}", handlers.UpdateTurno)
			router.Delete("/{id}", handlers.CancelarTurno)
		})

		r.Route("/metodos-pago", func(router chi.Router) {
			router.Post("/", handlers.CreateMetodoPago)
			router.Get("/", handlers.GetMetodosPago)
			router.Get("/{id}", handlers.GetMetodoPagoByID)
			router.Put("/{id}", handlers.UpdateMetodoPago)
			router.Delete("/{id}", handlers.DeleteMetodoPago)
		})
	})

	log.Println("Servidor escuchando en http://localhost:8080")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}