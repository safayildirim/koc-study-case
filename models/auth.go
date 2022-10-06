package models

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpRequest struct {
	Email            string `json:"email"`
	Password         string `json:"password"`
	SubscriptionType int    `json:"subscription_type"`
}

type User struct {
	Email            string `json:"email"`
	Password         string `json:"password"`
	SubscriptionType int    `json:"subscription_type"`
}
