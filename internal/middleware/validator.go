package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateRequest(payload interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBindJSON(payload); err != nil {
			c.Error(&ErrorResponse{
				Code:    400,
				Message: "Invalid request payload",
				Details: err.Error(),
			})
			c.Abort()
			return
		}

		if err := validate.Struct(payload); err != nil {
			if validationErrors, ok := err.(validator.ValidationErrors); ok {
				c.Error(&ErrorResponse{
					Code:    400,
					Message: "Validation failed",
					Details: formatValidationErrors(validationErrors),
				})
				c.Abort()
				return
			}
		}

		c.Set("payload", payload)
		c.Next()
	}
}
