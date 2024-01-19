package helper

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"go-api-native-basic/models"
	"net/http"
	"time"
)

var mySigningKey = []byte("thisIsSecret")

type MyCustomClaims struct {
	AdminID  uint   `json:"admin_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func CreateAccessToken(admin *models.Admin) (string, error) {
	Claims := MyCustomClaims{
		admin.ID,
		admin.Username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)
	return token.SignedString(mySigningKey)

}

func CreateRefreshToken(admin *models.Admin) (string, error) {
	claims := MyCustomClaims{
		admin.ID,
		admin.Username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour * 7)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)
}

func ValidateToken(tokenString string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}
	claims, _ := token.Claims.(*MyCustomClaims)
	return claims, nil
}

func RefreshToken(refreshTokenString string) (string, string, error) {
	// Parse refresh token
	token, err := jwt.ParseWithClaims(refreshTokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	//Check if error
	if err != nil || !token.Valid {
		return "", "", err
	}

	//Claims token
	claims, _ := token.Claims.(*MyCustomClaims)
	//Create new access token
	newAccessToken, err := CreateAccessToken(&models.Admin{ID: claims.AdminID, Username: claims.Username})
	if err != nil {
		return "", "", err
	}

	//Create new refresh token
	newRefreshToken, err := CreateRefreshToken(&models.Admin{ID: claims.AdminID, Username: claims.Username})
	if err != nil {
		return "", "", err
	}

	return newAccessToken, newRefreshToken, nil
}

func HandleTokenRefresh(w http.ResponseWriter, refreshToken string) (*MyCustomClaims, error) {
	//refresh token
	newAccessToken, newRefreshToken, err := RefreshToken(refreshToken)

	if err != nil {
		logrus.WithError(err).Error("Refresh Token Expired")
		return nil, err
	}

	logrus.Println("newRefreshToken : ", newRefreshToken)
	logrus.Println("newAccessToken : ", newAccessToken)

	http.SetCookie(w, &http.Cookie{
		Name:     "Authorization",
		Path:     "/",
		Value:    newAccessToken,
		HttpOnly: true,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "Refresh",
		Path:     "/",
		Value:    newRefreshToken,
		HttpOnly: true,
	})
	return ValidateToken(newAccessToken)
}
