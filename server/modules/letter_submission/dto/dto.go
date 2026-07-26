package dto

import (
	"time"

	"bemunair2026/server/database/entities"
)

type CreateRequest struct {
	Ministry   string `json:"ministry"`
	LetterType string `json:"letter_type"`
	Subject    string `json:"subject"`
	Body       string `json:"body"`
	Deadline   string `json:"deadline"`
	DeadlineAt time.Time
}

type StatusHistoryResponse struct {
	EventType  string       `json:"event_type"`
	ID         uint64       `json:"id"`
	Actor      *UserSummary `json:"actor,omitempty"`
	FromStatus *string      `json:"from_status,omitempty"`
	ToStatus   string       `json:"to_status"`
	Note       *string      `json:"note,omitempty"`
	CreatedAt  time.Time    `json:"created_at"`
	FromPJ     *UserSummary `json:"from_pj,omitempty"`
	ToPJ       *UserSummary `json:"to_pj,omitempty"`
}

type UpdateStatusRequest struct {
	Status string  `json:"status"`
	Notes  *string `json:"notes"`
}

type AssignPJRequest struct {
	AssignedPJID uint64 `json:"assigned_pj_id"`
}

type UserSummary struct {
	ID       uint64  `json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Role     string  `json:"role"`
	Ministry *string `json:"ministry,omitempty"`
	Phone    *string `json:"phone,omitempty"`
}

type LetterSubmissionResponse struct {
	ID             uint64       `json:"id"`
	RequestCode    *string      `json:"request_code,omitempty"`
	SubmitterName  string       `json:"submitter_name"`
	SubmitterPhone *string      `json:"submitter_phone,omitempty"`
	MinistryID     *uint64      `json:"ministry_id,omitempty"`
	SubmitterID    uint64       `json:"submitter_id"`
	Submitter      *UserSummary `json:"submitter,omitempty"`
	Ministry       string       `json:"ministry"`
	LetterType     string       `json:"letter_type"`
	Subject        string       `json:"subject"`
	Body           string       `json:"body"`
	Deadline       time.Time    `json:"deadline"`
	AssignedPJID   *uint64      `json:"assigned_pj_id"`
	AssignedPJ     *UserSummary `json:"assigned_pj,omitempty"`
	Status         string       `json:"status"`
	Notes          *string      `json:"notes,omitempty"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
}

func NewLetterSubmissionResponse(submission *entities.LetterSubmission) LetterSubmissionResponse {
	if submission == nil {
		return LetterSubmissionResponse{}
	}

	return LetterSubmissionResponse{
		ID:             submission.ID,
		RequestCode:    submission.RequestCode,
		SubmitterName:  submission.SubmitterName,
		SubmitterPhone: submission.SubmitterPhone,
		MinistryID:     submission.MinistryID,
		SubmitterID:    submission.SubmitterID,
		Submitter:      newUserSummary(submission.Submitter),
		Ministry:       submission.Ministry,
		LetterType:     submission.LetterType,
		Subject:        submission.Subject,
		Body:           submission.Body,
		Deadline:       submission.Deadline,
		AssignedPJID:   submission.AssignedPJID,
		AssignedPJ:     newUserSummary(submission.AssignedPJ),
		Status:         submission.Status,
		Notes:          submission.Notes,
		CreatedAt:      submission.CreatedAt,
		UpdatedAt:      submission.UpdatedAt,
	}
}

func NewStatusHistoryResponses(rows []entities.LetterSubmissionStatusHistory) []StatusHistoryResponse {
	responses := make([]StatusHistoryResponse, 0, len(rows))
	for i := range rows {
		responses = append(responses, StatusHistoryResponse{EventType: "STATUS_CHANGED", ID: rows[i].ID, Actor: newUserSummary(rows[i].Actor), FromStatus: rows[i].FromStatus, ToStatus: rows[i].ToStatus, Note: rows[i].Note, CreatedAt: rows[i].CreatedAt})
	}
	return responses
}

func NewAssignmentHistoryResponses(rows []entities.LetterSubmissionAssignmentHistory) []StatusHistoryResponse {
	responses := make([]StatusHistoryResponse, 0, len(rows))
	for i := range rows {
		eventType := "PJ_ASSIGNED"
		if rows[i].FromPJID != nil {
			eventType = "PJ_REASSIGNED"
		}
		responses = append(responses, StatusHistoryResponse{EventType: eventType, ID: rows[i].ID, Actor: newUserSummary(rows[i].Actor), Note: rows[i].Note, CreatedAt: rows[i].CreatedAt, FromPJ: newUserSummary(rows[i].FromPJ), ToPJ: newUserSummary(rows[i].ToPJ)})
	}
	return responses
}

func NewLetterSubmissionResponses(submissions []entities.LetterSubmission) []LetterSubmissionResponse {
	responses := make([]LetterSubmissionResponse, 0, len(submissions))
	for i := range submissions {
		responses = append(responses, NewLetterSubmissionResponse(&submissions[i]))
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
