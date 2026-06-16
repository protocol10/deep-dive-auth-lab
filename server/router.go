package server

import (
	"github/com/protocol10/deep-dive-auth-lab/auth"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRouters(authHandler *auth.Handler) {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/register", authHandler.RegisterHandler).Methods("POST")
	r.HandleFunc("/api/v1/login", authHandler.LoginHandler).Methods("POST")

	// Create a subrouter for protected routes that require authentication
	protected := r.PathPrefix("/api/v1/protected").Subrouter()
	protected.Use(auth.AuthMiddleware)
	protected.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"message": "Welcome to your protected profile!"}`))
	}).Methods("GET")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
