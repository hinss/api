package base

type SendSmsForm struct {
	Mobile string `form:"mobile" json:"mobile" validate:"required,mobile"`
	Type   int    `form:"type" json:"type" validate:"required,oneof=1 2"`
	// type 1-注册 2登录
}
