package routes

import (
	"github.com/gorilla/mux"
	"go-api-native-basic/controllers/bookcontroller"
)

func BooksRoute(r *mux.Router) {
	router := r.PathPrefix("/books").Subrouter()
	router.HandleFunc("", bookcontroller.Index).Methods("GET")
	router.HandleFunc("", bookcontroller.Create).Methods("POST")
	router.HandleFunc("/{id}/detail", bookcontroller.Detail).Methods("GET")
	router.HandleFunc("/{id}/update", bookcontroller.Update).Methods("PUT")
	router.HandleFunc("/{id}/delete", bookcontroller.Delete).Methods("DELETE")
}
