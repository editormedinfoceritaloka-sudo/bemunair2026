package submission_support

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"bemunair2026/server/database/entities"
	"bemunair2026/server/middlewares"
	"bemunair2026/server/pkg/constants"
	response "bemunair2026/server/pkg/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type controller struct {
	db *gorm.DB
}

type ministryRequest struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	IsActive *bool  `json:"is_active"`
}

type settingRequest struct {
	SOPURL              *string  `json:"sop_url"`
	MinistryTemplateURL *string  `json:"ministry_template_url"`
	BriefTemplateURL    *string  `json:"brief_template_url"`
	CaptionTemplateURL  *string  `json:"caption_template_url"`
	PICName             *string  `json:"pic_name"`
	PICWhatsApp         *string  `json:"pic_whatsapp"`
	Terms               []string `json:"terms"`
	MinimumLeadDays     uint     `json:"minimum_lead_days"`
	PublishTimeStart    string   `json:"publish_time_start"`
	PublishTimeEnd      string   `json:"publish_time_end"`
	SlotIntervalMinutes uint     `json:"slot_interval_minutes"`
	DailyCapacity       *uint    `json:"daily_capacity"`
}

type settingResponse struct {
	entities.MediaSubmissionSetting
	Terms []string `json:"terms"`
}

func RegisterRoutes(api *gin.RouterGroup, db *gorm.DB, jwtSecret string) {
	c := &controller{db: db}
	authenticated := api.Group("", middlewares.Auth(jwtSecret), middlewares.AuthenticatedAdmin())
	authenticated.GET("/ministries", c.listMinistries)
	authenticated.GET("/media-submission-settings/:serviceType", c.getSetting)

	medinfo := authenticated.Group("", middlewares.MedinfoOnly())
	medinfo.POST("/ministries", c.createMinistry)
	medinfo.PUT("/ministries/:id", c.updateMinistry)
	medinfo.PUT("/media-submission-settings/:serviceType", c.updateSetting)
}

func (c *controller) listMinistries(ctx *gin.Context) {
	var rows []entities.Ministry
	query := c.db.Order("name ASC")
	if middlewares.CurrentClaims(ctx).Role != constants.RoleAdminMedinfo {
		query = query.Where("is_active = ?", true)
	}
	if err := query.Find(&rows).Error; err != nil {
		response.Error(ctx, http.StatusInternalServerError, response.InternalError, err.Error())
		return
	}
	response.OK(ctx, "Daftar kementerian", rows)
}

func (c *controller) createMinistry(ctx *gin.Context) {
	var req ministryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil || strings.TrimSpace(req.Name) == "" {
		response.Error(ctx, http.StatusUnprocessableEntity, response.ValidationError, "Nama kementerian wajib diisi")
		return
	}
	active := true
	if req.IsActive != nil {
		active = *req.IsActive
	}
	row := entities.Ministry{
		Code: strings.ToUpper(strings.ReplaceAll(strings.TrimSpace(req.Code), " ", "_")),
		Name: strings.TrimSpace(req.Name), IsActive: active,
	}
	if row.Code == "" {
		row.Code = strings.ToUpper(strings.ReplaceAll(row.Name, " ", "_"))
	}
	if err := c.db.Create(&row).Error; err != nil {
		response.Error(ctx, http.StatusConflict, response.ValidationError, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, response.BuildResponseSuccess("Kementerian dibuat", row))
}

func (c *controller) updateMinistry(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	var req ministryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, response.ValidationError, "Payload tidak valid")
		return
	}
	updates := map[string]any{}
	if strings.TrimSpace(req.Code) != "" {
		updates["code"] = strings.ToUpper(strings.ReplaceAll(strings.TrimSpace(req.Code), " ", "_"))
	}
	if strings.TrimSpace(req.Name) != "" {
		updates["name"] = strings.TrimSpace(req.Name)
	}
	if req.IsActive != nil {
		updates["is_active"] = *req.IsActive
	}
	if len(updates) == 0 {
		response.Error(ctx, http.StatusUnprocessableEntity, response.ValidationError, "Tidak ada perubahan")
		return
	}
	result := c.db.Model(&entities.Ministry{}).Where("id = ?", id).Updates(updates)
	if result.Error != nil {
		response.Error(ctx, http.StatusConflict, response.ValidationError, result.Error.Error())
		return
	}
	if result.RowsAffected == 0 {
		response.Error(ctx, http.StatusNotFound, response.NotFound, "Kementerian tidak ditemukan")
		return
	}
	var row entities.Ministry
	c.db.First(&row, id)
	response.OK(ctx, "Kementerian diperbarui", row)
}

func (c *controller) getSetting(ctx *gin.Context) {
	serviceType := strings.ToUpper(ctx.Param("serviceType"))
	if serviceType != constants.ServiceTypeContent && serviceType != constants.ServiceTypeArticle {
		response.Error(ctx, http.StatusUnprocessableEntity, response.ValidationError, "serviceType tidak valid")
		return
	}
	var row entities.MediaSubmissionSetting
	if err := c.db.Where("service_type = ?", serviceType).First(&row).Error; err != nil {
		response.Error(ctx, http.StatusNotFound, response.NotFound, "Pengaturan tidak ditemukan")
		return
	}
	response.OK(ctx, "Pengaturan pengajuan media", toSettingResponse(row))
}

func (c *controller) updateSetting(ctx *gin.Context) {
	serviceType := strings.ToUpper(ctx.Param("serviceType"))
	if serviceType != constants.ServiceTypeContent && serviceType != constants.ServiceTypeArticle {
		response.Error(ctx, http.StatusUnprocessableEntity, response.ValidationError, "serviceType tidak valid")
		return
	}
	var req settingRequest
	if err := ctx.ShouldBindJSON(&req); err != nil || req.PublishTimeStart == "" || req.PublishTimeEnd == "" {
		response.Error(ctx, http.StatusUnprocessableEntity, response.ValidationError, "Payload pengaturan tidak valid")
		return
	}
	terms, _ := json.Marshal(req.Terms)
	updates := map[string]any{
		"sop_url": req.SOPURL, "ministry_template_url": req.MinistryTemplateURL,
		"brief_template_url": req.BriefTemplateURL, "caption_template_url": req.CaptionTemplateURL,
		"pic_name": req.PICName, "pic_whatsapp": req.PICWhatsApp, "terms_json": terms,
		"minimum_lead_days": req.MinimumLeadDays, "publish_time_start": req.PublishTimeStart,
		"publish_time_end": req.PublishTimeEnd, "slot_interval_minutes": req.SlotIntervalMinutes,
		"daily_capacity": req.DailyCapacity,
	}
	if err := c.db.Model(&entities.MediaSubmissionSetting{}).Where("service_type = ?", serviceType).Updates(updates).Error; err != nil {
		response.Error(ctx, http.StatusInternalServerError, response.InternalError, err.Error())
		return
	}
	var row entities.MediaSubmissionSetting
	c.db.Where("service_type = ?", serviceType).First(&row)
	response.OK(ctx, "Pengaturan diperbarui", toSettingResponse(row))
}

func toSettingResponse(row entities.MediaSubmissionSetting) settingResponse {
	var terms []string
	_ = json.Unmarshal(row.TermsJSON, &terms)
	return settingResponse{MediaSubmissionSetting: row, Terms: terms}
}
