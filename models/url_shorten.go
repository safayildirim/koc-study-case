package models

type CreateSURLRequest struct {
	Email string
	URL   string
}

type URLMapping struct {
	ID        int
	Original  string
	Shortened string
}
