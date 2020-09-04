package config

import (
	"fmt"

	"github.com/tkanos/gonfig"
)

type Properties struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Dbname   string `json:"dbname"`
	User     string `json:"user"`
	Password string `json:"password"`
}

var Configuration = Properties{}

func ReadProperties() {
	err := gonfig.GetConf("/Users/prashantbedi/workspace-go/calculator/config/properties.json", &Configuration)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Using GoConfig",Configuration)
}
