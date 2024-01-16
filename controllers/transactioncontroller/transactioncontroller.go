package transactioncontroller

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"go-api-native-basic/config"
	"go-api-native-basic/helper"
	"go-api-native-basic/models"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var transactions []models.Transactions
	var transactionResponse []models.TransactionsResponse

	if err := config.DB.
		Joins("Member").
		Joins("Book.Author").
		Joins("Book.Category").
		Joins("Admin").
		Find(&transactions).
		Find(&transactionResponse).
		Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "Success get all transactions", transactionResponse)
}

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

func Update(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	condition := r.URL.Query().Get("condition")
	if condition == "" {
		helper.Response(w, 400, "Missing parameter condition", nil)
		return
	}

	var transaction models.Transactions

	if err := config.DB.First(&transaction, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Transaction with that id are not found", nil)
			return
		}
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	if err := config.DB.
		Where("member_id = ? AND book_id = ? AND return_date != ?", transaction.MemberID, transaction.BookID, "").
		First(&transaction).
		Error; err == nil {
		helper.Response(w, 409, "Book has been returned!", nil)
		return
	}

	transaction.ReturnDate = time.Now().Format("02-01-2006")

	borrowingDate, err := time.Parse("02-01-2006", transaction.BorrowingDate)
	if err != nil {
		helper.Response(w, 500, "Cant parsing borrowingDate", nil)
		return
	}

	returnDate, err := time.Parse("02-01-2006", transaction.ReturnDate)
	if err != nil {
		helper.Response(w, 500, "Cant parsing borrowingDate", nil)
		return
	}

	totalDays := int32(returnDate.Sub(borrowingDate).Hours() / 24)
	lateDay := totalDays - config.ENV.MAXLOANDURATION

	if lateDay < 0 {
		lateDay = 0
	}

	if lateDay >= config.ENV.MAXLOANDURATION {
		transaction.Penalties += (lateDay - config.ENV.MAXLOANDURATION) * config.ENV.PENALTYPERDAY
	}

	transaction.Condition = condition
	switch condition {
	case "broken":
		transaction.Penalties += config.ENV.PENALTYBROKEN
	case "missing":
		transaction.Penalties += config.ENV.PINALTYLOST
	default:
	}

	if err := config.DB.Where("id = ?", transaction.ID).Updates(&transaction).Error; err != nil {
		helper.Response(w, 404, err.Error(), nil)
		return
	}

	var returnResponse models.TransactionReturnResponse

	returnResponse.LateDay = lateDay
	returnResponse.Penalty = transaction.Penalties

	helper.Response(w, 201, "Successfully returned the book", returnResponse)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var transaction models.Transactions
	response := config.DB.Delete(&transaction, id)

	if response.Error != nil {
		helper.Response(w, 500, response.Error.Error(), nil)
		return
	}

	if response.RowsAffected == 0 {
		helper.Response(w, 404, "Transaction Not Found", nil)
		return
	}

	helper.Response(w, 200, "Success Delete Member", nil)

}
