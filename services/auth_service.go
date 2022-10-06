package services

import (
	"errors"
	"koc-digital-case/models"

	"golang.org/x/crypto/bcrypt"
)

type (
	AuthService struct {
		authRepository AuthRepository
		tokenMaker     TokenMaker
	}

	AuthRepository interface {
		GetUser(email string) (*models.User, error)
		CreateUser(user *models.User) error
		CreateUserUsage(email string, limit int) error
	}

	TokenMaker interface {
		CreateToken(email string) (string, error)
	}
)

func NewAuthService(authRepository AuthRepository, tokenMaker TokenMaker) *AuthService {
	return &AuthService{authRepository: authRepository, tokenMaker: tokenMaker}
}

func (s *AuthService) SignUp(request *models.SignUpRequest) error {
	userExist := s.DoesUserExist(request.Email)
	if userExist {
		return errors.New("email is already taken")
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("password encryption failed")
	}

	user := models.User{
		Email:            request.Email,
		Password:         string(pass),
		SubscriptionType: request.SubscriptionType,
	}

	err = s.authRepository.CreateUser(&user)
	if err != nil {
		return err
	}

	userLimit := LimitForFreeUser
	if user.SubscriptionType == 1 {
		userLimit = LimitForPremiumUser
	}

	err = s.authRepository.CreateUserUsage(request.Email, userLimit)
	if err != nil {
		return err
	}
	return nil
}

func (s *AuthService) SignIn(request *models.SignInRequest) (string, error) {
	user, err := s.authRepository.GetUser(request.Email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", errors.New("invalid username or password")
	}

	token, err := s.tokenMaker.CreateToken(request.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) DoesUserExist(email string) bool {
	_, err := s.authRepository.GetUser(email)
	if err != nil {
		return false
	}
	return true
}
