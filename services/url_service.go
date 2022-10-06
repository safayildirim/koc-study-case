package services

import (
	"fmt"
	"koc-digital-case/models"
	"strings"
)

var baseID int

const (
	LimitForFreeUser    = 1
	LimitForPremiumUser = 10
	Alphabet            = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	Base                = len(Alphabet)
)

type URLService struct {
	urlRepository  URLRepository
	userRepository UserRepository
}

type UserRepository interface {
	GetUser(email string) (*models.User, error)
}

type URLRepository interface {
	StoreURLMapping(email, original string, id int) error
	GetUserRemainingBenefits(email string) (int, error)
	UpdateUserUsage(email string, remaining int) error
	GetURLs() ([]models.URLMapping, error)
	GetURL(id int) (string, error)
	DeleteURL(id int) error
}

func NewURLService(urlRepository URLRepository, userRepository UserRepository) *URLService {
	baseID = 100000
	return &URLService{urlRepository: urlRepository, userRepository: userRepository}
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

	n := baseID
	var shortenedURLID string
	for n > 0 {
		shortenedURLID = fmt.Sprintf("%s%s", string(Alphabet[n%62]), shortenedURLID)
		n = n / Base
	}

	shortenedURL := fmt.Sprintf("%s", shortenedURLID)
	err = s.urlRepository.StoreURLMapping(request.Email, request.URL, baseID)
	if err != nil {
		return "", err
	}
	err = s.urlRepository.UpdateUserUsage(user.Email, userBenefits-1)
	if err != nil {
		return "", err
	}
	baseID++
	return shortenedURL, nil
}

func (s *URLService) RedirectURL(shortenURL string) (string, error) {
	id := 0
	for _, c := range shortenURL {
		id = (id * Base) + strings.Index(Alphabet, string(c))
	}
	original, err := s.urlRepository.GetURL(id)
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
