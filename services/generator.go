package services

import (
	"fmt"
	"strings"
)

const (
	Alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	Base     = len(Alphabet)
)

type URLGenerator struct {
}

func NewURLGenerator() *URLGenerator {
	return &URLGenerator{}
}

func (s *URLGenerator) Decode(url string) int {
	convertedID := 0
	for _, c := range url {
		convertedID = (convertedID * Base) + strings.Index(Alphabet, string(c))
	}
	return convertedID
}

func (s *URLGenerator) Encode(id int) string {
	var shortenedURLID string
	for id > 0 {
		shortenedURLID = fmt.Sprintf("%s%s", string(Alphabet[id%62]), shortenedURLID)
		id = id / Base
	}
	return fmt.Sprintf("%s", shortenedURLID)
}
