package services

import (
	"errors"
	"koc-digital-case/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	authRepository AuthRepository
}

type AuthRepository interface {
	GetUser(email string) (*models.User, error)
	CreateUser(user *models.User) error
	CreateUserUsage(email string) error
}

func NewAuthService(authRepository AuthRepository) *AuthService {
	return &AuthService{authRepository: authRepository}
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

	err = s.authRepository.CreateUserUsage(request.Email)
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
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return "", errors.New("invalid username or password")
	}

	secret := os.Getenv("SECRET")
	claims := jwt.MapClaims{
		"email": request.Email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (s *AuthService) DoesUserExist(email string) bool {
	_, err := s.authRepository.GetUser(email)
	if err != nil {
		return false
	}
	return true
}
