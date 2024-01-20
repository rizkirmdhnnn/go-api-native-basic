package models

import "time"

type Book struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `json:"title"`
	AuthorID    uint      `json:"author_id"`
	Author      Author    `gorm:"foreignKey:AuthorID" json:"author"`
	CategoryID  uint      `json:"category_id"`
	Category    Category  `gorm:"foreignKey:CategoryID" json:"category"`
	Description string    `json:"description"`
	Stocks      int       `json:"stocks"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type BookResponse struct {
	ID          uint                 `json:"id"`
	Title       string               `json:"title"`
	AuthorID    uint                 `json:"-"`
	Author      AuthorBookResponse   `gorm:"foreignKey:AuthorID" json:"author"`
	CategoryID  uint                 `json:"-"`
	Category    CategoryBookResponse `gorm:"foreignKey:CategoryID" json:"category"`
	Description string               `json:"description"`
	Stocks      int                  `json:"stocks"`
	CreatedAt   time.Time            `json:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at"`
}

type BookResponseTransaction struct {
	ID          uint                 `json:"id"`
	Title       string               `json:"title"`
	AuthorID    uint                 `json:"-"`
	Author      AuthorBookResponse   `gorm:"foreignKey:AuthorID" json:"author"`
	CategoryID  uint                 `json:"-"`
	Category    CategoryBookResponse `gorm:"foreignKey:CategoryID" json:"category"`
	Description string               `json:"description"`
	Stocks      int                  `json:"stocks"`
}

type TotalBooks struct {
	TotalBooks int64 `json:"total_books"`
}
