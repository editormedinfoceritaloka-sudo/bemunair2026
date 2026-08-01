package entities

import "time"

type CabinetTerm struct {
	ID              uint64      `gorm:"primaryKey" json:"id"`
	Name            string      `gorm:"type:varchar(150);not null" json:"name"`
	Slug            string      `gorm:"type:varchar(160);uniqueIndex;not null" json:"slug"`
	Tagline         *string     `gorm:"type:varchar(255)" json:"tagline,omitempty"`
	Description     *string     `gorm:"type:longtext" json:"description,omitempty"`
	LogoMediaID     *uint64     `gorm:"index" json:"logo_media_id,omitempty"`
	LogoMedia       *MediaAsset `gorm:"foreignKey:LogoMediaID;references:ID;constraint:OnDelete:RESTRICT" json:"logo_media,omitempty"`
	HeroMediaID     *uint64     `gorm:"index" json:"hero_media_id,omitempty"`
	HeroMedia       *MediaAsset `gorm:"foreignKey:HeroMediaID;references:ID;constraint:OnDelete:RESTRICT" json:"hero_media,omitempty"`
	PeriodStart     *time.Time  `gorm:"type:date" json:"period_start,omitempty"`
	PeriodEnd       *time.Time  `gorm:"type:date" json:"period_end,omitempty"`
	IsActive        bool        `gorm:"not null;default:false;index" json:"is_active"`
	IsPublished     bool        `gorm:"not null;default:false;index" json:"is_published"`
	PublishedAt     *time.Time  `json:"published_at,omitempty"`
	MetaTitle       *string     `gorm:"type:varchar(180)" json:"meta_title,omitempty"`
	MetaDescription *string     `gorm:"type:varchar(255)" json:"meta_description,omitempty"`
	Timestamp
	Ministries []Ministry `gorm:"foreignKey:CabinetTermID" json:"ministries,omitempty"`
}
