package user

import "time"

// UserProfile represents business data of a user
type UserProfile struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"uniqueIndex;not null"`
	FullName  string `gorm:"type:varchar(100)"`
	Phone     string `gorm:"type:varchar(20)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Address represents a shipping/billing address
type Address struct {
	ID         uint   `gorm:"primaryKey"`
	UserID     uint   `gorm:"index;not null"`
	Label      string `gorm:"type:varchar(50)"` // Home, Work
	Line1      string `gorm:"type:varchar(255);not null"`
	Line2      string `gorm:"type:varchar(255)"`
	City       string `gorm:"type:varchar(100);not null"`
	State      string `gorm:"type:varchar(100);not null"`
	PostalCode string `gorm:"type:varchar(20);not null"`
	Country    string `gorm:"type:varchar(100);not null"`
	IsDefault  bool   `gorm:"default:false"`
	CreatedAt  time.Time
}

type UserWithRelations struct {
	UserProfile UserProfile `gorm:"foreignKey:UserID"`
	Addresses   []Address   `gorm:"foreignKey:UserID"`
}
