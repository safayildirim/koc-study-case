package main

import (
	"fmt"
	"koc-digital-case/controllers"
	"koc-digital-case/repositories"
	"koc-digital-case/services"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
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
	authRepository := repositories.NewAuthRepository(repository)
	urlRepository := repositories.NewURLRepository(repository)
	authService := services.NewAuthService(authRepository)
	urlService := services.NewURLService(urlRepository)
	urlHandler := controllers.NewURLHandler(urlService)
	authHandler := controllers.NewAuthHandler(authService)
	server := NewServer([]Handler{urlHandler, authHandler})
	log.Fatalln(server.app.Listen(fmt.Sprintf(":%v", port)))
}

func NewServer(handlers []Handler) *Server {
	app := fiber.New()
	app.Use(jwtware.New(jwtware.Config{
		ContextKey: "token",
		SigningKey: []byte("secret"),
	}))
	for _, handler := range handlers {
		handler.RegisterRoutes(app)
	}

	return &Server{app: app}
}
