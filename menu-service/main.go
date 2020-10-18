package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/YoriDigitalent/digitalentMicroservice/menu-service/config"
	"github.com/YoriDigitalent/digitalentMicroservice/menu-service/handler"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func main() {

	cfg, err := getConfig()
	if err != nil {
		log.Panic(err)
		return
	}

	router := mux.NewRouter()

	router.Handle("/add-product", http.HandlerFunc(handler.AddMenu))

	fmt.Printf("Server listen on :%s", cfg.Port)
	log.Panic(http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), router))
}

func getConfig() (config.Config, error) {

	viper.AddConfigPath(".")
	viper.SetConfigType("yml")
	viper.SetConfigName("config.yml")

	if err := viper.ReadInConfig(); err != nil {
		return config.Config{}, err
	}

	var cfg config.Config
	err := viper.Unmarshal(&cfg)
	if err != nil {
		return config.Config{}, err
	}

	return cfg, nil
}
