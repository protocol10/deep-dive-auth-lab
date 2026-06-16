package auth

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
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

func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Implement login logic here
	var req models.UserLoginRequest
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

	// Check if the user is already logged in
	cookie, err := r.Cookie("auth_token")
	fmt.Println("Cookie is", cookie)
	if err == nil && cookie.Value != "" {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(models.UserResponse{
			Message: "User already logged in",
			Success: true,
		})
		return
	}

	if err := h.svc.LoginUser(req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.UserResponse{
			Message: err.Error(),
			Success: false,
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	sessionId, err := generateSecureSessionID()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.UserResponse{
			Message: "Failed to generate session ID",
			Success: false,
		})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    sessionId,
		Path:     "/",
		MaxAge:   3600, // 1 hour
		Secure:   true, // Set to true in production with HTTPS
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.UserResponse{
		Message: "User logged in successfully",
		Success: true,
	})
}

func generateSecureSessionID() (string, error) {
	b := make([]byte, 32) // 32 bytes = 256 bits → very secure
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}
