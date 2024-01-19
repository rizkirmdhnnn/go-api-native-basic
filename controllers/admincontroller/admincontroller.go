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
	//Create array
	var admin []models.Admin

	//Find data
	if err := config.DB.Find(&admin).Error; err != nil {
		helper.Response(w, 500, "Error", nil)
	}
	helper.Response(w, 200, "List Admins", admin)
}

func Create(w http.ResponseWriter, r *http.Request) {
	//Create var
	var admin models.Admin

	//Decode data from body
	if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	//Close connection
	defer r.Body.Close()

	//Search for the same username
	if err := config.DB.Where("username = ?", admin.Username).First(&admin); err != nil {
		helper.Response(w, 400, "Username already exists", nil)
		return
	}

	//hash password using bcrypt
	hashedPassword, err := helper.HashPassword(admin.Password)
	if err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}
	admin.Password = hashedPassword

	//Create data
	if err := config.DB.Create(&admin).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "Success create admin", nil)
}

func Detail(w http.ResponseWriter, r *http.Request) {
	//Get id from params
	idParams := mux.Vars(r)["id"]
	//Convert string to int
	id, _ := strconv.Atoi(idParams)

	//Create var
	var admin models.Admin

	//Search for data
	if err := config.DB.First(&admin, id).Error; err != nil {
		//If data record not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Admin Not Found", nil)
			return
		}
		return
	}
	helper.Response(w, 200, "Detail Admin", admin)
}

func Update(w http.ResponseWriter, r *http.Request) {
	//Get id from params
	idParams := mux.Vars(r)["id"]
	//Convert string to int
	id, _ := strconv.Atoi(idParams)

	//Create var
	var admin models.Admin

	//Search for data
	if err := config.DB.First(&admin, id).Error; err != nil {
		//If data record not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Admin Not Found", nil)
			return
		}
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	//Get Query from url and assign to oldPassword
	oldPassword := r.URL.Query().Get("oldPassword")
	//Checking password
	if helper.CheckPasswordHash(oldPassword, admin.Password) != true {
		helper.Response(w, 400, "Old Password Not Match", nil)
		return
	}

	//Decode json from request body
	if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}
	//Close request body
	defer r.Body.Close()

	//Search if there is the same username
	if err := config.DB.Where("username = ? AND id != ?", admin.Username, id).First(&models.Admin{}).Error; err == nil {
		helper.Response(w, 400, "Username already exists", nil)
		return
	}

	//Hashing new password
	hashedPassword, err := helper.HashPassword(admin.Password)
	if err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}
	admin.Password = hashedPassword

	//Update to db
	if err := config.DB.Where("id = ?", id).Updates(&admin).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "Success update admin", nil)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	//Get id from params
	idParams := mux.Vars(r)["id"]
	//Convert string to int
	id, _ := strconv.Atoi(idParams)

	//Create var
	var admin models.Admin

	//Find data by id
	if err := config.DB.First(&admin, id).Error; err != nil {
		helper.Response(w, 404, "Admin Not Found", nil)
		return
	}

	//Get Query from url and assign to oldPassword
	oldPassword := r.URL.Query().Get("password")
	if helper.CheckPasswordHash(oldPassword, admin.Password) != true {
		helper.Response(w, 400, "Old Password Not Match", nil)
		return
	}

	//Delete data
	res := config.DB.Delete(&admin, id)

	//If respond err
	if res.Error != nil {
		helper.Response(w, 500, res.Error.Error(), nil)
		return
	}

	helper.Response(w, 200, "Success Delete Admin", nil)

}
