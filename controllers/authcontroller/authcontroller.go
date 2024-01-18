package authcontroller

import (
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"go-api-native-basic/config"
	"go-api-native-basic/helper"
	"go-api-native-basic/models"
	"gorm.io/gorm"
	"net/http"
	"time"
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

	//Create var expired token
	expTime := time.Now().Add(time.Minute * 60)

	//Create JwtToken
	claims := config.JWTClaim{
		ID:       user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "library-book",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	//using algorithm hs256
	tokenAlogitm := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenAlogitm.SignedString(config.JWT_KEY)
	if err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	//Set jwt token to cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "Token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})

	helper.Response(w, 200, "Success Login", nil)

}
func Logout(w http.ResponseWriter, r *http.Request) {
	//Remove cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "Token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})

	helper.Response(w, 200, "You have successfully logged out", nil)
}
