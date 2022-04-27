package helper

import (
	"github.com/afiefafian/go-pos/src/model"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(m interface{}) error {
	return validate.Struct(m)
}

func FormatValidationError(err validator.ValidationErrors) []*model.ValidationError {
	var errors = make([]*model.ValidationError, 0)
	if err != nil {
		for _, err := range err {
			var element model.ValidationError

			element.FailedField = err.StructField()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
