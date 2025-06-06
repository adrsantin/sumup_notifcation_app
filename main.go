package main

import (
	"fmt"
	"net/http"
	"sumup/notifications/internal/api"
)

func main() {
	handler := api.NewAPIs()

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", handler)

}
