package dto

import (
	"bemunair2026/server/database/entities"
	"bemunair2026/server/modules/medinfo_pj/repository"
)

type CreateRequest struct {
	UserID   uint64 `json:"user_id"`
	Position int    `json:"position"`
}

type ReorderRequest struct {
	IDs []uint64 `json:"ids"`
}

type UserSummary struct {
	ID       uint64  `json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Role     string  `json:"role"`
	Ministry *string `json:"ministry,omitempty"`
	Phone    *string `json:"phone,omitempty"`
}

type QueueResponse struct {
	ID                uint64       `json:"id"`
	UserID            uint64       `json:"user_id"`
	User              *UserSummary `json:"user,omitempty"`
	Position          int          `json:"position"`
	IsCurrent         bool         `json:"is_current"`
	IsBusy            bool         `json:"is_busy"`
	ActiveTaskType    string       `json:"active_task_type,omitempty"`
	ActiveTaskID      uint64       `json:"active_task_id,omitempty"`
	ActiveRequestCode string       `json:"active_request_code,omitempty"`
	ActiveTaskTitle   string       `json:"active_task_title,omitempty"`
}

func NewQueueResponse(row *entities.MedinfoPJQueue) QueueResponse {
	if row == nil {
		return QueueResponse{}
	}

	return QueueResponse{
		ID:        row.ID,
		UserID:    row.UserID,
		User:      newUserSummary(row.User),
		Position:  row.Position,
		IsCurrent: row.IsCurrent,
	}
}

func NewAvailabilityResponses(rows []repository.QueueAvailability) []QueueResponse {
	responses := make([]QueueResponse, 0, len(rows))
	for i := range rows {
		response := NewQueueResponse(&rows[i].Queue)
		response.IsBusy = rows[i].IsBusy
		response.ActiveTaskType = rows[i].ActiveTaskType
		response.ActiveTaskID = rows[i].ActiveTaskID
		response.ActiveRequestCode = rows[i].ActiveRequestCode
		response.ActiveTaskTitle = rows[i].ActiveTaskTitle
		responses = append(responses, response)
	}
	return responses
}

func NewQueueResponses(rows []entities.MedinfoPJQueue) []QueueResponse {
	responses := make([]QueueResponse, 0, len(rows))
	for i := range rows {
		responses = append(responses, NewQueueResponse(&rows[i]))
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
