package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/YoriDigitalent/digitalentMicroservice/menu-service/config"
	"github.com/YoriDigitalent/digitalentMicroservice/menu-service/database"
	"github.com/YoriDigitalent/digitalentMicroservice/menu-service/handler"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	cfg, err := getConfig()
	if err != nil {
		log.Panic(err)
		return
	}

	db, err := initDB(cfg.Database)

	router := mux.NewRouter()

	menuHandler := handler.Menus{Db: db}

	authMiddleware := handler.AuthMiddleware{
		AuthService: cfg.authService,
	}

	//router.Handle("/add-menu", http.HandlerFunc(menuHandler.AddMenu))
	router.Handle("/add-menu", authMiddleware.ValidateAuth(http.HandlerFunc(menuHandler.AddMenu)))

	router.Handle("/get-menu", http.HandlerFunc(menuHandler.GetMenu))

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

func initDB(dbConfig config.Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DbName, dbConfig.Config)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&database.Menu{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
