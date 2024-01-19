package middlewares

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"go-api-native-basic/helper"
	"net/http"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Read data cookie Authorization
		cAuthorization, err := r.Cookie("Authorization")
		if err != nil {
			//If cookie Authorization not found
			if errors.Is(err, http.ErrNoCookie) {
				helper.Response(w, 406, "Unauthorized, please login ", nil)
				return
			}
		}
		//Extract data from cookie Authorization
		accessToken := cAuthorization.Value
		//Validate token
		admin, err := helper.ValidateToken(accessToken)
		if err != nil {
			//If token expired
			if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
				//Read data cookie Authorization
				cRefresh, err := r.Cookie("Refresh")
				if err != nil {
					//If cookie Authorization not found
					if errors.Is(err, http.ErrNoCookie) {
						helper.Response(w, 406, "Unauthorized, please login", nil)
						return
					}
				}
				//Refresh Token
				admin, err = helper.HandleTokenRefresh(w, cRefresh.Value)
				if err != nil {
					helper.Response(w, 500, "Refresh token expired, please login again", nil)
					return
				}
			} else {
				logrus.WithError(err).Error("Error validating access token")
				helper.Response(w, 500, "Failed to validate access token", nil)
				return
			}
		}
		ctx := context.WithValue(r.Context(), "adminInfo", admin)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
