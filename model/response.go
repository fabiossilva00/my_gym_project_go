package model

type (
	Response[T any] struct {
		Data    T      `json:"data"`
		Message string `json:"message"`
	}

	DataError struct {
		Errors []ErrorResponse `json:"errors"`
	}

	ErrorResponse struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	}
)
