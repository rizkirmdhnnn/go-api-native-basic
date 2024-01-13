package admincontroller

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
	var admin []models.Admin

	if err := config.DB.Find(&admin).Error; err != nil {
		helper.Response(w, 500, "Error", nil)
	}
	helper.Response(w, 200, "List Admins", admin)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var admin models.Admin

	if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	if err := config.DB.Where("username = ?", admin.Username).First(&admin).Error; err == nil {
		helper.Response(w, 400, "Username already exists", nil)
		return
	}

	hashedPassword, err := helper.HashPassword(admin.Password)
	if err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}
	admin.Password = hashedPassword

	if err := config.DB.Create(&admin).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "Success create admin", nil)
}

func Detail(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var admin models.Admin

	if err := config.DB.First(&admin, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Admin Not Found", nil)
			return
		}
		return
	}
	helper.Response(w, 200, "Detail Admin", admin)
}

func Update(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var admin models.Admin

	if err := config.DB.First(&admin, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Admin Not Found", nil)
			return
		}
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	oldPassword := r.URL.Query().Get("oldPassword")
	if helper.CheckPasswordHash(oldPassword, admin.Password) != true {
		helper.Response(w, 400, "Old Password Not Match", nil)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	if err := config.DB.Where("username = ? AND id != ?", admin.Username, id).First(&models.Admin{}).Error; err == nil {
		helper.Response(w, 400, "Username already exists", nil)
		return
	}

	hashedPassword, err := helper.HashPassword(admin.Password)
	if err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}
	admin.Password = hashedPassword

	if err := config.DB.Where("id = ?", id).Updates(&admin).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "Success update admin", nil)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var admin models.Admin

	if err := config.DB.First(&admin, id).Error; err != nil {
		helper.Response(w, 404, "Admin Not Found", nil)
		return
	}

	oldPassword := r.URL.Query().Get("password")
	if helper.CheckPasswordHash(oldPassword, admin.Password) != true {
		helper.Response(w, 400, "Old Password Not Match", nil)
		return
	}

	res := config.DB.Delete(&admin, id)

	if res.Error != nil {
		helper.Response(w, 500, res.Error.Error(), nil)
		return
	}

	helper.Response(w, 200, "Success Delete Admin", nil)

}
