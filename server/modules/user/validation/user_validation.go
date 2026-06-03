package validation

import (
	"bemunair2026/server/modules/user/dto"
	"github.com/go-playground/validator/v10"
)

type UserValidation struct {
	validate *validator.Validate
}

func NewUserValidation() *UserValidation {
	validate := validator.New()

	_ = validate.RegisterValidation("password", validatePassword)

	return &UserValidation{validate: validate}
}

func (v *UserValidation) ValidateCreateRequest(req dto.UserCreateRequest) error {
	return v.validate.Struct(req)
}

func (v *UserValidation) ValidateUpdateRequest(req dto.UserUpdateRequest) error {
	return v.validate.Struct(req)
}

func validatePassword(fl validator.FieldLevel) bool {
	return len(fl.Field().String()) >= 8
}
