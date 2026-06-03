package entities

import "time"

type LetterSubmission struct {
	ID           uint64    `gorm:"primaryKey"`
	SubmitterID  uint64    `gorm:"index;not null"`
	Submitter    *User     `gorm:"foreignKey:SubmitterID;references:ID;constraint:OnDelete:RESTRICT"`
	Ministry     string    `gorm:"type:varchar(100);not null"`
	LetterType   string    `gorm:"type:varchar(100);not null"`
	Subject      string    `gorm:"type:varchar(200);not null"`
	Body         string    `gorm:"type:text"`
	Deadline     time.Time `gorm:"not null;index"`
	AssignedPJID *uint64   `gorm:"index"`
	AssignedPJ   *User     `gorm:"foreignKey:AssignedPJID;references:ID;constraint:OnDelete:SET NULL"`
	Status       string    `gorm:"type:enum('PENDING','IN_REVIEW','APPROVED','REJECTED');default:'PENDING';index"`
	Notes        *string   `gorm:"type:text"`
	Timestamp
}
