package main

import (
	"log"
	"net/http"

	"github.com/mgonzo/go-api-venues/route"
)

func main() {
	router := route.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
