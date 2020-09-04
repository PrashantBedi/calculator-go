package service

import (
	"errors"
	"go.uber.org/zap"
	"calculator/database"
	"calculator/logger"
	// "fmt"
    "encoding/json"
	"net/http"
	// "strconv"
	"github.com/gofiber/fiber"
)

// Used when using mux or default library

// func Sum(w http.ResponseWriter, req *http.Request) {
// 	var values = getValues(req)
// 	var result = values.Number1 + values.Number2
// 	var responseString = strconv.Itoa(result)
// 	// fmt.Fprintf(w, "hello\n")
// 	w.Write([]byte(responseString))
// }

// Using filter

func Sum(c *fiber.Ctx) {
	var values Input
	json.Unmarshal([]byte(c.Body()), &values)
    values.Result = values.Number1 + values.Number2
    values.Operation = string("+")
	database.DBConn.Create(&values)
	var valuesJson, err = json.Marshal(values)
	if(err != nil) {
		panic(err.Error())
	}
	logger.Logger.Info(string(valuesJson), zap.Int("response", values.Result))
	c.Send(values.Result)
}

// Using filter

func Sub(c *fiber.Ctx) {
	var values Input
	json.Unmarshal([]byte(c.Body()), &values)
    values.Result = values.Number1 - values.Number2
	values.Operation = string("-")
	database.DBConn.Create(&values)
	var valuesJson, err = json.Marshal(values)
	if(err != nil) {
		panic(err.Error())
	}
	logger.Logger.Info(string(valuesJson), zap.Int("response", values.Result))
	c.Send(values.Result)
}

// Using filter

func History(c *fiber.Ctx) {
    var history []Input
    database.DBConn.Find(&history)
    c.JSON(history)
}

func Exception(c *fiber.Ctx) {
	err := errors.New("Abc")
	c.Status(http.StatusBadRequest)
	c.Send(err)
}

// func Sub(w http.ResponseWriter, req *http.Request) {
// 	var values = getValues(req)
// 	var result = values.Number1 - values.Number2
// 	var responseString = strconv.Itoa(result)
// 	// fmt.Fprintf(w, "hello\n")
// 	w.Write([]byte(responseString))
// }

// Used when using mux or default library

// func getValues(req *http.Request) Input{
// 	decoder := json.NewDecoder(req.Body)
//     var t Input
//     err := decoder.Decode(&t)
//     if err != nil {
//         panic(err)
//     }
//     return t
// }
