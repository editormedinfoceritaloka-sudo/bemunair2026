package dto

import (
	"time"

	"bemunair2026/server/database/entities"
)

type CreateRequest struct {
	Ministry       string
	Platform       string
	SubmissionType string
	Caption        string
	Deadline       time.Time
	BriefFile      string
	PosterFile     string
}

type UpdateStatusRequest struct {
	Status string  `json:"status"`
	Notes  *string `json:"notes"`
}

type UserSummary struct {
	ID       uint64  `json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Role     string  `json:"role"`
	Ministry *string `json:"ministry,omitempty"`
	Phone    *string `json:"phone,omitempty"`
}

type ContentSubmissionResponse struct {
	ID             uint64       `json:"id"`
	SubmitterID    uint64       `json:"submitter_id"`
	Submitter      *UserSummary `json:"submitter,omitempty"`
	Ministry       string       `json:"ministry"`
	Platform       string       `json:"platform"`
	SubmissionType string       `json:"submission_type"`
	Caption        string       `json:"caption"`
	Deadline       time.Time    `json:"deadline"`
	BriefFile      string       `json:"brief_file"`
	PosterFile     string       `json:"poster_file"`
	AssignedPJID   *uint64      `json:"assigned_pj_id"`
	AssignedPJ     *UserSummary `json:"assigned_pj,omitempty"`
	Status         string       `json:"status"`
	Notes          *string      `json:"notes,omitempty"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
}

func NewContentSubmissionResponse(sub *entities.ContentSubmission) ContentSubmissionResponse {
	if sub == nil {
		return ContentSubmissionResponse{}
	}

	return ContentSubmissionResponse{
		ID:             sub.ID,
		SubmitterID:    sub.SubmitterID,
		Submitter:      newUserSummary(sub.Submitter),
		Ministry:       sub.Ministry,
		Platform:       sub.Platform,
		SubmissionType: sub.SubmissionType,
		Caption:        sub.Caption,
		Deadline:       sub.Deadline,
		BriefFile:      sub.BriefFile,
		PosterFile:     sub.PosterFile,
		AssignedPJID:   sub.AssignedPJID,
		AssignedPJ:     newUserSummary(sub.AssignedPJ),
		Status:         sub.Status,
		Notes:          sub.Notes,
		CreatedAt:      sub.CreatedAt,
		UpdatedAt:      sub.UpdatedAt,
	}
}

func NewContentSubmissionResponses(submissions []entities.ContentSubmission) []ContentSubmissionResponse {
	responses := make([]ContentSubmissionResponse, 0, len(submissions))
	for i := range submissions {
		responses = append(responses, NewContentSubmissionResponse(&submissions[i]))
	}
	return responses
}

func newUserSummary(user *entities.User) *UserSummary {
	if user == nil {
		return nil
	}

	return &UserSummary{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Role:     user.Role,
		Ministry: user.Ministry,
		Phone:    user.Phone,
	}
}
