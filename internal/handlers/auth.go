package handlers

import (
	"encoding/json"
	"myapp-backend/internal/infra/db"
	"myapp-backend/internal/service"
	"net/http"
)

type AuthHandler struct {
	AuthService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	token, err := h.AuthService.SignUp(input.Username, input.Password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]any{
		"message": "signup success",
		"token":   token,
	})
}

// パッケージレベルのSignUp関数
func SignUp(w http.ResponseWriter, r *http.Request) {
	userRepo := db.NewGormUserRepository(db.DB)
	authService := service.NewAuthService(userRepo)
	authHandler := NewAuthHandler(authService)
	authHandler.SignUp(w, r)
}
