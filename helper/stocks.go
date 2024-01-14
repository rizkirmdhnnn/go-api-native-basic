package helper

import (
	"errors"
	"github.com/gorilla/mux"
	"go-api-native-basic/config"
	"go-api-native-basic/models"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func HandleStockChange(w http.ResponseWriter, r *http.Request, increase bool) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var book models.Book

	if err := config.DB.First(&book, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			Response(w, 404, "Book Not Found", nil)
			return
		}
		Response(w, 500, err.Error(), nil)
		return
	}

	if increase {
		book.Stocks++
	}
	if !increase {
		if book.Stocks <= 0 {
			Response(w, 400, "Out of stock books", nil)
			return
		} else {
			book.Stocks--
		}
	}

	if err := config.DB.Save(&book).Error; err != nil {
		Response(w, 500, err.Error(), nil)
		return
	}

	Response(w, 201, "Stock updated successfully", nil)
}

func IncreaseStock(w http.ResponseWriter, r *http.Request) {
	HandleStockChange(w, r, true)
}

func DecreaseStock(w http.ResponseWriter, r *http.Request) {
	HandleStockChange(w, r, false)
}
