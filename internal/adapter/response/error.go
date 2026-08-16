package response

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

const (
	CODE_INVALID_INPUT         = "IVAL"
	CODE_UNAUTHORIZED          = "UNAU"
	CODE_NOT_FOUND             = "NOTF"
	CODE_INTERNAL_SERVER_ERROR = "INSE"

	MSG_INVALID_INPUT         = "invalid input"
	MSG_UNAUTHORIZED          = "unauthorized"
	MSG_NOT_FOUND             = "not found"
	MSG_INTERNAL_SERVER_ERROR = "internal server error"
)

type ErrorResponse struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type ValidateErrorResponse struct {
	Code        string       `json:"code"`
	Description string       `json:"description"`
	Errors      []FieldError `json:"errors,omitempty"`
}

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ValidationError(err error) ValidateErrorResponse {
	res := ValidateErrorResponse{
		Code:        CODE_INVALID_INPUT,
		Description: MSG_INVALID_INPUT,
	}

	var (
		validationErrs validator.ValidationErrors
		typeErr        *json.UnmarshalTypeError
		syntaxErr      *json.SyntaxError
	)

	switch {
	case errors.As(err, &validationErrs):
		res.Errors = make([]FieldError, 0, len(validationErrs))
		for _, fe := range validationErrs {
			res.Errors = append(res.Errors, FieldError{
				Field:   fieldPath(fe),
				Message: message(fe),
			})
		}

	case errors.As(err, &typeErr):
		res.Errors = []FieldError{{
			Field:   typeErr.Field,
			Message: fmt.Sprintf("must be of type %s", typeErr.Type),
		}}

	case errors.As(err, &syntaxErr), errors.Is(err, io.ErrUnexpectedEOF):
		res.Description = "malformed JSON body"

	case errors.Is(err, io.EOF):
		res.Description = "request body is empty"
	}

	return res
}

func fieldPath(fe validator.FieldError) string {
	ns := fe.Namespace()
	if i := strings.IndexByte(ns, '.'); i >= 0 {
		return ns[i+1:]
	}
	return fe.Field()
}

func message(fe validator.FieldError) string {
	field := fieldPath(fe)

	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "oneof":
		return fmt.Sprintf("%s must be one of: %s", field, strings.Join(strings.Fields(fe.Param()), ", "))
	case "min":
		return fmt.Sprintf("%s must be %s", field, quantity(fe, "at least"))
	case "max":
		return fmt.Sprintf("%s must be %s", field, quantity(fe, "at most"))
	case "len":
		return fmt.Sprintf("%s must be %s", field, quantity(fe, "exactly"))
	case "gt":
		return fmt.Sprintf("%s must be greater than %s", field, fe.Param())
	case "gte":
		return fmt.Sprintf("%s must be greater than or equal to %s", field, fe.Param())
	case "lt":
		return fmt.Sprintf("%s must be less than %s", field, fe.Param())
	case "lte":
		return fmt.Sprintf("%s must be less than or equal to %s", field, fe.Param())
	case "eqfield":
		return fmt.Sprintf("%s must match %s", field, fe.Param())
	case "nefield":
		return fmt.Sprintf("%s must not match %s", field, fe.Param())
	case "email":
		return fmt.Sprintf("%s must be a valid email address", field)
	case "url", "uri":
		return fmt.Sprintf("%s must be a valid %s", field, fe.Tag())
	case "uuid", "uuid4":
		return fmt.Sprintf("%s must be a valid UUID", field)
	case "alpha":
		return fmt.Sprintf("%s must contain letters only", field)
	case "alphanum":
		return fmt.Sprintf("%s must contain letters and digits only", field)
	case "numeric":
		return fmt.Sprintf("%s must be numeric", field)
	case "boolean":
		return fmt.Sprintf("%s must be true or false", field)
	case "datetime":
		return fmt.Sprintf("%s must be a date in the format %s", field, fe.Param())
	case "excluded_with", "excluded_without":
		return fmt.Sprintf("%s is not allowed here", field)
	default:
		return fmt.Sprintf("%s failed the %q rule", field, fe.Tag())
	}
}

func quantity(fe validator.FieldError, adverb string) string {
	switch fe.Kind() {
	case reflect.String:
		return fmt.Sprintf("%s %s characters long", adverb, fe.Param())
	case reflect.Slice, reflect.Array, reflect.Map:
		return fmt.Sprintf("%s %s items", adverb, fe.Param())
	default:
		return fmt.Sprintf("%s %s", adverb, fe.Param())
	}
}
