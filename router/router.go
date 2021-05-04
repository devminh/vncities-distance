package router

import (
	"simple-restapi/middleware"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/user/{id}", middleware.GetUser).Methods("GET")
	router.HandleFunc("/api/user", middleware.GetAllUser).Methods("GET")
	router.HandleFunc("/api/newuser", middleware.CreateUser).Methods("POST")
	router.HandleFunc("/api/updateduser", middleware.UpdateUser).Methods("POST")
	router.HandleFunc("/api/deleteuser/{id}", middleware.DeleteUser).Methods("DELETE")

	router.HandleFunc("/api/city", middleware.GetCityInfo).Queries("cityname", "{cityname}").Methods("GET")
	router.HandleFunc("/api/getdistance", middleware.GetDistance).Methods("POST")

	return router
}
