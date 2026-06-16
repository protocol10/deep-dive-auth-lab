package main

import (
	"context"
	"fmt"
	"github/com/protocol10/deep-dive-auth-lab/auth"
	"github/com/protocol10/deep-dive-auth-lab/auth/repository"
	"github/com/protocol10/deep-dive-auth-lab/auth/service"
	"github/com/protocol10/deep-dive-auth-lab/server"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	connectionStr := "postgres://akshaymukadam@localhost:5432/auth_lab?sslmode=disable"

	ctx := context.Background()
	pgxPool, err := pgxpool.New(ctx, connectionStr)
	if err != nil {
		panic(err)
	}
	defer pgxPool.Close()

	if err := pgxPool.Ping(ctx); err != nil {
		panic(err)
	}
	fmt.Println("Database connected successfully")
	repo := repository.NewRepository(pgxPool, ctx)
	authService := service.NewAuthService(repo)
	authHandler := auth.NewHandler(authService)
	server.InitializeRouters(authHandler)
}
