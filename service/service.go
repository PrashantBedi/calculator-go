package service

import (
    "encoding/json"
    "net/http"
)

type input struct {
	Number1 int
	Number2 int
}

func GetValues(req *http.Request) input{
	decoder := json.NewDecoder(req.Body)
    var t input
    err := decoder.Decode(&t)
    if err != nil {
        panic(err)
    }
    return t
}