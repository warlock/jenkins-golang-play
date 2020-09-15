package main

import (
	"log"
	"net/http"

	"github.com/warlock/jenkins-golang-play/handlers"
)

func main() {
	http.HandleFunc("/health", handlers.HealthCheckHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
