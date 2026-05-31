package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	ValidationError = "VALIDATION_ERROR"
	Unauthenticated = "UNAUTHENTICATED"
	Forbidden       = "FORBIDDEN"
	NotFound        = "NOT_FOUND"
	Conflict        = "CONFLICT"
	RateLimited     = "RATE_LIMITED"
	InternalError   = "INTERNAL_ERROR"
	WAEngineError   = "WA_ENGINE_ERROR"
)

type Meta struct {
	Page       int      `json:"page"`
	PerPage    int      `json:"per_page"`
	Total      int64    `json:"total"`
	TotalPages int      `json:"total_pages"`
	Warnings   []string `json:"warnings,omitempty"`
}

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ErrorBody struct {
	Code    string       `json:"code"`
	Details []FieldError `json:"details,omitempty"`
}

type Envelope struct {
	Success bool       `json:"success"`
	Message string     `json:"message"`
	Data    any        `json:"data,omitempty"`
	Meta    *Meta      `json:"meta,omitempty"`
	Error   *ErrorBody `json:"error,omitempty"`
}

func OK(c *gin.Context, message string, data any) {
	c.JSON(http.StatusOK, Envelope{Success: true, Message: message, Data: data})
}

func Created(c *gin.Context, message string, data any) {
	c.JSON(http.StatusCreated, Envelope{Success: true, Message: message, Data: data})
}

func List(c *gin.Context, message string, data any, meta Meta) {
	c.JSON(http.StatusOK, Envelope{Success: true, Message: message, Data: data, Meta: &meta})
}

func Error(c *gin.Context, status int, code, message string, details ...FieldError) {
	c.JSON(status, Envelope{Success: false, Message: message, Error: &ErrorBody{Code: code, Details: details}})
}
