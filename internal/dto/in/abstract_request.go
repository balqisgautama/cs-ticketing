package dtoin

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
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
				if fieldError.Field() == "Title" {
					return errors.New("ticket title must be at least 10 characters")
				} else if fieldError.Field() == "Msg" {
					return errors.New("ticket message must be at least 100 characters")
				}
			case "max":
				if fieldError.Field() == "Title" {
					return errors.New("ticket title must be between 10 and 100 characters")
				}
			case "oneof":
				if fieldError.Field() == "FilterType" {
					return errors.New("filter_type must be one of 'before', 'after', or 'between'")
				} else if fieldError.Field() == "FilterName" {
					return errors.New("filter_name must be 'created_at'")
				} else if fieldError.Field() == "SortDir" {
					return errors.New("sort_dir must be 'asc' or 'desc'")
				}
			}
		}
	}
	return err // Return the original error if no specific translation is found
}
