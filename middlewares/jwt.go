package middlewares

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"go-api-native-basic/config"
	"go-api-native-basic/helper"
	"net/http"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("Token")
		if err != nil {
			if err == http.ErrNoCookie {
				helper.Response(w, 400, "Unauthorised", nil)
				return
			}
		}

		tokenString := c.Value
		claims := &config.JWTClaim{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		})
		switch {
		case token.Valid:
			fmt.Println("You look nice today")
		case errors.Is(err, jwt.ErrTokenMalformed):
			fmt.Println("That's not even a token")
		case errors.Is(err, jwt.ErrTokenSignatureInvalid):
			// Invalid signature
			fmt.Println("Invalid signature")
		case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
			// Token is either expired or not active yet
			fmt.Println("Timing is everything")
		default:
			helper.Response(w, 400, err.Error(), nil)
			return
		}

		if !token.Valid {
			helper.Response(w, 400, "Unauthorized", nil)
			return
		}

		next.ServeHTTP(w, r)
	})
}
