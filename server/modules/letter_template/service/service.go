package service

import (
	"bemunair2026/server/database/entities"
	"bemunair2026/server/modules/letter_template/dto"
	"bemunair2026/server/modules/letter_template/repository"
)

type LetterTemplateService interface {
	Create(req dto.CreateRequest) (*dto.LetterTemplateResponse, error)
	List() ([]dto.LetterTemplateResponse, error)
	Get(id uint64) (*dto.LetterTemplateResponse, error)
	Update(id uint64, req dto.UpdateRequest) (*dto.LetterTemplateResponse, error)
	Delete(id uint64) error
}

type letterTemplateService struct {
	repository repository.LetterTemplateRepository
}

var _ LetterTemplateService = (*letterTemplateService)(nil)

func NewLetterTemplateService(repository repository.LetterTemplateRepository) LetterTemplateService {
	return &letterTemplateService{repository: repository}
}

func (s *letterTemplateService) Create(req dto.CreateRequest) (*dto.LetterTemplateResponse, error) {
	template := &entities.LetterTemplate{
		Name:    req.Name,
		Type:    req.Type,
		Subject: req.Subject,
		Body:    req.Body,
	}
	if err := s.repository.Create(template); err != nil {
		return nil, err
	}
	res := dto.NewLetterTemplateResponse(template)
	return &res, nil
}

func (s *letterTemplateService) List() ([]dto.LetterTemplateResponse, error) {
	rows, err := s.repository.List()
	if err != nil {
		return nil, err
	}
	return dto.NewLetterTemplateResponses(rows), nil
}

func (s *letterTemplateService) Get(id uint64) (*dto.LetterTemplateResponse, error) {
	template, err := s.repository.FindByID(id)
	if err != nil || template == nil {
		return nil, err
	}
	res := dto.NewLetterTemplateResponse(template)
	return &res, nil
}

func (s *letterTemplateService) Update(id uint64, req dto.UpdateRequest) (*dto.LetterTemplateResponse, error) {
	template, err := s.repository.FindByID(id)
	if err != nil || template == nil {
		return nil, err
	}
	template.Name = req.Name
	template.Type = req.Type
	template.Subject = req.Subject
	template.Body = req.Body
	if err := s.repository.Update(template); err != nil {
		return nil, err
	}
	res := dto.NewLetterTemplateResponse(template)
	return &res, nil
}

func (s *letterTemplateService) Delete(id uint64) error {
	return s.repository.Delete(id)
}
