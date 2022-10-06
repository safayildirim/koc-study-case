package services

import (
	"koc-digital-case/models"
)

var id int

const (
	LimitForFreeUser    = 1
	LimitForPremiumUser = 10
)

type (
	URLService struct {
		urlRepository  URLRepository
		userRepository UserRepository
		generator      Generator
	}

	UserRepository interface {
		GetUser(email string) (*models.User, error)
	}

	Generator interface {
		Decode(url string) int
		Encode(id int) string
	}

	URLRepository interface {
		StoreURLMapping(id int, email, original, shortenURL string) error
		GetUserRemainingBenefits(email string) (int, error)
		UpdateUserUsage(email string, remaining int) error
		GetURLs() ([]models.URLMapping, error)
		GetURL(id int) (string, error)
		GetShortenedURL(url string) (string, error)
		DeleteURL(id int) error
	}
)

func NewURLService(urlRepository URLRepository, userRepository UserRepository, generator Generator) *URLService {
	id = 100000
	return &URLService{urlRepository: urlRepository, userRepository: userRepository, generator: generator}
}

func (s *URLService) ShortenURL(request *models.CreateSURLRequest) (string, error) {
	user, err := s.userRepository.GetUser(request.Email)
	if err != nil {
		return "", err
	}

	userBenefits, err := s.urlRepository.GetUserRemainingBenefits(user.Email)
	if err != nil {
		return "", err
	}

	if !s.DoesUserHaveRemainingBenefit(user.SubscriptionType, userBenefits) {
		return "", &models.Response{
			Status: 500,
			Data:   nil,
			Err:    "you dont have remaining usage",
		}
	}

	shortenedURL, err := s.urlRepository.GetShortenedURL(request.URL)
	if err != nil {
		n := id
		shortenedURL = s.generator.Encode(n)
		err = s.urlRepository.StoreURLMapping(id, request.Email, request.URL, shortenedURL)
		if err != nil {
			return "", err
		}
		id++
	}
	err = s.urlRepository.UpdateUserUsage(user.Email, userBenefits-1)
	if err != nil {
		return "", err
	}
	return shortenedURL, nil
}

func (s *URLService) RedirectURL(shortenURL string) (string, error) {
	decodedID := s.generator.Decode(shortenURL)
	original, err := s.urlRepository.GetURL(decodedID)
	if err != nil {
		return "", err
	}
	return original, nil
}

func (s *URLService) DoesUserHaveRemainingBenefit(subscriptionType, remainingBenefit int) bool {
	if subscriptionType == 0 && remainingBenefit > 0 {
		return true
	}
	if subscriptionType == 1 && remainingBenefit > 0 {
		return true
	}
	return false
}

func (s *URLService) GetURLs() ([]models.URLMapping, error) {
	return s.urlRepository.GetURLs()
}

func (s *URLService) DeleteURL(id int) error {
	return s.urlRepository.DeleteURL(id)
}
