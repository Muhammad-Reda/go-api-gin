package validation

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Struct for errors of validation
type ValidationError struct {
	Tag    string `json:"tag"`
	Field  string `json:"field"`
	Reason string `json:"reason"`
}

func UserValidation(c *gin.Context, user interface{}) []ValidationError {
	var veer []ValidationError

	if err := c.ShouldBindJSON(user); err != nil {
		// Check if errors happen in unmarshalling json
		unmarshallTypeError, okUnmarshall := err.(*json.UnmarshalTypeError)
		if okUnmarshall {
			// UnmarshalTypeError true
			veer = append(veer, ValidationError{
				Reason: "Invalid data types",
				Field:  unmarshallTypeError.Field,
			})
		}

		// Take errors validation
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// If errors is not from validation
			veer = append(veer, ValidationError{
				Reason: "Invalid format",
				Field:  "Request",
			})
			return veer

		}

		// Transform error validation to custom error
		for _, err := range errs {
			veer = append(veer, ValidationError{
				Tag:    err.Tag(),
				Field:  err.Field(),
				Reason: getErrorMessage(err),
			})
		}
	}

	return veer
}

func getErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return err.Field() + " can not blank"
	case "email":
		return err.Field() + " is not a valid email address"
	case "gte":
		return err.Field() + " must be greater than or equal to " + err.Param()
	case "lte":
		return err.Field() + " must be less than or equal to " + err.Param()
	default:
		return "Invalid " + err.Field()
	}
}
