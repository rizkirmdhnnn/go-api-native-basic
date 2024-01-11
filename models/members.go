package models

import (
	"time"
)

type StatusAccount string

const (
	Allowed    StatusAccount = "Allowed"
	NotAllowed StatusAccount = "NotAllowed"
)

type Member struct {
	ID        uint          `gorm:"primaryKey" json:"id"`
	Name      string        `gorm:"type:varchar(150)" json:"name"`
	Gender    string        `gorm:"type:char(1)" json:"gender"`
	Username  string        `gorm:"type:varchar(100)" json:"username"`
	Email     string        `gorm:"type:varchar(100)" json:"email"`
	Status    StatusAccount `json:"status"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}
