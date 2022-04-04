package v1

import (
	"github.com/marmotedu/component-base/pkg/validation"
	"github.com/marmotedu/component-base/pkg/validation/field"
)

// Validate validates that a user object is valid.
func (c *Category) Validate() field.ErrorList {
	val := validation.NewValidator(c)
	allErrs := val.Validate()

	return allErrs
}
