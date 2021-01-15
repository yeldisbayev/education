package entity

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator"
)

// QueryError struct
type QueryError struct {
	Code        string       `json:"code"`
	Message     string       `json:"message"`
	Err         error        `json:"-"`
	FieldErrors []FieldError `json:"fieldErrors,omitempty"`
}

// FieldError struct
type FieldError struct {
	Field string `json:"field"`
	Rule  string `json:"rule"`
}

// NewQueryError constructor
func NewQueryError(code, message string, err error) QueryError {
	return QueryError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// NewQueryFormError constructor
func NewQueryFormError(err error) QueryError {
	validationErrors := new(validator.ValidationErrors)

	queryError := QueryError{
		Code:        "QUERY_FORM_ERROR",
		Message:     "Ошибка отправляемых данных",
		FieldErrors: []FieldError{},
	}

	if errors.As(err, validationErrors) {
		for _, fieldError := range err.(validator.ValidationErrors) {
			queryError.FieldErrors = append(
				queryError.FieldErrors,
				FieldError{
					Field: fieldError.Field(),
					Rule:  fieldError.Tag(),
				},
			)
		}
	}

	return queryError
}

func (e *QueryError) Error() string {
	return fmt.Sprintf(
		"Code: %s; Message: %s; DevMessage: %w",
		e.Code,
		e.Message,
		e.Err,
	)
}
