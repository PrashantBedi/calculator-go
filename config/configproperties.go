package config

import (
	"fmt"
	"github.com/magiconair/properties"
)

type ConfigProperties struct {
	Host     string `properties:"host,default=localhost"`
	Port     string `properties:"port,default=5432"`
	Dbname   string `properties:"dbname"`
	User     string `properties:"user"`
	Password string `properties:"password"`
}

var Cfg = ConfigProperties{}

func Read() {
	p := properties.MustLoadFile("/Users/prashantbedi/workspace-go/calculator/config/config.properties", properties.UTF8)
	if err := p.Decode(&Cfg); err != nil {
		panic(err)
	}
	fmt.Println("Using Properties",Cfg)
}