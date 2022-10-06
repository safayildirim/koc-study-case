package services_test

import (
	"koc-digital-case/mocks"
	"koc-digital-case/models"
	"koc-digital-case/services"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestShortenURL(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockAuthRepository := mocks.NewMockURLAuthRepository(controller)
	mockURLRepository := mocks.NewMockURLRepository(controller)
	urlGenerator := services.NewURLGenerator()

	t.Run("GivenUserExistsAndHaveBenefitWhenShortenURLCalledThenShouldReturnShortenURL", func(t *testing.T) {
		service := services.NewURLService(mockURLRepository, mockAuthRepository, urlGenerator)
		request := models.CreateSURLRequest{
			Email: "safa.yildirim54@gmail.com",
			URL:   "www.example.com/blablablabla",
		}
		mockAuthRepository.EXPECT().GetUser("safa.yildirim54@gmail.com").Return(&models.User{
			Email:            "safa.yildirim54@gmail.com",
			Password:         "123456",
			SubscriptionType: 1,
		}, nil).Times(1)
		mockURLRepository.EXPECT().GetUserRemainingBenefits("safa.yildirim54@gmail.com").Return(1, nil).Times(1)
		mockURLRepository.EXPECT().GetShortenedURL("www.example.com/blablablabla").Return("", &models.Response{
			Status: 404,
			Data:   nil,
			Err:    "url not found",
		}).Times(1)
		mockURLRepository.EXPECT().StoreURLMapping(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)
		mockURLRepository.EXPECT().UpdateUserUsage(gomock.Any(), gomock.Any()).Return(nil).Times(1)
		expectedShortenedURL := "Aa4"
		actualShortened, err := service.ShortenURL(&request)
		assert.Nil(t, err)
		assert.Equal(t, expectedShortenedURL, actualShortened)
	})
}

func TestRedirectURL(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockAuthRepository := mocks.NewMockURLAuthRepository(controller)
	mockURLRepository := mocks.NewMockURLRepository(controller)
	urlGenerator := services.NewURLGenerator()

	t.Run("GivenIdExistInDBWhenRedirectURLCalledThenShouldReturnOriginalURL", func(t *testing.T) {
		service := services.NewURLService(mockURLRepository, mockAuthRepository, urlGenerator)
		mockURLRepository.EXPECT().GetURL(100000).Return("www.example.com/blablablabla", nil).Times(1)
		expectedOriginalURL := "www.example.com/blablablabla"
		actualURL, err := service.RedirectURL("Aa4")
		assert.Nil(t, err)
		assert.Equal(t, expectedOriginalURL, actualURL)
	})
}

func TestGetURLs(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockAuthRepository := mocks.NewMockURLAuthRepository(controller)
	mockURLRepository := mocks.NewMockURLRepository(controller)
	urlGenerator := services.NewURLGenerator()

	t.Run("GivenURLsExistInDBWhenGetURLsCalledThenShouldReturnURLs", func(t *testing.T) {
		service := services.NewURLService(mockURLRepository, mockAuthRepository, urlGenerator)
		mockURLRepository.EXPECT().GetURLs().Return([]models.URLMapping{
			{
				ID:           100000,
				Original:     "test",
				ShortenedURL: "test",
			},
			{
				ID:           100001,
				Original:     "test2",
				ShortenedURL: "test2",
			},
		}, nil).Times(1)
		actualURLs, err := service.GetURLs()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(actualURLs))
	})
}

func TestDeleteURL(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockAuthRepository := mocks.NewMockURLAuthRepository(controller)
	mockURLRepository := mocks.NewMockURLRepository(controller)
	urlGenerator := services.NewURLGenerator()

	t.Run("GivenURLExistInDBWhenDeleteURLsCalledThenShouldDeleteURLFromDB", func(t *testing.T) {
		service := services.NewURLService(mockURLRepository, mockAuthRepository, urlGenerator)
		mockURLRepository.EXPECT().DeleteURL(100000).Return(nil).Times(1)
		err := service.DeleteURL(100000)
		assert.Nil(t, err)
	})

	t.Run("GivenURLNotExistInDBWhenDeleteURLsCalledThenShouldReturnError", func(t *testing.T) {
		service := services.NewURLService(mockURLRepository, mockAuthRepository, urlGenerator)
		mockURLRepository.EXPECT().DeleteURL(100000).Return(&models.Response{
			Status: 404,
			Data:   nil,
			Err:    "",
		}).Times(1)
		err := service.DeleteURL(100000)
		assert.NotNil(t, err)
	})
}
