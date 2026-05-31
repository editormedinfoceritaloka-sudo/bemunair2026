package letter_submission

import (
	"net/http"
	"strconv"
	"time"

	"bemunair2026/server/database/entities"
	"bemunair2026/server/middlewares"
	"bemunair2026/server/modules/user"
	"bemunair2026/server/pkg/response"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	repo    *Repository
	service *Service
	users   *user.Repository
}

func NewHandler(repo *Repository, service *Service, users *user.Repository) *Handler {
	return &Handler{repo: repo, service: service, users: users}
}

func (h *Handler) Create(c *gin.Context) {
	claims := middlewares.CurrentClaims(c)
	submitter, _ := h.users.FindByID(claims.UserID)
	var req struct{ Ministry, LetterType, Subject, Body, Deadline string }
	if c.ShouldBindJSON(&req) != nil {
		response.Error(c, http.StatusUnprocessableEntity, response.ValidationError, "Validasi gagal")
		return
	}
	deadline, err := time.Parse(time.RFC3339, req.Deadline)
	if err != nil {
		response.Error(c, http.StatusUnprocessableEntity, response.ValidationError, "Deadline harus RFC3339")
		return
	}
	if req.Ministry == "" && claims.Ministry != nil {
		req.Ministry = *claims.Ministry
	}
	sub := &entities.LetterSubmission{SubmitterID: claims.UserID, Ministry: req.Ministry, LetterType: req.LetterType, Subject: req.Subject, Body: req.Body, Deadline: deadline, Status: entities.StatusPending}
	created, warnings, err := h.service.Create(sub, submitter)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.InternalError, "Submission gagal dibuat")
		return
	}
	if len(warnings) > 0 {
		gin.DefaultWriter.Write([]byte("WA warning: " + warnings[0].Error() + "\n"))
	}
	response.Created(c, "Submission surat berhasil dibuat", created)
}

func (h *Handler) List(c *gin.Context) {
	claims := middlewares.CurrentClaims(c)
	rows, err := h.repo.ListForUser(claims.Role, claims.UserID, claims.Ministry)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.InternalError, "Gagal mengambil submission")
		return
	}
	response.List(c, "Daftar letter submission", rows, response.Meta{Page: 1, PerPage: len(rows), Total: int64(len(rows)), TotalPages: 1})
}

func (h *Handler) Get(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	row, err := h.repo.FindByID(id)
	if err != nil || row == nil {
		response.Error(c, http.StatusNotFound, response.NotFound, "Submission tidak ditemukan")
		return
	}
	response.OK(c, "Detail letter submission", row)
}

func (h *Handler) UpdateStatus(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var req struct {
		Status string  `json:"status"`
		Notes  *string `json:"notes"`
	}
	if c.ShouldBindJSON(&req) != nil || req.Status == "" {
		response.Error(c, http.StatusUnprocessableEntity, response.ValidationError, "Validasi gagal")
		return
	}
	row, err := h.service.UpdateStatus(id, req.Status, req.Notes)
	if err != nil {
		response.Error(c, http.StatusConflict, response.Conflict, err.Error())
		return
	}
	response.OK(c, "Status berhasil diperbarui", row)
}

func (h *Handler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.repo.Delete(id); err != nil {
		response.Error(c, http.StatusInternalServerError, response.InternalError, "Submission gagal dihapus")
		return
	}
	response.OK(c, "Submission berhasil dihapus", nil)
}
