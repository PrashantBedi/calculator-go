package main

import (
	"net/http"
	"calculator/service"
)

func main() {

    http.HandleFunc("/sum", service.Sum)
	http.HandleFunc("/sub", service.Sub)

    http.ListenAndServe(":8090", nil)
}