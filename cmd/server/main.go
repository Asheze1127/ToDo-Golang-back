package main 

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"myapp-backend/internal/database"
	"myapp-backend/internal/models"
	"myapp-backend/internal/handlers"
)

func main() {
	database.Connect()
	database.DB.AutoMigrate(&models.User{})

	r := chi.NewRouter()
	r.Post("/signup", handlers.SignUpHandler)
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	fmt.Println("🚀 Server started on :8080")
	http.ListenAndServe(":8080", r)
}
