package model

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func NewCustomValidator(validator *validator.Validate) *CustomValidator {
	return &CustomValidator{validator}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {

		var responseError Response[DataError]
		var errors []ErrorResponse

		for _, err := range err.(validator.ValidationErrors) {

			message := fmt.Sprintf("Problema na validação, %s", err.ActualTag())

			err := ErrorResponse{
				Field:   strings.ToLower(err.Field()),
				Message: message,
			}

			errors = append(errors, err)
		}

		responseError.Message = "Ocorreu um ou mais erros"
		responseError.Data = DataError{errors}

		return echo.NewHTTPError(http.StatusBadRequest, responseError)
	}

	return nil
}
