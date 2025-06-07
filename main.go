package main

import (
	"log"
	"net/http"
	"sumup/notifications/internal/api"
)

func main() {

	log.Println("Starting server...")
	handler := api.NewAPIs()

	log.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", handler)

}
