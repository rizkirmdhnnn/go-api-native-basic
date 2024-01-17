package routes

import (
	"github.com/gorilla/mux"
	"go-api-native-basic/controllers/transactioncontroller"
	"go-api-native-basic/middlewares"
)

func TransactionRouter(r *mux.Router) {
	router := r.PathPrefix("/transaction").Subrouter()
	router.HandleFunc("", transactioncontroller.Index).Methods("GET")
	router.HandleFunc("", transactioncontroller.Create).Methods("POST")
	router.HandleFunc("/{id}/return", transactioncontroller.Update).Methods("PUT")
	router.HandleFunc("/{id}/delete", transactioncontroller.Delete).Methods("DELETE")
	router.Use(middlewares.JWTMiddleware)

}
