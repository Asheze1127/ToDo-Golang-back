package handlers

import (
	"encoding/json"
	"net/http"
	"myapp-backend/internal/service"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(authService *service.AuthService)*AuthHandler{
	return &AuthHandler{AuthService: authService}
}

func (h *AuthHandler) SighUp(w http.ResponseWriter, r *http.Request){
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err!=nil{
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := h.AuthService.SignUp(input.Username, input.Password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]any{
		"message":"signup success",
		"user":map[string]any{
			"id":user.ID.String(),
			"username":user.Username,
		}
	})
}
