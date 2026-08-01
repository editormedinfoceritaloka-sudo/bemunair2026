package entities

import "time"

type MediaAsset struct {
	ID             uint64  `gorm:"primaryKey" json:"id"`
	UploadedBy     *uint64 `gorm:"index" json:"uploaded_by,omitempty"`
	ImageKitFileID string  `gorm:"column:imagekit_file_id;type:varchar(100);uniqueIndex;not null" json:"file_id"`
	FilePath       *string `gorm:"type:varchar(1000)" json:"file_path,omitempty"`
	URL            string  `gorm:"type:varchar(1000);not null" json:"url"`
	ThumbnailURL   *string `gorm:"type:varchar(1000)" json:"thumbnail_url,omitempty"`
	Name           string  `gorm:"type:varchar(255);not null" json:"name"`
	AltText        string  `gorm:"type:varchar(255);not null" json:"alt_text"`
	Caption        *string `gorm:"type:text" json:"caption,omitempty"`
	MimeType       string  `gorm:"type:varchar(120);not null" json:"mime_type"`
	SizeBytes      uint64  `gorm:"not null" json:"size_bytes"`
	Width          *uint   `json:"width,omitempty"`
	Height         *uint   `json:"height,omitempty"`
	Purpose        string  `gorm:"type:varchar(80);not null;index" json:"purpose"`
	Status         string  `gorm:"type:varchar(20);not null;default:'ACTIVE';index" json:"status"`
	Timestamp
}

type OrganizationMember struct {
	ID           uint64      `gorm:"primaryKey" json:"id"`
	MinistryID   uint64      `gorm:"index;not null" json:"ministry_id"`
	Ministry     *Ministry   `gorm:"foreignKey:MinistryID;references:ID;constraint:OnDelete:RESTRICT" json:"-"`
	UserID       *uint64     `gorm:"index" json:"user_id,omitempty"`
	User         *User       `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:SET NULL" json:"-"`
	Name         string      `gorm:"type:varchar(120);not null" json:"name"`
	Position     string      `gorm:"type:varchar(120);not null" json:"position"`
	PositionType string      `gorm:"type:varchar(40);not null;index" json:"position_type"`
	EmailPublic  *string     `gorm:"type:varchar(150)" json:"email_public,omitempty"`
	PhonePublic  *string     `gorm:"type:varchar(30)" json:"phone_public,omitempty"`
	Biography    *string     `gorm:"type:text" json:"biography,omitempty"`
	Quote        *string     `gorm:"type:text" json:"quote,omitempty"`
	PhotoMediaID *uint64     `gorm:"index" json:"photo_media_id,omitempty"`
	PhotoMedia   *MediaAsset `gorm:"foreignKey:PhotoMediaID;references:ID;constraint:OnDelete:RESTRICT" json:"photo_media,omitempty"`
	DisplayOrder uint        `gorm:"not null;default:0;index" json:"display_order"`
	IsLeader     bool        `gorm:"not null;default:false;index" json:"is_leader"`
	IsActive     bool        `gorm:"not null;default:true;index" json:"is_active"`
	StartedAt    *time.Time  `gorm:"type:date" json:"started_at,omitempty"`
	EndedAt      *time.Time  `gorm:"type:date" json:"ended_at,omitempty"`
	Timestamp
}

type UserOrganizationRole struct {
	ID         uint64    `gorm:"primaryKey" json:"id"`
	UserID     uint64    `gorm:"not null;index" json:"user_id"`
	MinistryID uint64    `gorm:"not null;index" json:"ministry_id"`
	Permission string    `gorm:"type:varchar(40);not null" json:"permission"`
	IsActive   bool      `gorm:"not null;default:true;index" json:"is_active"`
	User       *User     `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE" json:"user,omitempty"`
	Ministry   *Ministry `gorm:"foreignKey:MinistryID;references:ID;constraint:OnDelete:RESTRICT" json:"ministry,omitempty"`
	Timestamp
}

