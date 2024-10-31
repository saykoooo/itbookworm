package main

import (
	"github.com/go-chi/chi/v5"
)

func (app *application) routes() *chi.Mux {
	// Инициализируем новый маршрутизатор chi.
	router := chi.NewRouter()

	// Регистрируем шаблоны URL и методы обработчиков.
	router.Get("/v1/healthcheck", app.healhcheckHandler)
	router.Post("/v1/books", app.createBookHandler)
	router.Get("/v1/books/{id}", app.showBookHandler)

	// Возвращаем экземпляр chi.Mux
	return router
}
