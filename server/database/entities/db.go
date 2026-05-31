package entities

import "time"

const (
	RoleAdmin  = "ADMIN"
	RoleMentri = "MENTRI"

	StatusPending  = "PENDING"
	StatusInReview = "IN_REVIEW"
	StatusApproved = "APPROVED"
	StatusRejected = "REJECTED"
)

type User struct {
	ID           uint64    `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"type:varchar(100);not null" json:"name"`
	Email        string    `gorm:"type:varchar(150);uniqueIndex;not null" json:"email"`
	PasswordHash string    `gorm:"column:password_hash;type:varchar(255);not null" json:"-"`
	Role         string    `gorm:"type:enum('ADMIN','MENTRI');not null;index" json:"role"`
	Ministry     *string   `gorm:"type:varchar(100)" json:"ministry"`
	Phone        *string   `gorm:"type:varchar(30)" json:"phone"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (User) TableName() string { return "users" }

type ContentSubmission struct {
	ID             uint64    `gorm:"primaryKey" json:"id"`
	SubmitterID    uint64    `gorm:"column:submitter_id;index;not null" json:"submitter_id"`
	Submitter      *User     `gorm:"foreignKey:SubmitterID;references:ID;constraint:OnDelete:RESTRICT" json:"submitter,omitempty"`
	Ministry       string    `gorm:"type:varchar(100);not null" json:"ministry"`
	Platform       string    `gorm:"type:enum('INSTAGRAM','TWITTER');not null" json:"platform"`
	SubmissionType string    `gorm:"column:submission_type;type:varchar(100);not null" json:"submission_type"`
	Caption        string    `gorm:"type:text" json:"caption"`
	Deadline       time.Time `gorm:"not null;index" json:"deadline"`
	BriefFile      string    `gorm:"column:brief_file;type:varchar(255)" json:"brief_file"`
	PosterFile     string    `gorm:"column:poster_file;type:varchar(255)" json:"poster_file"`
	AssignedPJID   *uint64   `gorm:"column:assigned_pj_id;index" json:"assigned_pj_id"`
	AssignedPJ     *User     `gorm:"foreignKey:AssignedPJID;references:ID;constraint:OnDelete:SET NULL" json:"assigned_pj,omitempty"`
	Status         string    `gorm:"type:enum('PENDING','IN_REVIEW','APPROVED','REJECTED');default:'PENDING';index" json:"status"`
	Notes          *string   `gorm:"type:text" json:"notes"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (ContentSubmission) TableName() string { return "content_submissions" }

type LetterSubmission struct {
	ID           uint64    `gorm:"primaryKey" json:"id"`
	SubmitterID  uint64    `gorm:"column:submitter_id;index;not null" json:"submitter_id"`
	Submitter    *User     `gorm:"foreignKey:SubmitterID;references:ID;constraint:OnDelete:RESTRICT" json:"submitter,omitempty"`
	Ministry     string    `gorm:"type:varchar(100);not null" json:"ministry"`
	LetterType   string    `gorm:"column:letter_type;type:varchar(100);not null" json:"letter_type"`
	Subject      string    `gorm:"type:varchar(200);not null" json:"subject"`
	Body         string    `gorm:"type:text" json:"body"`
	Deadline     time.Time `gorm:"not null;index" json:"deadline"`
	AssignedPJID *uint64   `gorm:"column:assigned_pj_id;index" json:"assigned_pj_id"`
	AssignedPJ   *User     `gorm:"foreignKey:AssignedPJID;references:ID;constraint:OnDelete:SET NULL" json:"assigned_pj,omitempty"`
	Status       string    `gorm:"type:enum('PENDING','IN_REVIEW','APPROVED','REJECTED');default:'PENDING';index" json:"status"`
	Notes        *string   `gorm:"type:text" json:"notes"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (LetterSubmission) TableName() string { return "letter_submissions" }

type MedinfoPJQueue struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	UserID    uint64    `gorm:"column:user_id;uniqueIndex;not null" json:"user_id"`
	User      *User     `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE" json:"user,omitempty"`
	Position  int       `gorm:"not null;index" json:"position"`
	IsCurrent bool      `gorm:"column:is_current;not null;default:false;index" json:"is_current"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (MedinfoPJQueue) TableName() string { return "medinfo_pj_queues" }

type LetterTemplate struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(120);not null" json:"name"`
	Type      string    `gorm:"type:varchar(100);not null;index" json:"type"`
	Subject   string    `gorm:"type:varchar(200);not null" json:"subject"`
	Body      string    `gorm:"type:text;not null" json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (LetterTemplate) TableName() string { return "letter_templates" }
