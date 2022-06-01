package user

import (
	"github.com/marmotedu/component-base/pkg/validation/field"
	"github.com/marmotedu/errors"
	"regexp"
)

// RegisterForm use when register a user.
type RegisterForm struct {
	Mobile   string `form:"mobile" json:"mobile" validate:"required"`
	PassWord string `form:"password" json:"password" validate:"required,min=3,max=10"`
	Code     string `form:"code" json:"code" validate:"required"`
	// 兼容 iam loginInfo basic认证登录字段
	Username string `form:"username" json:"username" validate:"required"`

}

func (r *RegisterForm) ValidateMobile() field.ErrorList {
	allErrs := field.ErrorList{}
	matched, _ := regexp.MatchString(`^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`, r.Mobile)
	if !matched {
		allErrs = append(allErrs, field.Invalid(field.NewPath("mobile"), errors.New("mobile is invalid!"), ""))
	}

	return allErrs
}
