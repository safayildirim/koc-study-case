package services_test

import (
	"koc-digital-case/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	urlGenerator := services.NewURLGenerator()
	actualEncoded := urlGenerator.Encode(100000)
	assert.Equal(t, "Aa4", actualEncoded)
}

func TestDecode(t *testing.T) {
	urlGenerator := services.NewURLGenerator()
	actualDecoded := urlGenerator.Decode("Aa4")
	assert.Equal(t, 100000, actualDecoded)
}
