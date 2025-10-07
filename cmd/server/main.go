package main

import (
	"fmt"
	"log"
	"net/http"

	"myapp-backend/internal/handlers"
	"myapp-backend/internal/infra/db"
	"myapp-backend/internal/models"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	db.Connect()
	if err := db.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}
	log.Println("✅ Migration complete")

	r := chi.NewRouter()
	r.Post("/signup", handlers.SignUp)
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	fmt.Println("🚀 Server started on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("❌ Server failed: %v", err)
	}
}
