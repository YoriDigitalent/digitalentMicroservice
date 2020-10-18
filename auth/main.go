package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/YoriDigitalent/digitalentMicroservice/auth/config"
	"github.com/YoriDigitalent/digitalentMicroservice/auth/database"
	"github.com/YoriDigitalent/digitalentMicroservice/auth/handler"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	cfg, err := getConfig()
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println(cfg)
	}

	db, err := initDB(cfg.Database)
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println("DB Connection Success")
	}

	authHandler := handler.AuthDB{Db: db}

	router := mux.NewRouter()

	router.Handle("/auth/validate", http.HandlerFunc(authHandler.ValidateAdmin))
	router.Handle("/auth/sign-up", http.HandlerFunc(authHandler.SignUp))
	router.Handle("/auth/login", http.HandlerFunc(authHandler.Login))

	fmt.Printf("Auth service listen on port 8001")
	log.Panic(http.ListenAndServe(":8001", router))
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

func initDB(cfg config.Database) (*gorm.DB, error) {
	//root:password@tcp(localhost:3306)/dts_microservice_auth?cofig
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DbName, cfg.Config)
	log.Println(dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&database.Auth{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
