package main

import (
	"log"

	"github.com/spf13/viper"
)

func main() {
	//From current directory, read config.yml file (ignore yml extension)
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error in reading config", err)
	}

	//Get a value from config.yml
	myport := viper.Get("server.port")
	log.Println("Port number is :", myport)

	//Get a value from Environment variables
	viper.AutomaticEnv()
	mypath := viper.Get("PATH")
	log.Println("PATH is :", mypath)

}
