package dto

import (
	"time"

	"bemunair2026/server/database/entities"
)

type CreateRequest struct {
	ServiceType            string
	ContentFormat          *string
	SubmitterPhone         *string
	MinistryID             *uint64
	SongTitle              *string
	SongArtist             *string
	SongStartSeconds       *uint
	SongEndSeconds         *uint
	DocumentationDriveLink *string
	RequiredInformation    *string
	Ministry               string
	SubmissionType         string
	Title                  string
	AddSong                *string
	Caption                string
	AdditionalNotes        *string
	PublishDate            *time.Time
	PublishTime            *string
	DesignDriveLink        *string
	CanvaLink              *string
	ArticleDriveLink       *string
	Deadline               *time.Time
	BriefLink              string
}

type StatusHistoryResponse struct {
	ID         uint64       `json:"id"`
	Actor      *UserSummary `json:"actor,omitempty"`
	FromStatus *string      `json:"from_status,omitempty"`
	ToStatus   string       `json:"to_status"`
	Note       *string      `json:"note,omitempty"`
	CreatedAt  time.Time    `json:"created_at"`
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
	ID                     uint64       `json:"id"`
	RequestCode            *string      `json:"request_code,omitempty"`
	ServiceType            string       `json:"service_type"`
	ContentFormat          *string      `json:"content_format,omitempty"`
	SubmitterName          string       `json:"submitter_name"`
	SubmitterPhone         *string      `json:"submitter_phone,omitempty"`
	MinistryID             *uint64      `json:"ministry_id,omitempty"`
	SubmitterID            uint64       `json:"submitter_id"`
	Submitter              *UserSummary `json:"submitter,omitempty"`
	Ministry               string       `json:"ministry"`
	SubmissionType         string       `json:"submission_type"`
	Title                  string       `json:"title"`
	AddSong                *string      `json:"add_song,omitempty"`
	Caption                string       `json:"caption"`
	AdditionalNotes        *string      `json:"additional_notes,omitempty"`
	PublishDate            *time.Time   `json:"publish_date,omitempty"`
	PublishTime            *string      `json:"publish_time,omitempty"`
	DesignDriveLink        *string      `json:"design_drive_link,omitempty"`
	CanvaLink              *string      `json:"canva_link,omitempty"`
	ArticleDriveLink       *string      `json:"article_drive_link,omitempty"`
	DocumentationDriveLink *string      `json:"documentation_drive_link,omitempty"`
	RequiredInformation    *string      `json:"required_information,omitempty"`
	Deadline               *time.Time   `json:"deadline,omitempty"`
	BriefLink              string       `json:"brief_link"`
	AssignedPJID           *uint64      `json:"assigned_pj_id"`
	AssignedPJ             *UserSummary `json:"assigned_pj,omitempty"`
	Status                 string       `json:"status"`
	Notes                  *string      `json:"notes,omitempty"`
	CreatedAt              time.Time    `json:"created_at"`
	UpdatedAt              time.Time    `json:"updated_at"`
}

func NewContentSubmissionResponse(sub *entities.ContentSubmission) ContentSubmissionResponse {
	if sub == nil {
		return ContentSubmissionResponse{}
	}

	return ContentSubmissionResponse{
		ID:                     sub.ID,
		RequestCode:            sub.RequestCode,
		ServiceType:            sub.ServiceType,
		ContentFormat:          sub.ContentFormat,
		SubmitterName:          sub.SubmitterName,
		SubmitterPhone:         sub.SubmitterPhone,
		MinistryID:             sub.MinistryID,
		SubmitterID:            sub.SubmitterID,
		Submitter:              newUserSummary(sub.Submitter),
		Ministry:               sub.Ministry,
		SubmissionType:         sub.SubmissionType,
		Title:                  sub.Title,
		AddSong:                sub.AddSong,
		Caption:                sub.Caption,
		AdditionalNotes:        sub.AdditionalNotes,
		PublishDate:            sub.PublishDate,
		PublishTime:            sub.PublishTime,
		DesignDriveLink:        sub.DesignDriveLink,
		CanvaLink:              sub.CanvaLink,
		ArticleDriveLink:       sub.ArticleDriveLink,
		DocumentationDriveLink: sub.DocumentationDriveLink,
		RequiredInformation:    sub.RequiredInformation,
		Deadline:               sub.Deadline,
		BriefLink:              sub.BriefLink,
		AssignedPJID:           sub.AssignedPJID,
		AssignedPJ:             newUserSummary(sub.AssignedPJ),
		Status:                 sub.Status,
		Notes:                  sub.Notes,
		CreatedAt:              sub.CreatedAt,
		UpdatedAt:              sub.UpdatedAt,
	}
}

func NewStatusHistoryResponses(rows []entities.ContentSubmissionStatusHistory) []StatusHistoryResponse {
	responses := make([]StatusHistoryResponse, 0, len(rows))
	for i := range rows {
		responses = append(responses, StatusHistoryResponse{ID: rows[i].ID, Actor: newUserSummary(rows[i].Actor), FromStatus: rows[i].FromStatus, ToStatus: rows[i].ToStatus, Note: rows[i].Note, CreatedAt: rows[i].CreatedAt})
	}
	return responses
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
