package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// readIDParam получает параметр id из контекста запроса, преобразует его в
// целое число и возвращает его. Если операция не удалась, возвращает 0 и сообщение
// об ошибке.
func (app *application) readIDParam(r *http.Request) (int64, error) {
	param := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid parameter")
	}

	return id, nil
}
