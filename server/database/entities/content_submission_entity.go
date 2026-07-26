package entities

import "time"

type ContentSubmission struct {
	ID                     uint64     `gorm:"primaryKey"`
	RequestCode            *string    `gorm:"type:varchar(30);uniqueIndex"`
	SubmitterID            uint64     `gorm:"index;not null"`
	Submitter              *User      `gorm:"foreignKey:SubmitterID;references:ID;constraint:OnDelete:RESTRICT"`
	SubmitterName          string     `gorm:"type:varchar(100);not null"`
	SubmitterPhone         *string    `gorm:"type:varchar(30)"`
	Ministry               string     `gorm:"type:varchar(100);not null"`
	MinistryID             *uint64    `gorm:"index"`
	MinistryRef            *Ministry  `gorm:"foreignKey:MinistryID;references:ID;constraint:OnDelete:SET NULL"`
	ServiceType            string     `gorm:"type:enum('CONTENT','ARTICLE');not null;index"`
	ContentFormat          *string    `gorm:"type:varchar(50)"`
	SubmissionType         string     `gorm:"type:varchar(100);not null"`
	Title                  string     `gorm:"type:varchar(255);not null"`
	AddSong                *string    `gorm:"type:varchar(255)"`
	SongTitle              *string    `gorm:"type:varchar(180)"`
	SongArtist             *string    `gorm:"type:varchar(180)"`
	SongStartSeconds       *uint      `gorm:"type:int unsigned"`
	SongEndSeconds         *uint      `gorm:"type:int unsigned"`
	Caption                string     `gorm:"type:text"`
	AdditionalNotes        *string    `gorm:"type:text"`
	PublishDate            *time.Time `gorm:"type:date"`
	PublishTime            *string    `gorm:"type:varchar(5)"`
	DesignDriveLink        *string    `gorm:"type:varchar(500)"`
	CanvaLink              *string    `gorm:"type:varchar(500)"`
	ArticleDriveLink       *string    `gorm:"type:varchar(500)"`
	DocumentationDriveLink *string    `gorm:"type:varchar(500)"`
	RequiredInformation    *string    `gorm:"type:text"`
	Deadline               *time.Time `gorm:"index"`
	ConfirmedPublishAt     *time.Time
	SubmittedAt            *time.Time
	BriefLink              string  `gorm:"type:varchar(500)"`
	AssignedPJID           *uint64 `gorm:"index"`
	AssignedPJ             *User   `gorm:"foreignKey:AssignedPJID;references:ID;constraint:OnDelete:SET NULL"`
	Status                 string  `gorm:"type:varchar(40);default:'DRAFT';index"`
	Notes                  *string `gorm:"type:text"`
	Timestamp
}

type ContentSubmissionStatusHistory struct {
	ID           uint64 `gorm:"primaryKey"`
	SubmissionID uint64 `gorm:"index;not null"`
	Submission   *ContentSubmission
	ActorID      *uint64 `gorm:"index"`
	Actor        *User
	FromStatus   *string `gorm:"type:varchar(40)"`
	ToStatus     string  `gorm:"type:varchar(40);not null"`
	Note         *string `gorm:"type:text"`
	CreatedAt    time.Time
}

type ContentSubmissionAttachment struct {
	ID             uint64 `gorm:"primaryKey"`
	SubmissionID   *uint64
	UploadedBy     uint64
	ImageKitFileID string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Purpose        string `gorm:"type:varchar(40);not null"`
	Name           string `gorm:"type:varchar(255);not null"`
	URL            string `gorm:"type:varchar(1000);not null"`
	MimeType       string `gorm:"type:varchar(120);not null"`
	SizeBytes      uint64
	Status         string `gorm:"type:varchar(20);not null"`
	Timestamp
}
