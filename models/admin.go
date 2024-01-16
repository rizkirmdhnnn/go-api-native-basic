package models

import "time"

type Admin struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(100)" json:"name"`
	Username  string    `gorm:"type:varchar(100)" json:"username"`
	Password  string    `gorm:"type:varchar(150)" json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AdminResponse struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"type:varchar(100)" json:"name"`
	Username string `gorm:"type:varchar(100)" json:"username"`
}
