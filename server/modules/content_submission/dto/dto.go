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
	MediaFileID            *string
	MediaFileName          *string
	MediaFileMimeType      *string
	MediaFileSize          uint64
	BriefFileID            *string
	BriefFileName          *string
	BriefFileMimeType      *string
	BriefFileSize          uint64
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

type AttachmentResponse struct {
	ID             uint64 `json:"id"`
	ImageKitFileID string `json:"imagekit_file_id"`
	Purpose        string `json:"purpose"`
	Name           string `json:"name"`
	URL            string `json:"url"`
	MimeType       string `json:"mime_type"`
	SizeBytes      uint64 `json:"size_bytes"`
}

type ContentSubmissionResponse struct {
	ID                     uint64               `json:"id"`
	RequestCode            *string              `json:"request_code,omitempty"`
	ServiceType            string               `json:"service_type"`
	ContentFormat          *string              `json:"content_format,omitempty"`
	SubmitterName          string               `json:"submitter_name"`
	SubmitterPhone         *string              `json:"submitter_phone,omitempty"`
	MinistryID             *uint64              `json:"ministry_id,omitempty"`
	SubmitterID            uint64               `json:"submitter_id"`
	Submitter              *UserSummary         `json:"submitter,omitempty"`
	Ministry               string               `json:"ministry"`
	SubmissionType         string               `json:"submission_type"`
	Title                  string               `json:"title"`
	AddSong                *string              `json:"add_song,omitempty"`
	Caption                string               `json:"caption"`
	AdditionalNotes        *string              `json:"additional_notes,omitempty"`
	PublishDate            *time.Time           `json:"publish_date,omitempty"`
	PublishTime            *string              `json:"publish_time,omitempty"`
	DesignDriveLink        *string              `json:"design_drive_link,omitempty"`
	CanvaLink              *string              `json:"canva_link,omitempty"`
	ArticleDriveLink       *string              `json:"article_drive_link,omitempty"`
	DocumentationDriveLink *string              `json:"documentation_drive_link,omitempty"`
	RequiredInformation    *string              `json:"required_information,omitempty"`
	Deadline               *time.Time           `json:"deadline,omitempty"`
	BriefLink              string               `json:"brief_link"`
	Attachments            []AttachmentResponse `json:"attachments"`
	AssignedPJID           *uint64              `json:"assigned_pj_id"`
	AssignedPJ             *UserSummary         `json:"assigned_pj,omitempty"`
	Status                 string               `json:"status"`
	Notes                  *string              `json:"notes,omitempty"`
	CreatedAt              time.Time            `json:"created_at"`
	UpdatedAt              time.Time            `json:"updated_at"`
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
		Attachments:            newAttachmentResponses(sub.Attachments),
		AssignedPJID:           sub.AssignedPJID,
		AssignedPJ:             newUserSummary(sub.AssignedPJ),
		Status:                 sub.Status,
		Notes:                  sub.Notes,
		CreatedAt:              sub.CreatedAt,
		UpdatedAt:              sub.UpdatedAt,
	}
}

func newAttachmentResponses(rows []entities.ContentSubmissionAttachment) []AttachmentResponse {
	responses := make([]AttachmentResponse, 0, len(rows))
	for i := range rows {
		responses = append(responses, AttachmentResponse{ID: rows[i].ID, ImageKitFileID: rows[i].ImageKitFileID, Purpose: rows[i].Purpose, Name: rows[i].Name, URL: rows[i].URL, MimeType: rows[i].MimeType, SizeBytes: rows[i].SizeBytes})
	}
	return responses
}

func NewStatusHistoryResponses(rows []entities.ContentSubmissionStatusHistory) []StatusHistoryResponse {
	responses := make([]StatusHistoryResponse, 0, len(rows))
	for i := range rows {
		responses = append(responses, StatusHistoryResponse{EventType: "STATUS_CHANGED", ID: rows[i].ID, Actor: newUserSummary(rows[i].Actor), FromStatus: rows[i].FromStatus, ToStatus: rows[i].ToStatus, Note: rows[i].Note, CreatedAt: rows[i].CreatedAt})
	}
	return responses
}

func NewAssignmentHistoryResponses(rows []entities.ContentSubmissionAssignmentHistory) []StatusHistoryResponse {
	responses := make([]StatusHistoryResponse, 0, len(rows))
	for i := range rows {
		note := rows[i].Note
		responses = append(responses, StatusHistoryResponse{EventType: "PJ_ASSIGNED", ID: rows[i].ID, Actor: newUserSummary(rows[i].Actor), Note: note, CreatedAt: rows[i].CreatedAt, FromPJ: newUserSummary(rows[i].FromPJ), ToPJ: newUserSummary(rows[i].ToPJ)})
		if rows[i].FromPJID != nil {
			responses[len(responses)-1].EventType = "PJ_REASSIGNED"
		}
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
