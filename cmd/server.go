package main

import (
	"log"
	"net/http"

	"github.com/mentos1386/golang-rest-example/pkg/api"
	"github.com/mentos1386/golang-rest-example/pkg/openapi"
)

func main() {
	service := &api.ApiService{}

	srv, err := openapi.NewServer(service)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting server on :1234")
	if err := http.ListenAndServe(":1234", srv); err != nil {
		log.Fatal(err)
	}
}
