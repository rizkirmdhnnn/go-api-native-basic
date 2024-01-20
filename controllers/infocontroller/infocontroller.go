package infocontroller

import (
	"go-api-native-basic/config"
	"go-api-native-basic/helper"
	"go-api-native-basic/models"
	"net/http"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var info models.Info

	var books []models.Book
	if err := config.DB.Model(&books).Count(&info.TotalBooks).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	var member []models.Member
	if err := config.DB.Model(&member).Count(&info.TotalMember).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	var transactions models.Transactions
	if err := config.DB.Model(&transactions).Count(&info.TotalTransaction).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	today := time.Now().Format("02-01-2006")
	if err := config.DB.
		Model(&transactions).
		Where("borrowing_date = ? AND return_date = ?", today, "").
		Count(&info.TotalBorrowingToday).
		Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	if err := config.DB.
		Model(&transactions).
		Where("return_date = ? ", today).
		Count(&info.TotalReturnToday).
		Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}
	helper.Response(w, 200, "info", info)

}
