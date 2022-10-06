package services_test

import (
	"koc-digital-case/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetToken(t *testing.T) {
	t.Run("GivenEmailWhenGetTokenCalledThenShouldReturnToken", func(t *testing.T) {
		email := "safayildirim54@gmail.com"
		service := services.NewTokenService("mysecret")
		token, err := service.CreateToken(email)
		assert.Nil(t, err)
		assert.True(t, len(token) > 0)
	})
}
