package auth

import (
	"encoding/json"
	"github/com/protocol10/deep-dive-auth-lab/auth/models"
	"github/com/protocol10/deep-dive-auth-lab/auth/service"
	"net/http"
)

type Handler struct {
	svc service.AuthService
}

func NewHandler(svc service.AuthService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
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

	if err := h.svc.RegisterUser(req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError) // Consider using 400 or 409 depending on the specific error
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
