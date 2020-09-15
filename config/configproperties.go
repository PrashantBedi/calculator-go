package config

import (
	"github.com/spf13/viper"
	"os"
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
	p := properties.MustLoadFile(os.Args[1], properties.UTF8)
	if err := p.Decode(&Cfg); err != nil {
		panic(err)
	}
	fmt.Println("Using Properties",Cfg)
}


func ReadPropertiesViper() {
	viper.SetConfigName(os.Args[2])
	viper.AddConfigPath(os.Args[1])
	viper.AutomaticEnv()
	viper.SetConfigType("properties")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&Cfg)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
	fmt.Println("Using Viper " , Cfg)
}
