package user

import "time"

type User struct {
	ID           uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name         string    `json:"name" gorm:"size:120;not null"`
	Email        string    `json:"email" gorm:"size:180;uniqueIndex;not null"`
	PasswordHash string    `json:"-" gorm:"size:255;not null"`
	IsActive     bool      `json:"is_active" gorm:"default:true"`
	AvatarURL    string    `json:"avatar,omitempty" gorm:"size:255"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
