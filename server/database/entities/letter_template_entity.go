package entities

type LetterTemplate struct {
	ID      uint64 `gorm:"primaryKey"`
	Name    string `gorm:"type:varchar(120);not null"`
	Type    string `gorm:"type:varchar(100);not null;index"`
	Subject string `gorm:"type:varchar(200);not null"`
	Body    string `gorm:"type:text;not null"`
	Timestamp
}
