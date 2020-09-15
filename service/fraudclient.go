package service

import (
	"fmt"
	"bytes"
	"net/http"
)

func FraudService() {
	requestJSON := []byte(`{
		"amount": 1000,
		"beneficiaryName": "user1",
		"beneficiaryAccountNumber": 12345,
		"beneficiaryIfscCode": "HDFC1234",
		"payeeName": "user2",
		"payeeAccountNumber": 67890,
		"payeeIfscCode": "HDFC1234"
	}`)

	res, err := http.Post("http://host.docker.internal:8083/checkFraud","application/json",bytes.NewBuffer(requestJSON))
	if err!= nil {
		panic(err)
	}

	fmt.Println("Fraud Service ---> ",res.Status)
}