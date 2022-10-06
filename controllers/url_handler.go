package controllers

import (
	"koc-digital-case/models"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	_ "github.com/golang-jwt/jwt/v4"
)

type (
	URLHandler struct {
		urlService   URLService
		tokenService TokenService
	}

	URLService interface {
		ShortenURL(request *models.CreateSURLRequest) (string, error)
		GetURLs() ([]models.URLMapping, error)
		RedirectURL(shortenURL string) (string, error)
		DeleteURL(id int) error
	}

	TokenService interface {
		ValidateToken(email string, bearer string) error
	}
)

func NewURLHandler(urlService URLService, tokenService TokenService) *URLHandler {
	return &URLHandler{urlService: urlService, tokenService: tokenService}
}

func (h *URLHandler) RegisterRoutes(app *fiber.App) {
	app.Post("/urls", h.ShortenURL)
	app.Get("/urls", h.GetURLs)
	app.Get("/urls/redirect/:id", h.RedirectURL)
	app.Delete("/urls/:id", h.DeleteURL)
}

func (h *URLHandler) ShortenURL(ctx *fiber.Ctx) error {
	var request models.CreateSURLRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		return &models.Response{
			Status: http.StatusBadRequest,
			Data:   nil,
			Err:    err.Error(),
		}
	}

	token := ctx.Get("Authorization")
	err = h.tokenService.ValidateToken(request.Email, token)
	if err != nil {
		return &models.Response{
			Status: http.StatusBadRequest,
			Data:   nil,
			Err:    err.Error(),
		}
	}

	if request.Email == "" || request.URL == "" {
		return &models.Response{
			Status: http.StatusBadRequest,
			Data:   nil,
			Err:    "url or email cant be empty",
		}
	}

	shortenedURL, err := h.urlService.ShortenURL(&request)
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusCreated).JSON(models.Response{
		Data: shortenedURL,
		Err:  "",
	})
}

func (h *URLHandler) GetURLs(ctx *fiber.Ctx) error {
	email := ctx.Get("email")
	if email == "" {
		return &models.Response{
			Status: http.StatusBadRequest,
			Data:   nil,
			Err:    "email param should be given",
		}
	}

	token := ctx.Get("Authorization")
	err := h.tokenService.ValidateToken(email, token)
	if err != nil {
		return &models.Response{
			Status: http.StatusBadRequest,
			Data:   nil,
			Err:    err.Error(),
		}
	}

	urls, err := h.urlService.GetURLs()
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusOK).JSON(models.Response{
		Data: urls,
		Err:  "",
	})
}

func (h *URLHandler) DeleteURL(ctx *fiber.Ctx) error {
	email := ctx.Get("email")
	if email == "" {
		return &models.Response{
			Status: http.StatusBadRequest,
			Data:   nil,
			Err:    "email param should be given",
		}
	}

	token := ctx.Get("Authorization")
	err := h.tokenService.ValidateToken(email, token)
	if err != nil {
		return &models.Response{
			Status: http.StatusBadRequest,
			Data:   nil,
			Err:    err.Error(),
		}
	}

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return &models.Response{
			Status: http.StatusBadRequest,
			Data:   nil,
			Err:    "id could not parse",
		}
	}

	err = h.urlService.DeleteURL(id)
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusNoContent).JSON(models.Response{
		Data: "",
		Err:  "",
	})
}

func (h *URLHandler) RedirectURL(ctx *fiber.Ctx) error {
	email := ctx.Get("email")
	if email == "" {
		return &models.Response{
			Status: http.StatusBadRequest,
			Data:   nil,
			Err:    "email param should be given at header",
		}
	}

	token := ctx.Get("Authorization")
	err := h.tokenService.ValidateToken(email, token)
	if err != nil {
		return &models.Response{
			Status: http.StatusBadRequest,
			Data:   nil,
			Err:    err.Error(),
		}
	}

	id := ctx.Params("id")
	if id == "" {
		return &models.Response{
			Status: http.StatusBadRequest,
			Data:   nil,
			Err:    "cant parse id",
		}
	}

	original, err := h.urlService.RedirectURL(id)
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusOK).JSON(models.Response{
		Data: original,
		Err:  "",
	})
}
