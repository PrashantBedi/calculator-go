package service

import (
    "encoding/json"
    // "net/http"
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
	var result = values.Number1 + values.Number2
    c.Send(result)
}

// Using filter

func Sub(c *fiber.Ctx) {
    var values Input
    json.Unmarshal([]byte(c.Body()), &values)
	var result = values.Number1 - values.Number2
    c.Send(result)
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