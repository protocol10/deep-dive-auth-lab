package server

import (
	"github/com/protocol10/deep-dive-auth-lab/auth"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRouters(authHandler *auth.Handler) {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/register", authHandler.RegisterHandler).Methods("POST")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
