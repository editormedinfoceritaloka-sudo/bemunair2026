package entities

type LetterTemplate struct {
	ID           uint64      `gorm:"primaryKey"`
	Name         string      `gorm:"type:varchar(120);not null"`
	Type         string      `gorm:"type:varchar(100);not null;index"`
	Subject      string      `gorm:"type:varchar(200);not null"`
	Body         string      `gorm:"type:text;not null"`
	MediaAssetID *uint64     `gorm:"index" json:"media_asset_id,omitempty"`
	MediaAsset   *MediaAsset `gorm:"foreignKey:MediaAssetID;references:ID;constraint:OnDelete:SET NULL" json:"media_asset,omitempty"`
	IsActive     bool        `gorm:"not null;default:true;index" json:"is_active"`
	DisplayOrder uint        `gorm:"not null;default:0;index" json:"display_order"`
	Timestamp
}
