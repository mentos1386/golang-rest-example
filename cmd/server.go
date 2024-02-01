package main

import (
	"log"
	"net/http"

	"github.com/mentos1386/golang-rest-example/pkg/api"
)

func main() {

	srv, err := api.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal(err)
	}
}
