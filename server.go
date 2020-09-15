package main

import (
	// "net/http"
	"calculator/database"
	"calculator/service"
	"calculator/logger"
	"calculator/config"
	// "github.com/gorilla/mux"
	// "os"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// log "github.com/sirupsen/logrus"
)

// Used to set routes

func setupRoutes(app *fiber.App) {
	// Used with fiber
	app.Post("/sum", service.Sum)
	app.Post("/sub", service.Sub)
	app.Post("/div", service.Div)
	app.Get("/history", service.History)
	app.Get("/error", service.Exception)
	// logger.InfoLogger.Println("Routes Generated")
	// log.Info("Routes Generated")
	logger.Logger.Info("Routes Generated")	
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("postgres", "host="+config.Cfg.Host+" port="+config.Cfg.Port+" user="+config.Cfg.User+" dbname="+config.Cfg.Dbname+" password="+config.Cfg.Password+" sslmode=disable")
	if err != nil {
		panic("Failed to connect to database. Reason: "+ err.Error())
	}
	database.DBConn.AutoMigrate(&service.Input{})
	// logger.InfoLogger.Println("Connection with db established")
	// log.Info("Connection with db establlished")
	logger.Logger.Info("Connection with db established")
}

func main() {
	defer database.DBConn.Close()

	config.ReadPropertiesViper()
	// config.Read()
	// config.ReadProperties()
	// file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    // if err != nil {
    //     log.Fatal(err)
	// }

	// defer file.Close()

	// log.SetFormatter(&log.JSONFormatter{})
	// log.SetOutput(file)
	logger.InitializeLogger()
	// Using fiber

	// service.FraudService()
	// service.BankClient()

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
