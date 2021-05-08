package router

import (
	"vncities-distance/middleware"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/city", middleware.GetCityInfo).Queries("cityname", "{cityname}").Methods("GET")
	router.HandleFunc("/api/getdistance", middleware.GetDistance).Methods("POST")

	return router
}
