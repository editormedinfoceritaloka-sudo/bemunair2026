package test

import (
	"testing"

	"bemunair2026/server/modules/user/dto"
	"bemunair2026/server/modules/user/validation"
	"bemunair2026/server/pkg/constants"
)

func TestUserValidationCreateRequest(t *testing.T) {
	v := validation.NewUserValidation()

	if err := v.ValidateCreateRequest(dto.UserCreateRequest{
		Name:     "Admin",
		Email:    "admin@test.com",
		Password: "password",
		Role:     constants.RoleAdmin,
	}); err != nil {
		t.Fatalf("expected valid create request, got %v", err)
	}

	if err := v.ValidateCreateRequest(dto.UserCreateRequest{
		Name:     "Admin",
		Email:    "admin@test.com",
		Password: "short",
		Role:     constants.RoleAdmin,
	}); err == nil {
		t.Fatal("expected short password to be invalid")
	}
}

func TestUserValidationUpdateRequest(t *testing.T) {
	v := validation.NewUserValidation()

	if err := v.ValidateUpdateRequest(dto.UserUpdateRequest{Email: "admin@test.com"}); err != nil {
		t.Fatalf("expected valid update request, got %v", err)
	}

	if err := v.ValidateUpdateRequest(dto.UserUpdateRequest{Email: "not-email"}); err == nil {
		t.Fatal("expected invalid email to be rejected")
	}
}
