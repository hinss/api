package base

import (
	"github.com/marmotedu/component-base/pkg/validation/field"
	"github.com/marmotedu/errors"
	"regexp"
)

type SendSmsForm struct {
	Mobile string `form:"mobile" json:"mobile" validate:"required"`
	Type   int    `form:"type" json:"type" validate:"required,oneof=1 2"`
	// type 1-注册 2登录
}

func (s *SendSmsForm) ValidateMobile() field.ErrorList {
	allErrs := field.ErrorList{}
	matched, _ := regexp.MatchString(`^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`, s.Mobile)
	if !matched {
		allErrs = append(allErrs, field.Invalid(field.NewPath("mobile"), errors.New("mobile is invalid!"), ""))
	}

	return allErrs
}

