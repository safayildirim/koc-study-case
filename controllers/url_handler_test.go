package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"koc-digital-case/controllers"
	"koc-digital-case/mocks"
	"koc-digital-case/models"
	"koc-digital-case/services"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestShortenURL(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	app := fiber.New(fiber.Config{ErrorHandler: models.ErrorHandler})
	mockURLService := mocks.NewMockURLService(controller)
	tokenService := services.NewTokenService("mysecret")
	urlHandler := controllers.NewURLHandler(mockURLService, tokenService)
	urlHandler.RegisterRoutes(app)

	t.Run("GivenOriginalURLWhenShortenURLCalledThenShouldReturnShortenURL", func(t *testing.T) {
		requestBody := models.CreateSURLRequest{
			Email: "safayildirim54@gmail.com",
			URL:   "www.example.com/blablabla",
		}
		body, _ := json.Marshal(requestBody)
		req := httptest.NewRequest(http.MethodPost, "/urls", bytes.NewBuffer(body))
		token, err := tokenService.CreateToken("safayildirim54@gmail.com")
		assert.Nil(t, err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		mockURLService.EXPECT().ShortenURL(&requestBody).Return("token", nil).Times(1)
		res, err := app.Test(req)
		assert.Nil(t, err)
		bodyBytes, _ := io.ReadAll(res.Body)
		var actualResponse models.Response
		err = json.Unmarshal(bodyBytes, &actualResponse)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusCreated, res.StatusCode)
		assert.Equal(t, "token", actualResponse.Data)
		assert.Equal(t, "", actualResponse.Error())
	})

	t.Run("GivenServiceErrorExistWhenShortenURLCalledThenShouldReturnError", func(t *testing.T) {
		requestBody := models.CreateSURLRequest{
			Email: "safayildirim54@gmail.com",
			URL:   "www.example.com/blablabla",
		}
		body, _ := json.Marshal(requestBody)
		req := httptest.NewRequest(http.MethodPost, "/urls", bytes.NewBuffer(body))
		token, err := tokenService.CreateToken("safayildirim54@gmail.com")
		assert.Nil(t, err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		mockURLService.EXPECT().ShortenURL(&requestBody).Return("", &models.Response{
			Status: 500,
			Data:   nil,
			Err:    "error",
		}).Times(1)
		res, err := app.Test(req)
		assert.Nil(t, err)
		bodyBytes, _ := io.ReadAll(res.Body)
		var actualResponse models.Response
		err = json.Unmarshal(bodyBytes, &actualResponse)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
		assert.Equal(t, nil, actualResponse.Data)
		assert.Equal(t, "error", actualResponse.Error())
	})

	t.Run("GivenURLIsEmptyInRequestWhenShortenURLCalledThenShouldReturnError", func(t *testing.T) {
		requestBody := models.CreateSURLRequest{
			Email: "safayildirim54@gmail.com",
			URL:   "",
		}
		body, _ := json.Marshal(requestBody)
		req := httptest.NewRequest(http.MethodPost, "/urls", bytes.NewBuffer(body))
		token, err := tokenService.CreateToken("safayildirim54@gmail.com")
		assert.Nil(t, err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		res, err := app.Test(req)
		assert.Nil(t, err)
		bodyBytes, _ := io.ReadAll(res.Body)
		var actualResponse models.Response
		err = json.Unmarshal(bodyBytes, &actualResponse)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		assert.Equal(t, nil, actualResponse.Data)
		assert.Equal(t, "url or email cant be empty", actualResponse.Error())
	})

	t.Run("GivenTokenIsInvalidWhenShortenURLCalledThenShouldReturnError", func(t *testing.T) {
		requestBody := models.CreateSURLRequest{
			Email: "safayildirim54@gmail.com",
			URL:   "www.example.com/blablabla",
		}
		body, _ := json.Marshal(requestBody)
		req := httptest.NewRequest(http.MethodPost, "/urls", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		res, err := app.Test(req)
		assert.Nil(t, err)
		bodyBytes, _ := io.ReadAll(res.Body)
		var actualResponse models.Response
		err = json.Unmarshal(bodyBytes, &actualResponse)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		assert.Equal(t, nil, actualResponse.Data)
		assert.Equal(t, "could not find token at header", actualResponse.Error())
	})
}

func TestRedirectURL(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	app := fiber.New(fiber.Config{ErrorHandler: models.ErrorHandler})
	mockURLService := mocks.NewMockURLService(controller)
	tokenService := services.NewTokenService("mysecret")
	urlHandler := controllers.NewURLHandler(mockURLService, tokenService)
	urlHandler.RegisterRoutes(app)

	t.Run("GivenShortenedURLWhenRedirectURLCalledThenShouldReturnOriginalURL", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/urls/redirect/Aa5", nil)
		token, err := tokenService.CreateToken("safayildirim54@gmail.com")
		assert.Nil(t, err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		req.Header.Set("email", "safayildirim54@gmail.com")
		mockURLService.EXPECT().RedirectURL("Aa5").Return("www.google.com.tr", nil).Times(1)
		res, err := app.Test(req)
		assert.Nil(t, err)
		bodyBytes, _ := io.ReadAll(res.Body)
		var actualResponse models.Response
		err = json.Unmarshal(bodyBytes, &actualResponse)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, "www.google.com.tr", actualResponse.Data)
		assert.Equal(t, "", actualResponse.Error())
	})

	t.Run("GivenServiceErrorExistWhenRedirectURLCalledThenShouldReturnError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/urls/redirect/Aa5", nil)
		token, err := tokenService.CreateToken("safayildirim54@gmail.com")
		assert.Nil(t, err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		req.Header.Set("email", "safayildirim54@gmail.com")
		mockURLService.EXPECT().RedirectURL("Aa5").Return("", &models.Response{
			Status: 500,
			Data:   nil,
			Err:    "error",
		}).Times(1)
		res, err := app.Test(req)
		assert.Nil(t, err)
		bodyBytes, _ := io.ReadAll(res.Body)
		var actualResponse models.Response
		err = json.Unmarshal(bodyBytes, &actualResponse)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
		assert.Equal(t, nil, actualResponse.Data)
		assert.Equal(t, "error", actualResponse.Error())
	})

	t.Run("GivenShortenURLNotExistWhenRedirectURLCalledThenShouldReturnError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/urls/redirect/Aa5", nil)
		token, err := tokenService.CreateToken("safayildirim54@gmail.com")
		assert.Nil(t, err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		req.Header.Set("email", "safayildirim54@gmail.com")
		mockURLService.EXPECT().RedirectURL("Aa5").Return("", &models.Response{
			Status: 404,
			Data:   nil,
			Err:    "original url not found",
		}).Times(1)
		res, err := app.Test(req)
		assert.Nil(t, err)
		bodyBytes, _ := io.ReadAll(res.Body)
		var actualResponse models.Response
		err = json.Unmarshal(bodyBytes, &actualResponse)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
		assert.Equal(t, nil, actualResponse.Data)
		assert.Equal(t, "original url not found", actualResponse.Error())
	})

	t.Run("GivenTokenIsInvalidWhenRedirectURLCalledThenShouldReturnError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/urls/redirect/Aa5", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("email", "safayildirim54@gmail.com")
		res, err := app.Test(req)
		assert.Nil(t, err)
		bodyBytes, _ := io.ReadAll(res.Body)
		var actualResponse models.Response
		err = json.Unmarshal(bodyBytes, &actualResponse)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		assert.Equal(t, nil, actualResponse.Data)
		assert.Equal(t, "could not find token at header", actualResponse.Error())
	})
}

