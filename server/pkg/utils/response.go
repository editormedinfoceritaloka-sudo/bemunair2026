package utils

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

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Error   any    `json:"error,omitempty"`
	Data    any    `json:"data,omitempty"`
	Meta    any    `json:"meta,omitempty"`
}

type EmptyObj struct{}

func BuildResponseSuccess(message string, data any) Response {
	return Response{
		Status:  true,
		Message: message,
		Data:    data,
	}
}

func BuildResponseFailed(message string, err string, data any) Response {
	return Response{
		Status:  false,
		Message: message,
		Error:   err,
		Data:    data,
	}
}

func OK(c *gin.Context, message string, data any) {
	c.JSON(http.StatusOK, BuildResponseSuccess(message, data))
}

func Created(c *gin.Context, message string, data any) {
	c.JSON(http.StatusCreated, BuildResponseSuccess(message, data))
}

func List(c *gin.Context, message string, data any, meta Meta) {
	res := BuildResponseSuccess(message, data)
	res.Meta = meta
	c.JSON(http.StatusOK, res)
}

func Error(c *gin.Context, status int, code, message string, details ...FieldError) {
	res := BuildResponseFailed(message, code, nil)
	if len(details) > 0 {
		res.Error = ErrorBody{Code: code, Details: details}
	}
	c.JSON(status, res)
}
