package entities

type MedinfoPJQueue struct {
	ID        uint64 `gorm:"primaryKey"`
	UserID    uint64 `gorm:"uniqueIndex;not null"`
	User      *User  `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Position  int    `gorm:"not null;index"`
	IsCurrent bool   `gorm:"not null;default:false;index"`
	Timestamp
}
