package main

import (
	"log"
	"net/http"

	"github.com/kelseyhightower/envconfig"
)

func main() {
	//Load configuration variables
	err := envconfig.Process("audiobookthing", &spec)
	if err != nil {
		log.Fatal(err)
	}

	// Start HTTP server
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))

}
