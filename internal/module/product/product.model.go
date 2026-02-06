package product

import "time"

type Category struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Slug      string `gorm:"type:varchar(120);uniqueIndex;not null"`
	CreatedAt time.Time
}

type Product struct {
	ID          uint   `gorm:"primaryKey"`
	CategoryID  uint   `gorm:"index;not null"`
	Name        string `gorm:"type:varchar(200);not null"`
	Slug        string `gorm:"type:varchar(220);uniqueIndex;not null"`
	Description string `gorm:"type:text"`
	Price       int64  `gorm:"not null"` // stored in paise/cents
	IsActive    bool   `gorm:"default:true"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
