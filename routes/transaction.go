package routes

import (
	"github.com/gorilla/mux"
	"go-api-native-basic/controllers/transactioncontroller"
)

func TransactionRouter(r *mux.Router) {
	router := r.PathPrefix("/transaction").Subrouter()
	router.HandleFunc("", transactioncontroller.Create).Methods("POST")
	router.HandleFunc("/{id}/return", transactioncontroller.Update).Methods("PUT")
}
