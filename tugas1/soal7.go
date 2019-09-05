package main

import (
	"fmt"
	"net/http"
	"github.com/spf13/viper"
	"log"
)

type ServerConfiguration struct {
	Port string
}

type Configuration struct {
	Server ServerConfiguration
}

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	
	var configuration Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, you've requested: %s\n", r.URL.Path)
	})
	port := configuration.Server.Port
	http.ListenAndServe(":" + port, nil)
}