package transactioncontroller

import (
	"encoding/json"
	"errors"
	"go-api-native-basic/config"
	"go-api-native-basic/helper"
	"go-api-native-basic/models"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var transaction models.Transactions

	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		helper.Response(w, 400, err.Error(), nil)
		return
	}
	defer r.Body.Close()

	transaction.BorrowingDate = time.Now().Format("02-01-2006")
	transaction.ReturnDate = ""
	transaction.Penalties = 0

	if err := config.DB.
		Where("member_id = ? AND book_id = ? AND return_date = ?", transaction.MemberID, transaction.BookID, "").
		First(&transaction).
		Error; err == nil {
		helper.Response(w, 409, "The member still borrowed the book and has not returned it!", nil)
		return
	}

	var member models.Member
	if err := config.DB.First(&member, transaction.MemberID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 500, "Members with this id do not exist", nil)
			return
		}
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	var books models.Book
	if err := config.DB.First(&books, transaction.BookID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 500, "Books with this id do not exist", nil)
			return
		}
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	var admin models.Admin
	if err := config.DB.First(&admin, transaction.AdminID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 500, "Admins with this id do not exist", nil)
			return
		}
		helper.Response(w, 500, err.Error(), nil)
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
