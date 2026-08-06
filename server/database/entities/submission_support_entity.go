package entities

import "time"

type Ministry struct {
	ID            uint64               `gorm:"primaryKey" json:"id"`
	Code          string               `gorm:"type:varchar(50);uniqueIndex;not null" json:"code"`
	Name          string               `gorm:"type:varchar(120);uniqueIndex;not null" json:"name"`
	CabinetTermID *uint64              `gorm:"index" json:"cabinet_term_id,omitempty"`
	CabinetTerm   *CabinetTerm         `gorm:"foreignKey:CabinetTermID;references:ID;constraint:OnDelete:RESTRICT" json:"-"`
	ParentID      *uint64              `gorm:"index" json:"parent_id,omitempty"`
	Parent        *Ministry            `gorm:"foreignKey:ParentID;references:ID;constraint:OnDelete:RESTRICT" json:"-"`
	Children      []Ministry           `gorm:"foreignKey:ParentID" json:"children,omitempty"`
	Members       []OrganizationMember `gorm:"foreignKey:MinistryID" json:"members,omitempty"`
	Programs      []WorkProgram        `gorm:"foreignKey:MinistryID" json:"programs,omitempty"`
	UnitType      string               `gorm:"type:varchar(20);not null;default:'KEMENTERIAN';index" json:"unit_type"`
	Slug          string               `gorm:"type:varchar(160);index" json:"slug"`
	ShortName     *string              `gorm:"type:varchar(80)" json:"short_name,omitempty"`
	Description   *string              `gorm:"type:text" json:"description,omitempty"`
	Vision        *string              `gorm:"type:text" json:"vision,omitempty"`
	Mission       *string              `gorm:"type:text" json:"mission,omitempty"`
	LogoMediaID   *uint64              `gorm:"index" json:"logo_media_id,omitempty"`
	LogoMedia     *MediaAsset          `gorm:"foreignKey:LogoMediaID;references:ID;constraint:OnDelete:RESTRICT" json:"logo_media,omitempty"`
	CoverMediaID  *uint64              `gorm:"index" json:"cover_media_id,omitempty"`
	CoverMedia    *MediaAsset          `gorm:"foreignKey:CoverMediaID;references:ID;constraint:OnDelete:RESTRICT" json:"cover_media,omitempty"`
	DisplayOrder  uint                 `gorm:"not null;default:0;index" json:"display_order"`
	IsActive      bool                 `gorm:"not null;default:true;index" json:"is_active"`
	IsPublished   bool                 `gorm:"not null;default:false;index" json:"is_published"`
	PublishedAt   *time.Time           `json:"published_at,omitempty"`
	Timestamp
}

type MediaSubmissionSetting struct {
	ID                  uint64 `gorm:"primaryKey"`
	ServiceType         string `gorm:"type:enum('CONTENT','ARTICLE');uniqueIndex;not null"`
	SOPURL              *string
	MinistryTemplateURL *string
	BriefTemplateURL    *string
	CaptionTemplateURL  *string
	PICName             *string
	PICWhatsApp         *string
	TermsJSON           []byte `gorm:"type:json;not null"`
	MinimumLeadDays     uint
	PublishTimeStart    string
	PublishTimeEnd      string
	SlotIntervalMinutes uint
	DailyCapacity       *uint
	Timestamp
}

type PublicationBlackout struct {
	ID           uint64 `gorm:"primaryKey"`
	ServiceType  *string
	BlackoutDate time.Time `gorm:"type:date"`
	Reason       *string
	CreatedBy    uint64
	CreatedAt    time.Time
}
