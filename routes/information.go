package routes

import (
	"github.com/gorilla/mux"
	"go-api-native-basic/controllers/infocontroller"
)

func InformationRouter(r *mux.Router) {
	router := r.PathPrefix("/information").Subrouter()
	router.HandleFunc("", infocontroller.Index).Methods("GET")
}
