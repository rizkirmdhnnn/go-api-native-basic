package routes

import (
	"github.com/gorilla/mux"
	"go-api-native-basic/controllers/authorcontroller"
)

func AuthorRoutes(r *mux.Router) {
	router := r.PathPrefix("/authors").Subrouter()
	router.HandleFunc("", authorcontroller.Index).Methods("GET")
	router.HandleFunc("", authorcontroller.Create).Methods("POST")
	router.HandleFunc("/{id}/detail", authorcontroller.Detail).Methods("GET")
	router.HandleFunc("/{id}/update", authorcontroller.Update).Methods("PUT")
	router.HandleFunc("/{id}/delete", authorcontroller.Delete).Methods("DELETE")
}
