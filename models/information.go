package models

type Info struct {
	TotalBooks          int64 `json:"total_books"`
	TotalMember         int64 `json:"total_member"`
	TotalBorrowingToday int64 `json:"total_borrowing_today"`
	TotalReturnToday    int64 `json:"total_return_today"`
	TotalTransaction    int64 `json:"total_transaction"`
}
