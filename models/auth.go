package models

type SignInRequest struct {
	Email    string
	Password string
}

type SignUpRequest struct {
	Email            string
	Password         string
	SubscriptionType int
}

type User struct {
	Email            string
	Password         string
	SubscriptionType int
}
