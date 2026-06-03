package entities

type User struct {
	ID           uint64    `gorm:"primaryKey"`
	Name         string    `gorm:"type:varchar(100);not null"`
	Email        string    `gorm:"type:varchar(150);uniqueIndex;not null"`
	PasswordHash string    `gorm:"type:varchar(255);not null"`
	Role         string    `gorm:"type:enum('ADMIN','MENTRI');not null;index"`
	Ministry     *string   `gorm:"type:varchar(100)"`
	Phone        *string   `gorm:"type:varchar(30)"`
	Timestamp
}