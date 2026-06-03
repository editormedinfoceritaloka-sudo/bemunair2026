package service

import (
	"errors"
	"time"

	"bemunair2026/server/database/entities"
	"bemunair2026/server/middlewares"
	"bemunair2026/server/modules/auth/dto"
	"bemunair2026/server/modules/auth/repository"
	"bemunair2026/server/pkg/constants"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo      *repository.Repository
	jwtSecret string
}

func NewService(repo *repository.Repository, jwtSecret string) *Service {
	return &Service{repo: repo, jwtSecret: jwtSecret}
}

func (s *Service) Register(req dto.RegisterRequest) (*dto.UserResponse, error) {
	if req.Role != constants.RoleAdmin && req.Role != constants.RoleMentri {
		return nil, errors.New("invalid role")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &entities.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: string(hash),
		Role:         req.Role,
		Ministry:     req.Ministry,
		Phone:        req.Phone,
	}
	if err := s.repo.Create(user); err != nil {
		return nil, err
	}
	res := dto.NewUserResponse(user)
	return &res, nil
}

func (s *Service) Login(req dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil || user == nil {
		return nil, errors.New("invalid credentials")
	}
	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)) != nil {
		return nil, errors.New("invalid credentials")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, middlewares.Claims{
		UserID: user.ID, Role: user.Role, Ministry: user.Ministry,
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour))},
	})
	signed, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return nil, err
	}
	return &dto.LoginResponse{Token: signed, User: dto.NewUserResponse(user)}, nil
}

func (s *Service) Me(userID uint64) (*dto.UserResponse, error) {
	user, err := s.repo.FindByID(userID)
	if err != nil || user == nil {
		return nil, err
	}
	res := dto.NewUserResponse(user)
	return &res, nil
}
