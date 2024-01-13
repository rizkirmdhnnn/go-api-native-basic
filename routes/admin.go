package routes

import (
	"github.com/gorilla/mux"
	"go-api-native-basic/controllers/admincontroller"
)

func AdminRoutes(r *mux.Router) {
	router := r.PathPrefix("/admin").Subrouter()
	router.HandleFunc("", admincontroller.Index).Methods("GET")
	router.HandleFunc("", admincontroller.Create).Methods("POST")
	router.HandleFunc("/{id}/detail", admincontroller.Detail).Methods("GET")
	router.HandleFunc("/{id}/update", admincontroller.Update).Methods("PUT")
	router.HandleFunc("/{id}/delete", admincontroller.Delete).Methods("DELETE")

}
