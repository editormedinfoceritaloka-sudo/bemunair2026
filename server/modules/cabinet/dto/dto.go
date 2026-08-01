package dto

import (
	"time"

	"bemunair2026/server/database/entities"
)

type CabinetRequest struct {
	Name            string  `json:"name"`
	Slug            string  `json:"slug"`
	Tagline         *string `json:"tagline"`
	Description     *string `json:"description"`
	LogoMediaID     *uint64 `json:"logo_media_id"`
	HeroMediaID     *uint64 `json:"hero_media_id"`
	PeriodStart     *string `json:"period_start"`
	PeriodEnd       *string `json:"period_end"`
	IsActive        bool    `json:"is_active"`
	IsPublished     bool    `json:"is_published"`
	MetaTitle       *string `json:"meta_title"`
	MetaDescription *string `json:"meta_description"`
}

type UnitRequest struct {
	CabinetTermID *uint64 `json:"cabinet_term_id"`
	ParentID      *uint64 `json:"parent_id"`
	Code          string  `json:"code"`
	Name          string  `json:"name"`
	UnitType      string  `json:"unit_type"`
	Slug          string  `json:"slug"`
	ShortName     *string `json:"short_name"`
	Description   *string `json:"description"`
	Vision        *string `json:"vision"`
	Mission       *string `json:"mission"`
	LogoMediaID   *uint64 `json:"logo_media_id"`
	CoverMediaID  *uint64 `json:"cover_media_id"`
	DisplayOrder  uint    `json:"display_order"`
	IsActive      bool    `json:"is_active"`
	IsPublished   bool    `json:"is_published"`
}

type MemberRequest struct {
	MinistryID   uint64  `json:"ministry_id"`
	Name         string  `json:"name"`
	Position     string  `json:"position"`
	PositionType string  `json:"position_type"`
	Biography    *string `json:"biography"`
	Quote        *string `json:"quote"`
	PhotoMediaID *uint64 `json:"photo_media_id"`
	DisplayOrder uint    `json:"display_order"`
	IsLeader     bool    `json:"is_leader"`
	IsActive     bool    `json:"is_active"`
}

type ProgramRequest struct {
	MinistryID       uint64  `json:"ministry_id"`
	Name             string  `json:"name"`
	Slug             string  `json:"slug"`
	ShortDescription *string `json:"short_description"`
	Description      *string `json:"description"`
	Objectives       *string `json:"objectives"`
	TargetAudience   *string `json:"target_audience"`
	ExecutionMonth   *string `json:"execution_month"`
	LifecycleStatus  string  `json:"status"`
	CoverMediaID     *uint64 `json:"cover_media_id"`
	DisplayOrder     uint    `json:"display_order"`
	IsFeatured       bool    `json:"is_featured"`
	IsPublished      bool    `json:"is_published"`
}

type MilestoneRequest struct {
	WorkProgramID uint64  `json:"work_program_id"`
	Title         string  `json:"title"`
	Description   *string `json:"description"`
	StartDate     *string `json:"start_date"`
	EndDate       *string `json:"end_date"`
	Status        string  `json:"status"`
	DisplayOrder  uint    `json:"display_order"`
}

type MediaRequest struct {
	UploadedBy     *uint64 `json:"uploaded_by"`
	ImageKitFileID string  `json:"file_id"`
	FilePath       *string `json:"file_path"`
	URL            string  `json:"url"`
	ThumbnailURL   *string `json:"thumbnail_url"`
	Name           string  `json:"name"`
	AltText        string  `json:"alt_text"`
	Caption        *string `json:"caption"`
	MimeType       string  `json:"mime_type"`
	SizeBytes      uint64  `json:"size_bytes"`
	Width          *uint   `json:"width"`
	Height         *uint   `json:"height"`
	Purpose        string  `json:"purpose"`
}

type DocumentationRequest struct {
	MediaAssetID uint64  `json:"media_asset_id"`
	Title        *string `json:"title"`
	Caption      *string `json:"caption"`
	DisplayOrder uint    `json:"display_order"`
	IsCover      bool    `json:"is_cover"`
}

