package services

import (
	"errors"
	"fmt"
	"koc-digital-case/models"
)

const (
	LimitForFreeUser    = 1
	LimitForPremiumUser = 10
)

type URLService struct {
	urlRepository  URLRepository
	authRepository URLAuthRepository
}

type URLAuthRepository interface {
	GetUser(email string) (*models.User, error)
	GetLastGivenID() int
}

type URLRepository interface {
	StoreURLMapping(email, original, shortened string) error
	GetUserRemainingBenefits(email string) (int, error)
	UpdateUserUsage(email string, remaining int) error
	GetURLs() ([]models.URLMapping, error)
	DeleteURL(id int) error
}

func NewURLService(urlRepository URLRepository, authRepository URLAuthRepository) *URLService {
	return &URLService{urlRepository: urlRepository, authRepository: authRepository}
}

func (s *URLService) ShortenURL(request *models.CreateSURLRequest) (string, error) {
	user, err := s.authRepository.GetUser(request.Email)
	if err != nil {
		return "", err
	}

	userBenefits, err := s.urlRepository.GetUserRemainingBenefits(user.Email)
	if err != nil {
		return "", err
	}

	if !s.DoesUserHaveRemainingBenefit(user.SubscriptionType, userBenefits) {
		return "", errors.New("you dont have remaining usage")
	}

	lastGivenID := s.authRepository.GetLastGivenID()
	shortenedURL := fmt.Sprintf("shorter.com/u%d", lastGivenID+1)
	err = s.urlRepository.StoreURLMapping(request.Email, request.URL, shortenedURL)
	if err != nil {
		return "", err
	}
	err = s.urlRepository.UpdateUserUsage(user.Email, userBenefits+1)
	if err != nil {
		return "", err
	}
	return shortenedURL, nil
}

func (s *URLService) DoesUserHaveRemainingBenefit(subscriptionType, remainingBenefit int) bool {
	if subscriptionType == 0 && remainingBenefit < LimitForFreeUser {
		return true
	}
	if subscriptionType == 1 && remainingBenefit < LimitForPremiumUser {
		return true
	}
	return false
}

func (s *URLService) GetURls() ([]models.URLMapping, error) {
	return s.urlRepository.GetURLs()
}

func (s *URLService) DeleteURL(id int) error {
	return s.urlRepository.DeleteURL(id)
}
