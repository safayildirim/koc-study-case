package repositories

import (
	"database/sql"
	"errors"
	"koc-digital-case/models"
)

type UserRepository struct {
	repository *Repository
}

func NewUserRepository(repository *Repository) *UserRepository {
	return &UserRepository{repository: repository}
}

func (s *UserRepository) CreateUser(user *models.User) error {
	res, err := s.repository.Db.Exec("INSERT INTO users (email, password, subscription_type) VALUES ($1, $2, $3)",
		user.Email, user.Password, user.SubscriptionType)
	if err != nil {
		return &models.Response{
			Status: 500,
			Data:   nil,
			Err:    err.Error(),
		}
	}
	ar, err := res.RowsAffected()
	if err != nil {
		return &models.Response{
			Status: 500,
			Data:   nil,
			Err:    err.Error(),
		}
	}
	if ar == 0 {
		return &models.Response{
			Status: 500,
			Data:   nil,
			Err:    "could not insert user to db",
		}
	}
	return nil
}

func (s *UserRepository) CreateUserUsage(email string, limit int) error {
	res, err := s.repository.Db.Exec("INSERT INTO usages (email, remaining) VALUES ($1,$2)", email, limit)
	if err != nil {
		return &models.Response{
			Status: 500,
			Data:   nil,
			Err:    err.Error(),
		}
	}
	affectedRows, err := res.RowsAffected()
	if err != nil {
		return &models.Response{
			Status: 404,
			Data:   nil,
			Err:    "user not found",
		}
	}
	if affectedRows == 0 {
		return errors.New("could not insert to user usage")
	}
	return nil
}

func (s *UserRepository) GetUser(email string) (*models.User, error) {
	row := s.repository.Db.QueryRow("SELECT email, password, subscription_type FROM users WHERE email = $1", email)
	var user models.User
	err := row.Scan(&user.Email, &user.Password, &user.SubscriptionType)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &models.Response{
				Status: 404,
				Data:   nil,
				Err:    "user not found",
			}
		}
		return nil, &models.Response{
			Status: 500,
			Data:   nil,
			Err:    err.Error(),
		}
	}
	return &user, nil
}
