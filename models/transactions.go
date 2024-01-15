package models

import "time"

type Transactions struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	MemberID      uint      `json:"member_id"`
	Member        Member    `gorm:"foreignKey:MemberID" json:"member"`
	BookID        uint      `json:"book_id"`
	Book          Book      `gorm:"foreignKey:BookID" json:"book"`
	AdminID       uint      `json:"admin_id"`
	Admin         Admin     `gorm:"foreignKey:AdminID" json:"admin"`
	BorrowingDate string    `json:"borrowing_date"`
	ReturnDate    string    `json:"return_date"`
	Penalties     int32     `json:"penalties"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type TransactionsResponse struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	MemberID      uint      `json:"-"`
	Member        Member    `gorm:"foreignKey:MemberID" json:"member"`
	BookID        uint      `json:"-"`
	Book          Book      `gorm:"foreignKey:BookID" json:"book"`
	AdminID       uint      `json:"-"`
	Admin         Admin     `gorm:"foreignKey:AdminID" json:"admin"`
	BorrowingDate time.Time `json:"borrowing_date"`
	ReturnDate    time.Time `json:"return_date"`
	Penalties     int32     `json:"penalties"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
