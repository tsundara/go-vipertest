package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error in reading config", err)
	}

	//Get a value from the config.yml
	myport := viper.Get("server.port")
	fmt.Println("Value:", myport)

	//Get a value from Environment variables
	viper.AutomaticEnv()
	mypath := viper.Get("PATH")
	fmt.Println("gopath:", mypath)

}
