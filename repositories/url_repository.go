package repositories

import (
	"database/sql"
	"fmt"
	"koc-digital-case/models"
)

type URLRepository struct {
	repository *Repository
}

func NewURLRepository(repository *Repository) *URLRepository {
	return &URLRepository{repository: repository}
}

func (r *URLRepository) StoreURLMapping(id int, email, original, shortenURL string) error {
	_, err := r.repository.Db.Exec("INSERT INTO urls (id, email, original, shortened_url) VALUES ($1,$2,$3,$4)",
		id, email, original, shortenURL)
	if err != nil {
		return &models.Response{
			Status: 500,
			Data:   nil,
			Err:    err.Error(),
		}
	}
	return nil
}

func (r *URLRepository) GetShortenedURL(url string) (string, error) {
	row := r.repository.Db.QueryRow("SELECT shortened_url FROM urls WHERE original = $1", url)
	var shortenedID *string
	err := row.Scan(&shortenedID)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", &models.Response{
				Status: 404,
				Data:   nil,
				Err:    "url not found",
			}
		}
		return "", &models.Response{
			Status: 500,
			Data:   nil,
			Err:    err.Error(),
		}
	}
	return *shortenedID, nil
}

func (r *URLRepository) GetURL(id int) (string, error) {
	row := r.repository.Db.QueryRow("SELECT original FROM urls WHERE id = $1", id)
	var url *string
	err := row.Scan(&url)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", &models.Response{
				Status: 404,
				Data:   nil,
				Err:    "original url not found",
			}
		}
		return "", &models.Response{
			Status: 500,
			Data:   nil,
			Err:    err.Error(),
		}
	}
	return *url, nil
}

func (r *URLRepository) GetUserRemainingBenefits(email string) (int, error) {
	row := r.repository.Db.QueryRow("SELECT remaining FROM usages WHERE email = $1", email)
	var remaining *int
	err := row.Scan(&remaining)
	if err != nil {
		if err == sql.ErrNoRows {
			return -1, &models.Response{
				Status: 404,
				Data:   nil,
				Err:    fmt.Sprintf("benefit information for user '%s' is not found", email),
			}
		}
		return -1, &models.Response{
			Status: 500,
			Data:   nil,
			Err:    err.Error(),
		}
	}
	return *remaining, nil
}

func (r *URLRepository) UpdateUserUsage(email string, remaining int) error {
	res, err := r.repository.Db.Exec("UPDATE usages SET remaining=$1 WHERE email=$2", remaining, email)
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
			Status: 500,
			Data:   nil,
			Err:    err.Error(),
		}
	}
	if affectedRows == 0 {
		return &models.Response{
			Status: 500,
			Data:   nil,
			Err:    "could not update user remaining benefits",
		}
	}
	return nil
}

func (r *URLRepository) GetURLs() ([]models.URLMapping, error) {
	var urlMappings []models.URLMapping
	rows, err := r.repository.Db.Query("SELECT id, original, shortened_url FROM urls")
	defer rows.Close()
	if err != nil {
		return nil, &models.Response{
			Status: 500,
			Data:   nil,
			Err:    err.Error(),
		}
	}
	for rows.Next() {
		var urlMapping models.URLMapping
		err = rows.Scan(&urlMapping.ID, &urlMapping.Original, &urlMapping.ShortenedURL)
		if err != nil {
			return nil, &models.Response{
				Status: 500,
				Data:   nil,
				Err:    err.Error(),
			}
		}
		urlMappings = append(urlMappings, urlMapping)
	}
	return urlMappings, nil
}

func (r *URLRepository) DeleteURL(id int) error {
	res, err := r.repository.Db.Exec("DELETE FROM urls WHERE id=$1", id)
	if err != nil {
		return &models.Response{
			Status: 500,
			Data:   nil,
			Err:    err.Error(),
		}
	}
	affectedRows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affectedRows == 0 {
		return &models.Response{
			Status: 404,
			Data:   nil,
			Err:    "id not found",
		}
	}
	return nil
}
