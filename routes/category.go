package routes

import (
	"github.com/gorilla/mux"
	"go-api-native-basic/controllers/categorycontroller"
)

func CategoryRoute(r *mux.Router) {
	router := r.PathPrefix("/category").Subrouter()
	router.HandleFunc("", categorycontroller.Index).Methods("GET")
	router.HandleFunc("", categorycontroller.Create).Methods("POST")
	router.HandleFunc("/{id}/detail", categorycontroller.Detail).Methods("GET")
	router.HandleFunc("/{id}/update", categorycontroller.Update).Methods("PUT")
	router.HandleFunc("/{id}/delete", categorycontroller.Delete).Methods("DELETE")
}
