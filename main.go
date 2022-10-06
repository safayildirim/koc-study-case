package main

import (
	"fmt"
	"koc-digital-case/controllers"
	"koc-digital-case/models"
	"koc-digital-case/repositories"
	"koc-digital-case/services"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

type Handler interface {
	RegisterRoutes(app *fiber.App)
}

type Server struct {
	app *fiber.App
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	repository := repositories.NewRepository()
	authRepository := repositories.NewUserRepository(repository)
	urlRepository := repositories.NewURLRepository(repository)
	tokenService := services.NewTokenService(os.Getenv("SECRET"))
	authService := services.NewAuthService(authRepository, tokenService)
	urlService := services.NewURLService(urlRepository, authRepository)
	urlHandler := controllers.NewURLHandler(urlService, tokenService)
	authHandler := controllers.NewAuthHandler(authService)
	server := NewServer([]Handler{urlHandler, authHandler})
	log.Fatalln(server.app.Listen(fmt.Sprintf(":%v", port)))
}

func NewServer(handlers []Handler) *Server {
	app := fiber.New(fiber.Config{ErrorHandler: models.ErrorHandler})
	for _, handler := range handlers {
		handler.RegisterRoutes(app)
	}

	return &Server{app: app}
}