type MediaResponse struct {
	ID           uint64  `json:"id"`
	FileID       string  `json:"file_id"`
	URL          string  `json:"url"`
	ThumbnailURL *string `json:"thumbnail_url,omitempty"`
	Name         string  `json:"name"`
	AltText      string  `json:"alt_text"`
	Caption      *string `json:"caption,omitempty"`
	MimeType     string  `json:"mime_type"`
	SizeBytes    uint64  `json:"size_bytes"`
	Width        *uint   `json:"width,omitempty"`
	Height       *uint   `json:"height,omitempty"`
	Purpose      string  `json:"purpose"`
	Status       string  `json:"status"`
}

type MemberResponse struct {
	ID           uint64         `json:"id"`
	Name         string         `json:"name"`
	Position     string         `json:"position"`
	PositionType string         `json:"position_type"`
	Biography    *string        `json:"biography,omitempty"`
	Quote        *string        `json:"quote,omitempty"`
	Photo        *MediaResponse `json:"photo,omitempty"`
	DisplayOrder uint           `json:"display_order"`
	IsLeader     bool           `json:"is_leader"`
}

type UnitResponse struct {
	ID            uint64           `json:"id"`
	CabinetTermID *uint64          `json:"cabinet_term_id,omitempty"`
	ParentID      *uint64          `json:"parent_id,omitempty"`
	Code          string           `json:"code"`
	Name          string           `json:"name"`
	UnitType      string           `json:"unit_type"`
	Slug          string           `json:"slug"`
	ShortName     *string          `json:"short_name,omitempty"`
	Description   *string          `json:"description,omitempty"`
	Vision        *string          `json:"vision,omitempty"`
	Mission       *string          `json:"mission,omitempty"`
	Logo          *MediaResponse   `json:"logo,omitempty"`
	Cover         *MediaResponse   `json:"cover,omitempty"`
	DisplayOrder  uint             `json:"display_order"`
	IsActive      bool             `json:"is_active"`
	IsPublished   bool             `json:"is_published"`
	Members       []MemberResponse `json:"members,omitempty"`
	Children      []UnitResponse   `json:"children,omitempty"`
}

type MilestoneResponse struct {
	ID           uint64     `json:"id"`
	Title        string     `json:"title"`
	Description  *string    `json:"description,omitempty"`
	StartDate    *time.Time `json:"start_date,omitempty"`
	EndDate      *time.Time `json:"end_date,omitempty"`
	Status       string     `json:"status"`
	DisplayOrder uint       `json:"display_order"`
}

type DocumentationResponse struct {
	ID           uint64         `json:"id"`
	Media        *MediaResponse `json:"media,omitempty"`
	Title        *string        `json:"title,omitempty"`
	Caption      *string        `json:"caption,omitempty"`
	TakenAt      *time.Time     `json:"taken_at,omitempty"`
	DisplayOrder uint           `json:"display_order"`
	IsCover      bool           `json:"is_cover"`
}

type ProgramResponse struct {
	ID               uint64                  `json:"id"`
	MinistryID       uint64                  `json:"ministry_id"`
	MinistryName     string                  `json:"ministry_name,omitempty"`
	Name             string                  `json:"name"`
	Slug             string                  `json:"slug"`
	ShortDescription *string                 `json:"short_description,omitempty"`
	Description      *string                 `json:"description,omitempty"`
	Objectives       *string                 `json:"objectives,omitempty"`
	TargetAudience   *string                 `json:"target_audience,omitempty"`
	ExecutionMonth   *string                 `json:"execution_month,omitempty"`
	Status           string                  `json:"status"`
	Cover            *MediaResponse          `json:"cover,omitempty"`
	DisplayOrder     uint                    `json:"display_order"`
	IsFeatured       bool                    `json:"is_featured"`
	IsPublished      bool                    `json:"is_published"`
	PublishedAt      *time.Time              `json:"published_at,omitempty"`
	Milestones       []MilestoneResponse     `json:"milestones"`
	Documentations   []DocumentationResponse `json:"documentations"`
}

type CabinetResponse struct {
	ID              uint64         `json:"id"`
	Name            string         `json:"name"`
	Slug            string         `json:"slug"`
	Tagline         *string        `json:"tagline,omitempty"`
	Description     *string        `json:"description,omitempty"`
	Logo            *MediaResponse `json:"logo,omitempty"`
	Hero            *MediaResponse `json:"hero,omitempty"`
	PeriodStart     *time.Time     `json:"period_start,omitempty"`
	PeriodEnd       *time.Time     `json:"period_end,omitempty"`
	IsActive        bool           `json:"is_active"`
	IsPublished     bool           `json:"is_published"`
	MetaTitle       *string        `json:"meta_title,omitempty"`
	MetaDescription *string        `json:"meta_description,omitempty"`
	Kemenkoan       []UnitResponse `json:"kemenkoan,omitempty"`
}