func TestGetURLs(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	app := fiber.New(fiber.Config{ErrorHandler: models.ErrorHandler})
	mockURLService := mocks.NewMockURLService(controller)
	tokenService := services.NewTokenService("mysecret")
	urlHandler := controllers.NewURLHandler(mockURLService, tokenService)
	urlHandler.RegisterRoutes(app)

	t.Run("GivenUserWhenGetURLsCalledThenShouldReturnURLs", func(t *testing.T) {
		urls := []models.URLMapping{
			{
				ID:           0,
				Original:     "test1",
				ShortenedURL: "test1",
			},
			{
				ID:           1,
				Original:     "test2",
				ShortenedURL: "test2",
			},
		}

		req := httptest.NewRequest(http.MethodGet, "/urls", nil)
		token, err := tokenService.CreateToken("safayildirim54@gmail.com")
		assert.Nil(t, err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		req.Header.Set("email", "safayildirim54@gmail.com")
		mockURLService.EXPECT().GetURLs().Return(urls, nil).Times(1)
		res, err := app.Test(req)
		assert.Nil(t, err)
		bodyBytes, _ := io.ReadAll(res.Body)
		var actualResponse models.Response
		err = json.Unmarshal(bodyBytes, &actualResponse)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode)
		actualURLs := actualResponse.Data.([]interface{})
		assert.Equal(t, 2, len(actualURLs))
		assert.Equal(t, "", actualResponse.Error())
	})

	t.Run("GivenEmailIsEmptyWhenGetURLsCalledThenShouldReturnBadRequest", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/urls", nil)
		token, err := tokenService.CreateToken("safayildirim54@gmail.com")
		assert.Nil(t, err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		res, err := app.Test(req)
		assert.Nil(t, err)
		bodyBytes, _ := io.ReadAll(res.Body)
		var actualResponse models.Response
		err = json.Unmarshal(bodyBytes, &actualResponse)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		assert.Equal(t, "email param should be given", actualResponse.Error())
	})
}

func TestDeleteURL(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	app := fiber.New(fiber.Config{ErrorHandler: models.ErrorHandler})
	mockURLService := mocks.NewMockURLService(controller)
	tokenService := services.NewTokenService("mysecret")
	urlHandler := controllers.NewURLHandler(mockURLService, tokenService)
	urlHandler.RegisterRoutes(app)

	t.Run("GivenIdExistWhenDeleteURLCalledThenShouldDeleteURL", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/urls/100000?email=safayildirim54@gmail.com", nil)
		token, err := tokenService.CreateToken("safayildirim54@gmail.com")
		assert.Nil(t, err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		req.Header.Set("email", "safayildirim54@gmail.com")
		mockURLService.EXPECT().DeleteURL(100000).Return(nil).Times(1)
		res, err := app.Test(req)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNoContent, res.StatusCode)
	})
}
