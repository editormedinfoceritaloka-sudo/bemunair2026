package auth

import (
	"errors"
	"time"

	"bemunair2026/server/database/entities"
	"bemunair2026/server/middlewares"
	"bemunair2026/server/modules/user"
	"bemunair2026/server/pkg/constants"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	users     *user.Repository
	jwtSecret string
}

func NewService(users *user.Repository, jwtSecret string) *Service {
	return &Service{users: users, jwtSecret: jwtSecret}
}

func (s *Service) Register(name, email, password, role string, ministry, phone *string) (*entities.User, error) {
	if role != constants.RoleAdmin && role != constants.RoleMentri {
		return nil, errors.New("invalid role")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	u := &entities.User{Name: name, Email: email, PasswordHash: string(hash), Role: role, Ministry: ministry, Phone: phone}
	return u, s.users.Create(u)
}

func (s *Service) Login(email, password string) (string, *entities.User, error) {
	u, err := s.users.FindByEmail(email)
	if err != nil || u == nil {
		return "", nil, errors.New("invalid credentials")
	}
	if bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)) != nil {
		return "", nil, errors.New("invalid credentials")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, middlewares.Claims{
		UserID: u.ID, Role: u.Role, Ministry: u.Ministry,
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour))},
	})
	signed, err := token.SignedString([]byte(s.jwtSecret))
	return signed, u, err
}
