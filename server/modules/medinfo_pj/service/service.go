package service

import (
	"bemunair2026/server/database/entities"
	"bemunair2026/server/modules/medinfo_pj/dto"
	"bemunair2026/server/modules/medinfo_pj/repository"
)

type MedinfoPJService interface {
	List() ([]dto.QueueResponse, error)
	Create(req dto.CreateRequest) (*dto.QueueResponse, error)
	Reorder(req dto.ReorderRequest) error
	Delete(id uint64) error
}

type medinfoPJService struct {
	repository repository.MedinfoPJRepository
}

var _ MedinfoPJService = (*medinfoPJService)(nil)

func NewMedinfoPJService(repository repository.MedinfoPJRepository) MedinfoPJService {
	return &medinfoPJService{repository: repository}
}

func (s *medinfoPJService) List() ([]dto.QueueResponse, error) {
	rows, err := s.repository.List()
	if err != nil {
		return nil, err
	}
	return dto.NewQueueResponses(rows), nil
}

func (s *medinfoPJService) Create(req dto.CreateRequest) (*dto.QueueResponse, error) {
	row := &entities.MedinfoPJQueue{UserID: req.UserID, Position: req.Position}
	if row.Position == 0 {
		row.Position = 1
	}
	if err := s.repository.Create(row); err != nil {
		return nil, err
	}
	res := dto.NewQueueResponse(row)
	return &res, nil
}

func (s *medinfoPJService) Reorder(req dto.ReorderRequest) error {
	return s.repository.Reorder(req.IDs)
}

func (s *medinfoPJService) Delete(id uint64) error {
	return s.repository.Delete(id)
}
