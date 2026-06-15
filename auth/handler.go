package auth

import (
	"encoding/json"
	"github/com/protocol10/deep-dive-auth-lab/auth/models"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req models.UserRegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.UserResponse{
			Message: "Invalid request body",
			Success: false,
		})
		return
	}

	if err := req.Validate(); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.UserResponse{
			Message: err.Error(),
			Success: false,
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.UserResponse{
		Message: "User registered successfully",
		Success: true,
	})
}
