package controllers

import (
	"koc-digital-case/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService AuthService
}

type AuthService interface {
	SignUp(request *models.SignUpRequest) error
	SignIn(request *models.SignInRequest) (string, error)
}

func NewAuthHandler(authService AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) RegisterRoutes(app *fiber.App) {
	app.Group("/api")
	app.Post("/login", h.signIn)
	app.Post("/signup", h.signUp)
}

func (h *AuthHandler) signIn(ctx *fiber.Ctx) error {
	var request models.SignInRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		return err
	}
	token, err := h.authService.SignIn(&request)
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusOK).JSON(models.Response{
		Data:  token,
		Error: "",
	})
}

func (h *AuthHandler) signUp(ctx *fiber.Ctx) error {
	var request models.SignUpRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		return err
	}
	err = h.authService.SignUp(&request)
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusOK).JSON(models.Response{
		Data:  "",
		Error: "",
	})
}
