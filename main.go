package main

import (
	"fmt"
	"log"
	"net/http"
	"vncities-distance/router"
)

func main() {
	r := router.Router()
	// fs := http.FileServer(http.Dir("build"))
	// http.Handle("/", fs)
	fmt.Println("Starting server on the port 10000...")

	log.Fatal(http.ListenAndServe(":10000", r))
}
