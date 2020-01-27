package config

import (
	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	//r.HandleFunc("/dashboard", DashboardPage).Methods("GET")
	//`:wr.HandleFunc("/", AppPage).Methods("GET")

	return r
}
