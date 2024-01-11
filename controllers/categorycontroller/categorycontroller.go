package categorycontroller

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
)

func Index(w http.ResponseWriter, r *http.Request) {
	var category []models.Category
	if err := config.DB.Find(&category).Error; err != nil {
		helper.Response(w, 500, "Error", nil)
	}
	helper.Response(w, 200, "List Category", category)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}
	defer r.Body.Close()

	if err := config.DB.Create(&category).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "Success create new category "+category.Description, nil)
}

func Detail(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var category models.Category

	if err := config.DB.First(&category, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Category Not Found", nil)
			return
		}

		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "Detail Category", category)
}

func Update(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var category models.Category

	if err := config.DB.First(&category, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Category Not Found", nil)
			return
		}
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	if err := config.DB.Where("id = ?", id).Updates(&category).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "Success Update Category", nil)

}

func Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(params)
	var category models.Category
	res := config.DB.Delete(&category, id)

	if res.Error != nil {
		helper.Response(w, 500, res.Error.Error(), nil)
		return
	}

	if res.RowsAffected == 0 {
		helper.Response(w, 404, "Category Not Found", nil)
		return
	}

	helper.Response(w, 200, "Success Delete Category "+category.Description, nil)

}
