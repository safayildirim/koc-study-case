package controllers_test

import (
	"bytes"
	"encoding/json"
	"io"
	"koc-digital-case/controllers"
	"koc-digital-case/mocks"
	"koc-digital-case/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSignIn(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	app := fiber.New(fiber.Config{ErrorHandler: models.ErrorHandler})
	mockAuthService := mocks.NewMockAuthService(controller)
	urlHandler := controllers.NewAuthHandler(mockAuthService)
	urlHandler.RegisterRoutes(app)

	t.Run("GivenEmailPasswordValidWhenSignInCalledThenShouldReturnToken", func(t *testing.T) {
		request := models.SignInRequest{
			Email:    "safayildirim54@gmail.com",
			Password: "123456",
		}

		requestBytes, _ := json.Marshal(&request)
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(requestBytes))
		req.Header.Set("Content-Type", "application/json")
		mockAuthService.EXPECT().SignIn(&request).Return("token", nil).Times(1)
		res, err := app.Test(req)
		assert.Nil(t, err)
		bodyBytes, _ := io.ReadAll(res.Body)
		var actualResponse models.Response
		err = json.Unmarshal(bodyBytes, &actualResponse)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, "token", actualResponse.Data)
	})
}
