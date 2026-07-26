package service

import (
	"errors"

	"bemunair2026/server/database/entities"
	"bemunair2026/server/modules/user/dto"
	"bemunair2026/server/modules/user/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	List() ([]dto.UserResponse, error)
	Get(id uint64) (*dto.UserResponse, error)
	Create(req dto.UserCreateRequest) (*dto.UserResponse, error)
	Update(id uint64, req dto.UserUpdateRequest) (*dto.UserResponse, error)
	Delete(id uint64) error
}

type userService struct {
	userRepository repository.UserRepository
}

var _ UserService = (*userService)(nil)

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (s *userService) List() ([]dto.UserResponse, error) {
	users, err := s.userRepository.List()
	if err != nil {
		return nil, err
	}
	return dto.NewUserResponses(users), nil
}

func (s *userService) Get(id uint64) (*dto.UserResponse, error) {
	user, err := s.userRepository.FindByID(id)
	if err != nil || user == nil {
		return nil, err
	}
	res := dto.NewUserResponse(user)
	return &res, nil
}

func (s *userService) Create(req dto.UserCreateRequest) (*dto.UserResponse, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	ministry, err := s.resolveMinistry(req.MinistryID)
	if err != nil {
		return nil, err
	}

	user := &entities.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: string(hash),
		Role:         req.Role,
		MinistryID:   req.MinistryID,
		Ministry:     ministry,
		Phone:        req.Phone,
	}
	if err := s.userRepository.Create(user); err != nil {
		return nil, err
	}
	res := dto.NewUserResponse(user)
	return &res, nil
}

func (s *userService) Update(id uint64, req dto.UserUpdateRequest) (*dto.UserResponse, error) {
	user, err := s.userRepository.FindByID(id)
	if err != nil || user == nil {
		return nil, err
	}
	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Role != "" {
		user.Role = req.Role
	}
	ministry, err := s.resolveMinistry(req.MinistryID)
	if err != nil {
		return nil, err
	}
	user.MinistryID = req.MinistryID
	user.Ministry = ministry
	user.MinistryRef = nil
	user.Phone = req.Phone

	if err := s.userRepository.Update(user); err != nil {
		return nil, err
	}
	res := dto.NewUserResponse(user)
	return &res, nil
}

func (s *userService) Delete(id uint64) error {
	return s.userRepository.Delete(id)
}

func (s *userService) resolveMinistry(id *uint64) (*string, error) {
	if id == nil {
		return nil, nil
	}
	ministry, err := s.userRepository.FindMinistryByID(*id)
	if err != nil {
		return nil, err
	}
	if ministry == nil {
		return nil, errors.New("kementerian tidak ditemukan")
	}
	if !ministry.IsActive {
		return nil, errors.New("kementerian tidak aktif")
	}
	return &ministry.Code, nil
}
