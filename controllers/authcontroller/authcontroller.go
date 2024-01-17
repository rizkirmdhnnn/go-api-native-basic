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
	var userInput models.Admin
	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}
	defer r.Body.Close()

	var user models.Admin
	if err := config.DB.Where("username = ?", userInput.Username).First(&user).Error; err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			helper.Response(w, 400, "username not found", nil)
			return
		default:
			helper.Response(w, 500, err.Error(), nil)
			return
		}
	}

	if err := helper.CheckPasswordHash(userInput.Password, user.Password); err != true {
		helper.Response(w, 400, "Wrong password", nil)
		return
	}

	expTime := time.Now().Add(time.Minute * 60)
	claims := config.JWTClaim{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "library-book",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	tokenAlogitm := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenAlogitm.SignedString(config.JWT_KEY)
	if err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "Token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})

	helper.Response(w, 200, "Success Login", nil)

}
func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "Token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})

	helper.Response(w, 200, "You have successfully logged out", nil)
}
