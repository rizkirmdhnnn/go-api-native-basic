package models

import "time"

type Category struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CategoryResponse struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Description string `json:"category_name"`
}
