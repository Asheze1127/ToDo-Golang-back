package main 

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"myapp-backend/internal/database"
	"myapp-backend/internal/models"
	"myapp-backend/internal/handlers"
)

func main() {
	_ = godotenv.Load()

	database.Connect()
	if err := database.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}
	log.Println("✅ Migration complete")

	r := chi.NewRouter()
	r.Post("/signup", handlers.SignUpHandler)
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	fmt.Println("🚀 Server started on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("❌ Server failed: %v", err)
	}
}
