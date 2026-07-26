package entities

type User struct {
	ID           uint64    `gorm:"primaryKey"`
	Name         string    `gorm:"type:varchar(100);not null"`
	Email        string    `gorm:"type:varchar(150);uniqueIndex;not null"`
	PasswordHash string    `gorm:"type:varchar(255);not null"`
	Role         string    `gorm:"type:enum('ADMIN','ADMIN_MEDINFO');not null;index"`
	MinistryID   *uint64   `gorm:"index"`
	MinistryRef  *Ministry `gorm:"foreignKey:MinistryID;references:ID;constraint:OnDelete:SET NULL"`
	// Ministry is retained as a compatibility snapshot for existing API clients.
	Ministry *string `gorm:"type:varchar(100)"`
	Phone    *string `gorm:"type:varchar(30)"`
	Timestamp
}
