package entities

import "time"

type LetterSubmission struct {
	ID             uint64    `gorm:"primaryKey"`
	RequestCode    *string   `gorm:"type:varchar(30);uniqueIndex"`
	SubmitterID    uint64    `gorm:"index;not null"`
	Submitter      *User     `gorm:"foreignKey:SubmitterID;references:ID;constraint:OnDelete:RESTRICT"`
	SubmitterName  string    `gorm:"type:varchar(100);not null"`
	SubmitterPhone *string   `gorm:"type:varchar(30)"`
	Ministry       string    `gorm:"type:varchar(100);not null"`
	MinistryID     *uint64   `gorm:"index"`
	MinistryRef    *Ministry `gorm:"foreignKey:MinistryID;references:ID;constraint:OnDelete:SET NULL"`
	LetterType     string    `gorm:"type:varchar(100);not null"`
	Subject        string    `gorm:"type:varchar(200);not null"`
	Body           string    `gorm:"type:text"`
	Deadline       time.Time `gorm:"not null;index"`
	SubmittedAt    *time.Time
	AssignedPJID   *uint64 `gorm:"index"`
	AssignedPJ     *User   `gorm:"foreignKey:AssignedPJID;references:ID;constraint:OnDelete:SET NULL"`
	Status         string  `gorm:"type:varchar(40);default:'DRAFT';index"`
	Notes          *string `gorm:"type:text"`
	Timestamp
}

type LetterSubmissionStatusHistory struct {
	ID           uint64 `gorm:"primaryKey"`
	SubmissionID uint64 `gorm:"index;not null"`
	Submission   *LetterSubmission
	ActorID      *uint64 `gorm:"index"`
	Actor        *User
	FromStatus   *string `gorm:"type:varchar(40)"`
	ToStatus     string  `gorm:"type:varchar(40);not null"`
	Note         *string `gorm:"type:text"`
	CreatedAt    time.Time
}
