package authcontroller

import (
	"encoding/json"
	"errors"
	"go-api-native-basic/config"
	"go-api-native-basic/helper"
	"go-api-native-basic/models"
	"gorm.io/gorm"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	//Create var
	var userInput models.Admin

	//Decode data from request body
	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}
	//Close connection request body
	defer r.Body.Close()

	var user models.Admin
	//Find username in db
	if err := config.DB.Where("username = ?", userInput.Username).First(&user).Error; err != nil {
		switch {
		//If username not found
		case errors.Is(err, gorm.ErrRecordNotFound):
			helper.Response(w, 400, "username not found", nil)
			return
		default:
			helper.Response(w, 500, err.Error(), nil)
			return
		}
	}

	//Checking password hash
	if err := helper.CheckPasswordHash(userInput.Password, user.Password); err != true {
		helper.Response(w, 400, "Wrong password", nil)
		return
	}

	//Create access token JWT
	accessToken, err := helper.CreateAccessToken(&user)
	if err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	//Create refresh token JWT
	refreshToken, err := helper.CreateRefreshToken(&user)
	if err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	//Set access token to cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "Authorization",
		Path:     "/",
		Value:    accessToken,
		HttpOnly: true,
	})

	//Set refresh token to cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "Refresh",
		Path:     "/",
		Value:    refreshToken,
		HttpOnly: true,
	})
	helper.Response(w, 200, "Success Login", nil)

}
func Logout(w http.ResponseWriter, r *http.Request) {

	//Remove access token cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "Authorization",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})

	//Remove refresh token cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "Refresh",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})

	helper.Response(w, 200, "You have successfully logged out", nil)
}
