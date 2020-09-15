package main

import (
	"log"
	"net/http"

	"gitlab.git.girona.dev/josep/jenkins-golang-play/handlers"
)

func main() {
	http.HandleFunc("/health", handlers.HealthCheckHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
