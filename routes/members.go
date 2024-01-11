package routes

import (
	"github.com/gorilla/mux"
	"go-api-native-basic/controllers/memberscontroller"
)

func MembersRoute(r *mux.Router) {
	router := r.PathPrefix("/member").Subrouter()
	router.HandleFunc("", memberscontroller.Index).Methods("GET")
	router.HandleFunc("", memberscontroller.Create).Methods("POST")
	router.HandleFunc("/{id}/detail", memberscontroller.Detail).Methods("GET")
	router.HandleFunc("/{id}/update", memberscontroller.Update).Methods("PUT")
	router.HandleFunc("/{id}/delete", memberscontroller.Delete).Methods("DELETE")
}
