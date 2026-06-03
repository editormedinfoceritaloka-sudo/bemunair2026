package dto

import (
	"time"

	"bemunair2026/server/database/entities"
)

type CreateRequest struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type UpdateRequest struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type LetterTemplateResponse struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	Subject   string    `json:"subject"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewLetterTemplateResponse(template *entities.LetterTemplate) LetterTemplateResponse {
	if template == nil {
		return LetterTemplateResponse{}
	}

	return LetterTemplateResponse{
		ID:        template.ID,
		Name:      template.Name,
		Type:      template.Type,
		Subject:   template.Subject,
		Body:      template.Body,
		CreatedAt: template.CreatedAt,
		UpdatedAt: template.UpdatedAt,
	}
}

func NewLetterTemplateResponses(templates []entities.LetterTemplate) []LetterTemplateResponse {
	responses := make([]LetterTemplateResponse, 0, len(templates))
	for i := range templates {
		responses = append(responses, NewLetterTemplateResponse(&templates[i]))
	}
	return responses
}
