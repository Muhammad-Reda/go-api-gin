package validation

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Struct for errors of validation
type ValidationError struct {
	Tag    string `json:"tag"`
	Field  string `json:"field"`
	Reason string `json:"reason"`
	Type   string `json:"type"`
}

func (ve *ValidationError) Error() string {
	return fmt.Sprintf("Reason: %s, Field: %s, Type: %s, Tag: %s", ve.Reason, ve.Field, ve.Type, ve.Tag)
}

func BodyValidation(c *gin.Context, bodyDeference interface{}) error {
	if err := c.ShouldBindJSON(bodyDeference); err != nil {
		// Check if errors happen in unmarshalling json
		unmarshallTypeError, okUnmarshall := err.(*json.UnmarshalTypeError)
		if okUnmarshall {
			// UnmarshalTypeError true
			return &ValidationError{
				Reason: "Invalid data types",
				Field:  unmarshallTypeError.Field,
				Type:   unmarshallTypeError.Type.Name(),
			}
		}

		// Take errors validation
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// If errors is not from validation
			return &ValidationError{
				Reason: "Invalid format",
				Field:  "Request",
			}
		}

		veer := ValidationError{}
		// Transform error validation to custom error
		for _, err := range errs {
			veer.Field = veer.Field + err.Field() + " "
			veer.Reason = veer.Reason + getErrorMessage(err) + " "
			veer.Tag = veer.Tag + err.Tag() + " "
			veer.Type = veer.Type + err.Type().String() + " "
		}
		return &veer
	}

	return nil
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
	case "oneof":
		return err.Field() + " must be one of: " + err.Param()
	default:
		return "Invalid " + err.Field()
	}
}
