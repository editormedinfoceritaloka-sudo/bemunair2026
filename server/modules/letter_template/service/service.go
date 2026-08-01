package service

import (
	"bemunair2026/server/database/entities"
	"bemunair2026/server/modules/letter_template/dto"
	"bemunair2026/server/modules/letter_template/repository"
	"errors"
	"strings"
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
	if err := s.validateRequest(req); err != nil {
		return nil, err
	}
	template := &entities.LetterTemplate{
		Name:         req.Name,
		Type:         req.Type,
		Subject:      req.Subject,
		Body:         req.Body,
		MediaAssetID: req.MediaAssetID,
		IsActive:     true,
		DisplayOrder: req.DisplayOrder,
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
	if strings.TrimSpace(req.Name) == "" || strings.TrimSpace(req.Type) == "" {
		return nil, errors.New("nama dan jenis template wajib diisi")
	}
	if req.MediaAssetID != nil {
		if err := s.validateMedia(*req.MediaAssetID); err != nil {
			return nil, err
		}
	}
	template, err := s.repository.FindByID(id)
	if err != nil || template == nil {
		return nil, err
	}
	template.Name = req.Name
	template.Type = req.Type
	template.Subject = req.Subject
	template.Body = req.Body
	if req.MediaAssetID != nil {
		template.MediaAssetID = req.MediaAssetID
	}
	template.IsActive = req.IsActive
	template.DisplayOrder = req.DisplayOrder
	if err := s.repository.Update(template); err != nil {
		return nil, err
	}
	res := dto.NewLetterTemplateResponse(template)
	return &res, nil
}

func (s *letterTemplateService) Delete(id uint64) error {
	return s.repository.Delete(id)
}

func (s *letterTemplateService) validateRequest(req dto.CreateRequest) error {
	if strings.TrimSpace(req.Name) == "" || strings.TrimSpace(req.Type) == "" {
		return errors.New("nama dan jenis template wajib diisi")
	}
	if req.MediaAssetID == nil {
		return errors.New("file PDF template wajib diunggah")
	}
	return s.validateMedia(*req.MediaAssetID)
}

func (s *letterTemplateService) validateMedia(id uint64) error {
	media, err := s.repository.FindMediaAsset(id)
	if err != nil {
		return err
	}
	if media == nil || strings.ToLower(media.MimeType) != "application/pdf" || media.Purpose != "letter_template" {
		return errors.New("media template harus berupa PDF")
	}
	return nil
}
