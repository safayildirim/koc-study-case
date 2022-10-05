package repositories

import (
	"errors"
	"fmt"
	"koc-digital-case/models"
)

type AuthRepository struct {
	repository *Repository
}

func NewAuthRepository(repository *Repository) *AuthRepository {
	return &AuthRepository{repository: repository}
}

func (s *AuthRepository) CreateUser(user *models.User) error {
	query := fmt.Sprintf("INSERT INTO users VALUES (NULL, %s, %s, %d)", user.Email, user.Password, user.SubscriptionType)
	rows, err := s.repository.Db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}

func (s *AuthRepository) CreateUserUsage(email string) error {
	res, err := s.repository.Db.Exec("INSERT INTO usages VALUES (?,0)", email)
	if err != nil {
		return err
	}
	affectedRows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affectedRows == 0 {
		return errors.New("could not insert to user usage")
	}
	return nil
}

func (s *AuthRepository) GetUser(email string) (*models.User, error) {
	row := s.repository.Db.QueryRow("SELECT * FROM users WHERE email = ?", email)
	var user models.User
	err := row.Scan(&user.Email, &user.Password, &user.SubscriptionType)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
