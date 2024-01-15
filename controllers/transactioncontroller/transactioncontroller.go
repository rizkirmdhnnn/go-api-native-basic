package transactioncontroller

import (
	"encoding/json"
	"go-api-native-basic/config"
	"go-api-native-basic/helper"
	"go-api-native-basic/models"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var transaction models.Transactions

	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}
	defer r.Body.Close()

	if transaction.Status != "Borrowing" && transaction.Status != "Returned" {
		helper.Response(w, 400, "Only accepts status Borrowing or Returned", nil)
		return
	}

	if err := config.DB.
		Where("member_id = ? AND book_id = ? AND status = ?", transaction.MemberID, transaction.BookID, "Borrowing").
		First(&transaction).
		Error; err == nil {
		helper.Response(w, 409, "Transaction with the same member_id, book_id, and status 'Borrowing' already exists", nil)
		return
	}

	var member models.Member
	if err := config.DB.First(&member, transaction.MemberID).Error; err != nil {
		helper.Response(w, 404, err.Error(), nil)
		return
	}

	var books models.Book
	if err := config.DB.First(&books, transaction.BookID).Error; err != nil {
		helper.Response(w, 404, err.Error(), nil)
		return
	}

	var admin models.Admin
	if err := config.DB.First(&admin, transaction.AdminID).Error; err != nil {
		helper.Response(w, 404, err.Error(), nil)
		return
	}

	if err := config.DB.Create(&transaction).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	books.Stocks--
	if err := config.DB.Save(&books).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "Success create transaction", nil)
}
