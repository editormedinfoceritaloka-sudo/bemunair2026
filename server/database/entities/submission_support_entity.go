package entities

import "time"

type Ministry struct {
	ID       uint64 `gorm:"primaryKey" json:"id"`
	Code     string `gorm:"type:varchar(50);uniqueIndex;not null" json:"code"`
	Name     string `gorm:"type:varchar(120);uniqueIndex;not null" json:"name"`
	IsActive bool   `gorm:"not null;default:true;index" json:"is_active"`
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
