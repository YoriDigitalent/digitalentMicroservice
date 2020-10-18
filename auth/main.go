package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/YoriDigitalent/digitalentMicroservice/auth/handler"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.Handle("/admin-auth", http.HandlerFunc(handler.ValidateAuth))

	fmt.Printf("Auth service listen on :8001")
	log.Panic(http.ListenAndServe(":8001", router))
}
