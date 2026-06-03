package validation

import (
	"bemunair2026/server/modules/auth/dto"
	"github.com/go-playground/validator/v10"
)

type AuthValidation struct {
	validate *validator.Validate
}

func NewAuthValidation() *AuthValidation {
	validate := validator.New()

	_ = validate.RegisterValidation("password", validatePassword)

	return &AuthValidation{validate: validate}
}

func (v *AuthValidation) ValidateRegisterRequest(req dto.RegisterRequest) error {
	return v.validate.Struct(req)
}

func (v *AuthValidation) ValidateLoginRequest(req dto.LoginRequest) error {
	return v.validate.Struct(req)
}

func validatePassword(fl validator.FieldLevel) bool {
	return len(fl.Field().String()) >= 8
}
