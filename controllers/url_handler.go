package controllers

import (
	"koc-digital-case/models"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	_ "github.com/golang-jwt/jwt/v4"
)

type URLHandler struct {
	urlService URLService
}

type URLService interface {
	ShortenURL(request *models.CreateSURLRequest) (string, error)
	GetURLs() ([]models.URLMapping, error)
	DeleteURL(id int) error
}

func NewURLHandler(urlService URLService) *URLHandler {
	return &URLHandler{urlService: urlService}
}

func (h *URLHandler) RegisterRoutes(app *fiber.App) {
	app.Group("/api")
	app.Post("/urls", h.ShortenURL)
	app.Get("/urls", h.GetURLs)
	app.Delete("/urls/:id", h.DeleteURL)
}

func (h *URLHandler) ShortenURL(ctx *fiber.Ctx) error {
	var request models.CreateSURLRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		return err
	}

	err = validateToken(ctx, request.Email)
	if err != nil {
		return err
	}

	shortenedURL, err := h.urlService.ShortenURL(&request)
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusCreated).JSON(models.Response{
		Data:  shortenedURL,
		Error: "",
	})
}

func (h *URLHandler) GetURLs(ctx *fiber.Ctx) error {
	email := ctx.Query("email")
	if email == "" {
		return ctx.Status(http.StatusBadRequest).JSON(models.Response{Error: "email param should be given"})
	}
	err := validateToken(ctx, email)
	if err != nil {
		return err
	}

	urls, err := h.urlService.GetURLs()
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusCreated).JSON(models.Response{
		Data:  urls,
		Error: "",
	})
}

func (h *URLHandler) DeleteURL(ctx *fiber.Ctx) error {
	email := ctx.Query("email")
	if email == "" {
		return ctx.Status(http.StatusBadRequest).JSON(models.Response{Error: "email param should be given"})
	}
	err := validateToken(ctx, email)
	if err != nil {
		return err
	}

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(models.Response{Error: "id could not parse"})
	}

	err = h.urlService.DeleteURL(id)
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusNoContent).JSON(models.Response{
		Data:  "",
		Error: "",
	})
}

func validateToken(ctx *fiber.Ctx, email string) error {
	token := ctx.Locals("token").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	emailInToken := claims["email"].(string)
	if emailInToken != email {
		return ctx.Status(http.StatusBadRequest).JSON(models.Response{Error: "could not authenticated"})
	}
	return nil
}
