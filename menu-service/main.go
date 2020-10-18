package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/YoriDigitalent/digitalentMicroservice/menu-service/handler"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.Handle("/add-product", http.HandlerFunc(handler.AddMenu))

	fmt.Println("Server listen on :8000")
	log.Panic(http.ListenAndServe(":8000", router))
}
