package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"myapp-backend/internal/handlers"
	"myapp-backend/internal/infra/db"
	"myapp-backend/internal/middleware"
	"myapp-backend/internal/models"
)

func main() {
	_ = godotenv.Load()

	db.Connect()
	if err := db.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}
	log.Println("✅ Migration complete")

	r := chi.NewRouter()
	r.Use(middleware.CORS)
	r.Post("/signup", handlers.SignUp)
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	r.Post("/signin", handlers.SignIn)

	fmt.Println("🚀 Server started on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("❌ Server failed: %v", err)
	}
}
