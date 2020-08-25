package main

import (
	// "net/http"
	"calculator/service" 
	// "github.com/gorilla/mux"
	"github.com/gofiber/fiber"
)

func main() {

	// Using fiber

	var router = fiber.New()

	router.Post("/sum", service.Sum)
	router.Post("/sub", service.Sub)

	
	// Using mux
	
	// var router = mux.NewRouter()
	// router.HandleFunc("/sum",service.Sum).Methods(http.MethodPost)
	// router.HandleFunc("/sub",service.Sub).Methods(http.MethodPost)

	// Using default library

    // http.HandleFunc("/sum", service.Sum)
	// http.HandleFunc("/sub", service.Sub)

	// Used with default and mux library

	// http.ListenAndServe(":8090", router)

	// Used with fiber library
	
	router.Listen(8080)
}