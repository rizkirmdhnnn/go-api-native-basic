package models

import "time"

type Transactions struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	MemberID  uint      `json:"-"`
	Member    Member    `gorm:"foreignKey:MemberID" json:"member"`
	BookID    uint      `json:"-"`
	Book      Book      `gorm:"foreignKey:BookID" json:"book"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