func mediaResponse(value *entities.MediaAsset) *MediaResponse {
	if value == nil {
		return nil
	}
	return &MediaResponse{ID: value.ID, FileID: value.ImageKitFileID, URL: value.URL, ThumbnailURL: value.ThumbnailURL, Name: value.Name, AltText: value.AltText, Caption: value.Caption, MimeType: value.MimeType, SizeBytes: value.SizeBytes, Width: value.Width, Height: value.Height, Purpose: value.Purpose, Status: value.Status}
}

func memberResponse(value entities.OrganizationMember) MemberResponse {
	return MemberResponse{ID: value.ID, Name: value.Name, Position: value.Position, PositionType: value.PositionType, Biography: value.Biography, Quote: value.Quote, Photo: mediaResponse(value.PhotoMedia), DisplayOrder: value.DisplayOrder, IsLeader: value.IsLeader}
}

func unitResponse(value entities.Ministry) UnitResponse {
	result := UnitResponse{ID: value.ID, CabinetTermID: value.CabinetTermID, ParentID: value.ParentID, Code: value.Code, Name: value.Name, UnitType: value.UnitType, Slug: value.Slug, ShortName: value.ShortName, Description: value.Description, Vision: value.Vision, Mission: value.Mission, Logo: mediaResponse(value.LogoMedia), Cover: mediaResponse(value.CoverMedia), DisplayOrder: value.DisplayOrder, IsActive: value.IsActive, IsPublished: value.IsPublished}
	for _, member := range value.Members {
		result.Members = append(result.Members, memberResponse(member))
	}
	for _, child := range value.Children {
		result.Children = append(result.Children, unitResponse(child))
	}
	return result
}

func programResponse(value entities.WorkProgram) ProgramResponse {
	result := ProgramResponse{ID: value.ID, MinistryID: value.MinistryID, Name: value.Name, Slug: value.Slug, ShortDescription: value.ShortDescription, Description: value.Description, Objectives: value.Objectives, TargetAudience: value.TargetAudience, ExecutionMonth: value.ExecutionMonth, Status: value.LifecycleStatus, Cover: mediaResponse(value.CoverMedia), DisplayOrder: value.DisplayOrder, IsFeatured: value.IsFeatured, IsPublished: value.IsPublished, PublishedAt: value.PublishedAt, Milestones: make([]MilestoneResponse, 0, len(value.Milestones)), Documentations: make([]DocumentationResponse, 0, len(value.Documentations))}
	if value.Ministry != nil {
		result.MinistryName = value.Ministry.Name
	}
	for _, milestone := range value.Milestones {
		result.Milestones = append(result.Milestones, MilestoneResponse{ID: milestone.ID, Title: milestone.Title, Description: milestone.Description, StartDate: milestone.StartDate, EndDate: milestone.EndDate, Status: milestone.Status, DisplayOrder: milestone.DisplayOrder})
	}
	for _, documentation := range value.Documentations {
		result.Documentations = append(result.Documentations, DocumentationResponse{ID: documentation.ID, Media: mediaResponse(documentation.MediaAsset), Title: documentation.Title, Caption: documentation.Caption, TakenAt: documentation.TakenAt, DisplayOrder: documentation.DisplayOrder, IsCover: documentation.IsCover})
	}
	return result
}

func cabinetResponse(value entities.CabinetTerm, units []entities.Ministry) CabinetResponse {
	result := CabinetResponse{ID: value.ID, Name: value.Name, Slug: value.Slug, Tagline: value.Tagline, Description: value.Description, Logo: mediaResponse(value.LogoMedia), Hero: mediaResponse(value.HeroMedia), PeriodStart: value.PeriodStart, PeriodEnd: value.PeriodEnd, IsActive: value.IsActive, IsPublished: value.IsPublished, MetaTitle: value.MetaTitle, MetaDescription: value.MetaDescription}
	for _, unit := range units {
		if unit.ParentID == nil {
			result.Kemenkoan = append(result.Kemenkoan, unitResponse(unit))
		}
	}
	return result
}
