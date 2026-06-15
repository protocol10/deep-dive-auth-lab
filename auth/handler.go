package auth

import (
	"encoding/json"
	"fmt"
	"github/com/protocol10/deep-dive-auth-lab/auth/models"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req models.UserRegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("request is ", req)

}
