package handlers

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"myapp-backend/internal/database"
	"myapp-backend/internal/models"
)

func SignUpHandler(w http.ResponseWriter, r * http.Request) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil{
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password),bcrypt.DefaultCost)
	if err !=nil{
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	user := models.User{
		Username: input.Username,
		Password: string(hashedPassword),
	}

	if err := database.DB.Create(&user).Error; err != nil{
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Cotent-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":"signup success",
		"user":map[string]interface{}{
			"id":user.ID,
			"username":user.Username,
		},
	})

}
