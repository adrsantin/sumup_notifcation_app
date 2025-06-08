package main

import (
	"database/sql"
	"log"
	"net/http"
	"sumup/notifications/internal/api"
	"sumup/notifications/internal/business"
	"sumup/notifications/internal/repositories"

	"github.com/go-chi/chi/v5"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	log.Println("Starting server...")
	handler := loadAPIs()

	log.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", handler)

}

func loadAPIs() *chi.Mux {
	r := chi.NewRouter()

	db, err := sql.Open("mysql", "root:password@tcp(go-mysql:3306)/sumup")
	if err != nil {
		log.Fatal(err)
	}

	userRepository := repositories.NewUserRepository(db)

	notificationService := business.NewNotificationService(userRepository)

	healthAPI := api.NewHealthAPI()
	notificationsAPI := api.NewNotificationsAPI(
		notificationService,
	)

	api.Routes(r,
		healthAPI,
		notificationsAPI,
	)

	return r
}
