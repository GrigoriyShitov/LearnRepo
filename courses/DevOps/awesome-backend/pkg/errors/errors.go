package errors

import (
	"net/http"

	"git.itmo.su/go-modules/xerr/v4"
)

var (
	ErrValidationFailed = xerr.New("validation_error", http.StatusUnprocessableEntity).
				Russian("ошибка валидации входных данных").
				English("input data validation error")

	ErrNotFound = xerr.New("not_found", http.StatusNotFound).
			Russian("ресурс не найден").
			English("resource not found")
)
