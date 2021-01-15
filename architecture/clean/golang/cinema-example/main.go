package main

import (
	"log"
	"net/http"

	auth_http "github.com/yeldisbayev/education/architecture/clean/golang/cinema-example/auth/delivery/http"
	auth_mongo_repository "github.com/yeldisbayev/education/architecture/clean/golang/cinema-example/auth/repository/mongo"
	auth_usecase "github.com/yeldisbayev/education/architecture/clean/golang/cinema-example/auth/usecase"
)

func main() {
	userRepo := auth_mongo_repository.NewUserRepository()
	authUseCase := auth_usecase.NewAuthUseCase(userRepo)
	authHandler := auth_http.NewHandler(authUseCase)

	http.HandleFunc("/auth/signup", authHandler.SignUp)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
