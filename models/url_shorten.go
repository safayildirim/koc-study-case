package models

type CreateSURLRequest struct {
	Email string `json:"email"`
	URL   string `json:"url"`
}

type URLMapping struct {
	ID           int    `json:"id"`
	Original     string `json:"original"`
	ShortenedURL string `json:"shortened_url"`
}
