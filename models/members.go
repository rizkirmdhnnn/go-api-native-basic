package models

import (
	"time"
)

type Member struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(150)" json:"name"`
	Gender    string    `gorm:"type:char(1)" json:"gender"`
	Username  string    `gorm:"type:varchar(100)" json:"username"`
	Email     string    `gorm:"type:varchar(100)" json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type MemberTransactionResponse struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"type:varchar(150)" json:"name"`
	Gender   string `gorm:"type:char(1)" json:"gender"`
	Username string `gorm:"type:varchar(100)" json:"username"`
	Email    string `gorm:"type:varchar(100)" json:"email"`
}
