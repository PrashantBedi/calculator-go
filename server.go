package main

import (
	"net/http"
	"calculator/service" 
	"github.com/gorilla/mux"
)

func main() {

	var router = mux.NewRouter()

	router.HandleFunc("/sum",service.Sum).Methods(http.MethodPost)
	router.HandleFunc("/sub",service.Sub).Methods(http.MethodPost)

    // http.HandleFunc("/sum", service.Sum)
	// http.HandleFunc("/sub", service.Sub)

    http.ListenAndServe(":8090", nil)
}