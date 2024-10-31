package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// Добавляем обработчик createBookHandler для эндпоинта `POST /v1/books`.
// Пока мы просто возвращаем текст.
func (app *application) createBookHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create a new book")
}

// Добавляем обработчик showBookHandler для конечной точки `GET /v1/books/:id`.
// На данный момент мы получаем параметр id из URL-адреса и включаем его в ответ.
func (app *application) showBookHandler(w http.ResponseWriter, r *http.Request) {
	// Мы можем использовать функцию chi.URLParam(), чтобы получить параметр id из URL.
  // Первый параметр в функции — это http.Request.
	param := chi.URLParam(r, "id")

	// В нашем проекте все книги будут иметь уникальный положительный идентификатор,
	// но возвращаемое значение метода URLParam() всегда является строкой,
	// поэтому мы пытаемся преобразовать её в целое число. Если не удаётся
	// её преобразовать или id меньше 1, то идентификатор является недействительным, и мы
	// вызываем http.NotFound() для возврата ошибки 404.
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// Если идентификатор валидный, мы возвращаем его в ответе.
	fmt.Fprintf(w, "show the details of book %d\n", id)
}
