package service

import (
	"fmt"
	"net/http"
)

func BankClient() {
	res, err := http.Get("http://host.docker.internal:8082/checkDetails?accountNumber=67890&ifscCode=AXIS1234")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("BankClient ---> ",res.Status)
}