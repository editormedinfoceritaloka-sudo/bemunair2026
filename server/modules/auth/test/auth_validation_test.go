package test

import (
	"testing"

	"bemunair2026/server/modules/auth/dto"
	"bemunair2026/server/modules/auth/validation"
	"bemunair2026/server/pkg/constants"
)

func TestAuthValidationRegisterRequest(t *testing.T) {
	v := validation.NewAuthValidation()

	if err := v.ValidateRegisterRequest(dto.RegisterRequest{
		Name:     "Admin",
		Email:    "admin@test.com",
		Password: "password",
		Role:     constants.RoleAdmin,
	}); err != nil {
		t.Fatalf("expected valid register request, got %v", err)
	}

	if err := v.ValidateRegisterRequest(dto.RegisterRequest{
		Name:     "Admin",
		Email:    "admin@test.com",
		Password: "short",
		Role:     constants.RoleAdmin,
	}); err == nil {
		t.Fatal("expected short password to be invalid")
	}

	if err := v.ValidateRegisterRequest(dto.RegisterRequest{
		Name:     "Admin",
		Email:    "not-email",
		Password: "password",
		Role:     constants.RoleAdmin,
	}); err == nil {
		t.Fatal("expected invalid email to be rejected")
	}
}

func TestAuthValidationLoginRequest(t *testing.T) {
	v := validation.NewAuthValidation()

	if err := v.ValidateLoginRequest(dto.LoginRequest{Email: "admin@test.com", Password: "password"}); err != nil {
		t.Fatalf("expected valid login request, got %v", err)
	}

	if err := v.ValidateLoginRequest(dto.LoginRequest{Email: "", Password: "password"}); err == nil {
		t.Fatal("expected missing email to be invalid")
	}
}