type WorkProgram struct {
	ID               uint64      `gorm:"primaryKey" json:"id"`
	MinistryID       uint64      `gorm:"not null;index" json:"ministry_id"`
	Ministry         *Ministry   `gorm:"foreignKey:MinistryID;references:ID;constraint:OnDelete:RESTRICT" json:"ministry,omitempty"`
	Name             string      `gorm:"type:varchar(180);not null" json:"name"`
	Slug             string      `gorm:"type:varchar(200);not null" json:"slug"`
	ShortDescription *string     `gorm:"type:varchar(500)" json:"short_description,omitempty"`
	Description      *string     `gorm:"type:longtext" json:"description,omitempty"`
	Objectives       *string     `gorm:"type:text" json:"objectives,omitempty"`
	TargetAudience   *string     `gorm:"type:text" json:"target_audience,omitempty"`
	StartDate        *time.Time  `gorm:"type:date" json:"start_date,omitempty"`
	EndDate          *time.Time  `gorm:"type:date" json:"end_date,omitempty"`
	ExecutionMonth   *string     `gorm:"type:varchar(50)" json:"execution_month,omitempty"`
	LifecycleStatus  string      `gorm:"type:varchar(20);not null;default:'DRAFT';index" json:"status"`
	CoverMediaID     *uint64     `gorm:"index" json:"cover_media_id,omitempty"`
	CoverMedia       *MediaAsset `gorm:"foreignKey:CoverMediaID;references:ID;constraint:OnDelete:RESTRICT" json:"cover_media,omitempty"`
	DisplayOrder     uint        `gorm:"not null;default:0;index" json:"display_order"`
	IsFeatured       bool        `gorm:"not null;default:false;index" json:"is_featured"`
	IsPublished      bool        `gorm:"not null;default:false;index" json:"is_published"`
	PublishedAt      *time.Time  `json:"published_at,omitempty"`
	CreatedBy        *uint64     `gorm:"index" json:"created_by,omitempty"`
	UpdatedBy        *uint64     `gorm:"index" json:"updated_by,omitempty"`
	MetaTitle        *string     `gorm:"type:varchar(180)" json:"meta_title,omitempty"`
	MetaDescription  *string     `gorm:"type:varchar(255)" json:"meta_description,omitempty"`
	Timestamp
	Milestones     []WorkProgramMilestone     `gorm:"foreignKey:WorkProgramID" json:"milestones,omitempty"`
	Documentations []WorkProgramDocumentation `gorm:"foreignKey:WorkProgramID" json:"documentations,omitempty"`
}

type WorkProgramMilestone struct {
	ID            uint64     `gorm:"primaryKey" json:"id"`
	WorkProgramID uint64     `gorm:"not null;index" json:"work_program_id"`
	Title         string     `gorm:"type:varchar(180);not null" json:"title"`
	Description   *string    `gorm:"type:text" json:"description,omitempty"`
	StartDate     *time.Time `gorm:"type:date" json:"start_date,omitempty"`
	EndDate       *time.Time `gorm:"type:date" json:"end_date,omitempty"`
	Status        string     `gorm:"type:varchar(20);not null;default:'PLANNED';index" json:"status"`
	DisplayOrder  uint       `gorm:"not null;default:0;index" json:"display_order"`
	Timestamp
}

type WorkProgramDocumentation struct {
	ID            uint64      `gorm:"primaryKey" json:"id"`
	WorkProgramID uint64      `gorm:"not null;index" json:"work_program_id"`
	MediaAssetID  uint64      `gorm:"not null;index" json:"media_asset_id"`
	MediaAsset    *MediaAsset `gorm:"foreignKey:MediaAssetID;references:ID;constraint:OnDelete:RESTRICT" json:"media,omitempty"`
	Title         *string     `gorm:"type:varchar(180)" json:"title,omitempty"`
	Caption       *string     `gorm:"type:text" json:"caption,omitempty"`
	TakenAt       *time.Time  `json:"taken_at,omitempty"`
	DisplayOrder  uint        `gorm:"not null;default:0;index" json:"display_order"`
	IsCover       bool        `gorm:"not null;default:false;index" json:"is_cover"`
	Timestamp
}
