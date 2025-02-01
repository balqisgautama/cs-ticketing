package dtoin

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator"
)

func Validate(input interface{}) error {
	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		return translateValidationErrors(err)
	}
	return nil
}

// translateValidationErrors translates validation errors to user-friendly messages
func translateValidationErrors(err error) error {
	var validationErrors validator.ValidationErrors
	if ok := errors.As(err, &validationErrors); ok {
		for _, fieldError := range validationErrors {
			switch fieldError.Tag() {
			case "required":
				return fmt.Errorf("%s is required", fieldError.Field())
			case "min":
				if fieldError.Field() == "TicketTitle" {
					return errors.New("Ticket title must be at least 10 characters")
				} else if fieldError.Field() == "TicketMsg" {
					return errors.New("Ticket message must be at least 100 characters")
				}
			case "max":
				if fieldError.Field() == "TicketTitle" {
					return errors.New("Ticket title must be between 10 and 100 characters")
				}
			}
		}
	}
	return err // Return the original error if no specific translation is found
}
