package user

import (
	"github.com/marmotedu/component-base/pkg/validation/field"
	"github.com/marmotedu/errors"
	"regexp"
)

// RegisterForm use when
type RegisterForm struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required"`
	PassWord string `form:"password" json:"password" binding:"required,min=3,max=10"`
	Code     string `form:"code" json:"code" binding:"required"`
}

func (r *RegisterForm) ValidateMobile() field.ErrorList {
	allErrs := field.ErrorList{}
	matched, _ := regexp.MatchString(`^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`, r.Mobile)
	if !matched {
		allErrs = append(allErrs, field.Invalid(field.NewPath("mobile"), errors.New("mobile is invalid!"), ""))
	}

	return allErrs
}
