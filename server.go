package main

import (
	// "net/http"
	"calculator/database"
	"calculator/service"

	// "github.com/gorilla/mux"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Used to set routes

func setupRoutes(app *fiber.App) {
	// Used with fiber
	app.Post("/sum", service.Sum)
	app.Post("/sub", service.Sub)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("postgres", "host=localhost port=5432 user=prashantbedi dbname=calculator password=postgres sslmode=disable")
	if err != nil {
		panic("Failed to connect to database. Reason: "+ err.Error())
	}
}

func main() {
	defer database.DBConn.Close()
	
	// Using fiber

	var router = fiber.New()

	initDatabase()

	setupRoutes(router)

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
