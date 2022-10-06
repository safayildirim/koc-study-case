package services_test

import (
	"fmt"
	"koc-digital-case/mocks"
	"koc-digital-case/models"
	"koc-digital-case/services"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestShortenURL(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockAuthRepository := mocks.NewMockURLAuthRepository(controller)
	mockURLRepository := mocks.NewMockURLRepository(controller)

	t.Run("GivenUserExistsAndHaveBenefitWhenShortenURLCalledThenShouldReturnShortenURL", func(t *testing.T) {
		service := services.NewURLService(mockURLRepository, mockAuthRepository)
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
		mockURLRepository.EXPECT().StoreURLMapping(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)
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

	t.Run("GivenIdExistInDBWhenRedirectURLCalledThenShouldReturnOriginalURL", func(t *testing.T) {
		service := services.NewURLService(mockURLRepository, mockAuthRepository)
		mockURLRepository.EXPECT().GetURL(100000).Return("www.example.com/blablablabla", nil).Times(1)
		expectedOriginalURL := "www.example.com/blablablabla"
		actualURL, err := service.RedirectURL("Aa4")
		assert.Nil(t, err)
		assert.Equal(t, expectedOriginalURL, actualURL)
	})
}

func TestName(t *testing.T) {
	n := 100000
	var shortenedURLID string
	for n > 0 {
		shortenedURLID = fmt.Sprintf("%s%s", string(services.Alphabet[n%62]), shortenedURLID)
		n = n / 62
	}
	fmt.Println(shortenedURLID)
}

func TestName2(t *testing.T) {
	id := 0
	s := "Aa4"
	for _, c := range s {
		id = (id * 62) + strings.Index(services.Alphabet, string(c))
	}
	fmt.Println(id)
}
