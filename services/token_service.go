package services

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	Bearer = "Bearer"
)

type TokenService struct {
	secret string
}

func NewTokenService(secret string) *TokenService {
	if secret == "" {
		panic("secret can not be blank")
	}
	return &TokenService{secret: secret}
}

func (s *TokenService) extractToken(token string) (string, error) {
	if token == "" {
		return "", errors.New("could not find token at header")
	}
	l := len(Bearer)
	if len(token) > l+1 && strings.EqualFold(token[:l], Bearer) {
		return token[l+1:], nil
	}
	return "", errors.New("missing or malformed JWT")
}

type MyClaims struct {
	Email string
	Exp   int64
}

func (c *MyClaims) Valid() error {
	if c.Exp < time.Now().Unix() {
		return errors.New("token expired")
	}
	return nil
}

func (s *TokenService) ValidateToken(email string, tokenString string) error {
	tokenString, err := s.extractToken(tokenString)
	if err != nil {
		return err
	}
	tkn, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secret), nil
	})
	if err != nil {
		return err
	}
	if !tkn.Valid {
		return errors.New("invalid token")
	}
	claims, ok := tkn.Claims.(*MyClaims)
	if !ok {
		return errors.New("cant parse claims")
	}
	if claims.Email != email {
		return errors.New("could not validate user")
	}
	return nil
}

func (s *TokenService) CreateToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(s.secret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
