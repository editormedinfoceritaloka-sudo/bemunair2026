package entities

import "time"

type Article struct {
	ID          uint64     `gorm:"primaryKey"`
	Slug        string     `gorm:"type:varchar(255);uniqueIndex;not null"`
	Title       string     `gorm:"type:varchar(255);not null"`
	Excerpt     *string    `gorm:"type:varchar(500)"`
	Body        string     `gorm:"type:longtext;not null"`
	CoverImage  *string    `gorm:"type:varchar(500)"`
	AuthorID    uint64     `gorm:"index;not null"`
	Author      *User      `gorm:"foreignKey:AuthorID;references:ID;constraint:OnDelete:RESTRICT"`
	Status      string     `gorm:"type:enum('DRAFT','PUBLISHED');default:'DRAFT';index"`
	PublishedAt *time.Time `gorm:"index"`
	Timestamp
}
