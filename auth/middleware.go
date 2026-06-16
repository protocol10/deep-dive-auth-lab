package auth

import (
	"encoding/json"
	"net/http"

	"github/com/protocol10/deep-dive-auth-lab/auth/models"
)

// AuthMiddleware checks if the user has a valid auth_token cookie.
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("auth_token")
		if err != nil || cookie.Value == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(models.UserResponse{
				Message: "Unauthorized: Missing or invalid token",
				Success: false,
			})
			return
		}

		// If the token exists, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}
