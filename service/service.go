package service

import (
    "encoding/json"
    "net/http"
	"strconv"
)

type input struct {
	Number1 int
	Number2 int
}

func Sum(w http.ResponseWriter, req *http.Request) {
	var values = getValues(req)
	var result = values.Number1 + values.Number2
	var responseString = strconv.Itoa(result)
	// fmt.Fprintf(w, "hello\n")
	w.Write([]byte(responseString))
}

func Sub(w http.ResponseWriter, req *http.Request) {
	var values = getValues(req)
	var result = values.Number1 - values.Number2
	var responseString = strconv.Itoa(result)
	// fmt.Fprintf(w, "hello\n")
	w.Write([]byte(responseString))
}

func getValues(req *http.Request) input{
	decoder := json.NewDecoder(req.Body)
    var t input
    err := decoder.Decode(&t)
    if err != nil {
        panic(err)
    }
    return t
}