package routes

import (
	"github.com/gorilla/mux"
	"go-api-native-basic/controllers/authcontroller"
)

func AuthRouter(r *mux.Router) {
	router := r.PathPrefix("/auth").Subrouter()
	router.HandleFunc("/login", authcontroller.Login).Methods("POST")
	router.HandleFunc("/logout", authcontroller.Logout).Methods("GET")
}
