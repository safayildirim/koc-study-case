package controllers

import (
	"koc-digital-case/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type (
	AuthHandler struct {
		authService AuthService
	}

	AuthService interface {
		SignUp(request *models.SignUpRequest) error
		SignIn(request *models.SignInRequest) (string, error)
	}
)

func NewAuthHandler(authService AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) RegisterRoutes(app *fiber.App) {
	app.Post("/login", h.SignIn)
	app.Post("/signup", h.SignUp)
}

func (h *AuthHandler) SignIn(ctx *fiber.Ctx) error {
	var request models.SignInRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		return &models.Response{
			Status: http.StatusBadRequest,
			Data:   nil,
			Err:    "cant parse body",
		}
	}

	if request.Email == "" || request.Password == "" {
		return &models.Response{
			Status: http.StatusBadRequest,
			Data:   nil,
			Err:    "email or password is not valid",
		}
	}

	token, err := h.authService.SignIn(&request)
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusOK).JSON(models.Response{
		Data: token,
		Err:  "",
	})
}

func (h *AuthHandler) SignUp(ctx *fiber.Ctx) error {
	var request models.SignUpRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		return &models.Response{
			Status: http.StatusBadRequest,
			Data:   nil,
			Err:    "cant parse body",
		}
	}

	if request.Email == "" || request.Password == "" {
		return &models.Response{
			Status: http.StatusBadRequest,
			Data:   nil,
			Err:    "email or password is not valid",
		}
	}

	err = h.authService.SignUp(&request)
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusCreated).JSON(models.Response{
		Data: nil,
		Err:  "",
	})
}
