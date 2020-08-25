package main

import (
	"net/http"
	"strconv"
	"calculator/service"
)

func sum(w http.ResponseWriter, req *http.Request) {
	var values = service.GetValues(req)
	var result = values.Number1 + values.Number2
	var responseString = strconv.Itoa(result)
	// fmt.Fprintf(w, "hello\n")
	w.Write([]byte(responseString))
}

func sub(w http.ResponseWriter, req *http.Request) {
	var values = service.GetValues(req)
	var result = values.Number1 - values.Number2
	var responseString = strconv.Itoa(result)
	// fmt.Fprintf(w, "hello\n")
	w.Write([]byte(responseString))
}

func main() {

    http.HandleFunc("/sum", sum)
	http.HandleFunc("/sub", sub)

    http.ListenAndServe(":8090", nil)
}