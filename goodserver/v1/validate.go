package v1

import (
	"github.com/marmotedu/component-base/pkg/validation"
	"github.com/marmotedu/component-base/pkg/validation/field"
)

// Validate validates that a user object is valid.
func Validate(validated interface{}) field.ErrorList {
	val := validation.NewValidator(validated)
	allErrs := val.Validate()

	return allErrs
}
