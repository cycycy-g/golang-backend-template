package middleware

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Only handle errors if there are any
		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last()

		var response ErrorResponse

		switch e := err.Err.(type) {
		case validator.ValidationErrors:
			response = ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: "Validation failed",
				Details: formatValidationErrors(e),
			}
		case *ErrorResponse:
			response = *e
		default:
			response = ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "Internal server error",
			}
		}

		c.JSON(response.Code, response)
	}
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("code: %d, message: %s", e.Code, e.Message)
}

func formatValidationErrors(errs validator.ValidationErrors) map[string]string {
	errors := make(map[string]string)
	for _, err := range errs {
		errors[err.Field()] = formatValidationError(err)
	}
	return errors
}

func formatValidationError(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	// Add more cases as needed
	default:
		return err.Error()
	}
}
