package dto

import (
	"time"

	"bemunair2026/server/database/entities"
)

type CreateRequest struct {
	Name         string  `json:"name"`
	Type         string  `json:"type"`
	Subject      string  `json:"subject"`
	Body         string  `json:"body"`
	MediaAssetID *uint64 `json:"media_asset_id"`
	IsActive     bool    `json:"is_active"`
	DisplayOrder uint    `json:"display_order"`
}

type UpdateRequest struct {
	Name         string  `json:"name"`
	Type         string  `json:"type"`
	Subject      string  `json:"subject"`
	Body         string  `json:"body"`
	MediaAssetID *uint64 `json:"media_asset_id"`
	IsActive     bool    `json:"is_active"`
	DisplayOrder uint    `json:"display_order"`
}

type TemplateFileResponse struct {
	ID        uint64 `json:"id"`
	URL       string `json:"url"`
	Name      string `json:"name"`
	MimeType  string `json:"mime_type"`
	SizeBytes uint64 `json:"size_bytes"`
}

type LetterTemplateResponse struct {
	ID           uint64                `json:"id"`
	Name         string                `json:"name"`
	Type         string                `json:"type"`
	Subject      string                `json:"subject"`
	Body         string                `json:"body"`
	MediaAssetID *uint64               `json:"media_asset_id,omitempty"`
	File         *TemplateFileResponse `json:"file,omitempty"`
	DownloadURL  string                `json:"download_url,omitempty"`
	IsActive     bool                  `json:"is_active"`
	DisplayOrder uint                  `json:"display_order"`
	CreatedAt    time.Time             `json:"created_at"`
	UpdatedAt    time.Time             `json:"updated_at"`
}

func NewLetterTemplateResponse(template *entities.LetterTemplate) LetterTemplateResponse {
	if template == nil {
		return LetterTemplateResponse{}
	}

	return LetterTemplateResponse{
		ID:           template.ID,
		Name:         template.Name,
		Type:         template.Type,
		Subject:      template.Subject,
		Body:         template.Body,
		MediaAssetID: template.MediaAssetID,
		IsActive:     template.IsActive,
		DisplayOrder: template.DisplayOrder,
		CreatedAt:    template.CreatedAt,
		UpdatedAt:    template.UpdatedAt,
		File: func() *TemplateFileResponse {
			if template.MediaAsset == nil {
				return nil
			}
			return &TemplateFileResponse{ID: template.MediaAsset.ID, URL: template.MediaAsset.URL, Name: template.MediaAsset.Name, MimeType: template.MediaAsset.MimeType, SizeBytes: template.MediaAsset.SizeBytes}
		}(),
		DownloadURL: func() string {
			if template.MediaAsset == nil {
				return ""
			}
			return template.MediaAsset.URL
		}(),
	}
}

func NewLetterTemplateResponses(templates []entities.LetterTemplate) []LetterTemplateResponse {
	responses := make([]LetterTemplateResponse, 0, len(templates))
	for i := range templates {
		responses = append(responses, NewLetterTemplateResponse(&templates[i]))
	}
	return responses
}
