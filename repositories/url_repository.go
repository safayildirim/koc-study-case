package repositories

import (
	"errors"
	"koc-digital-case/models"
)

type URLRepository struct {
	repository *Repository
}

func NewURLRepository(repository *Repository) *URLRepository {
	return &URLRepository{repository: repository}
}

func (r *URLRepository) StoreURLMapping(email, original, shortened string) error {
	_, err := r.repository.Db.Exec("INSERT INTO urls VALUES (NULL, ?,?,?)", email, original, shortened)
	if err != nil {
		return err
	}
	return nil
}

func (r *URLRepository) GetUserRemainingBenefits(email string) (int, error) {
	row := r.repository.Db.QueryRow("SELECT remaining FROM usages WHERE email = ?", email)
	var remaining *int
	err := row.Scan(&remaining)
	if err != nil {
		return -1, err
	}
	return *remaining, nil
}

func (r *URLRepository) UpdateUserUsage(email string, remaining int) error {
	res, err := r.repository.Db.Exec("UPDATE usages SET remaining=? WHERE email=?", remaining, email)
	if err != nil {
		return err
	}
	affectedRows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affectedRows == 0 {
		return errors.New("could not update user remaining benefits")
	}
	return nil
}

func (r *URLRepository) GetURLs() ([]models.URLMapping, error) {
	var urlMappings []models.URLMapping
	rows, err := r.repository.Db.Query("SELECT id, original, shortened FROM urls")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var urlMapping models.URLMapping
		err = rows.Scan(&urlMapping.ID, &urlMapping.Original, &urlMapping.Shortened)
		if err != nil {
			return nil, err
		}
		urlMappings = append(urlMappings, urlMapping)
	}
	return urlMappings, nil
}

func (r *URLRepository) DeleteURL(id int) error {
	_, err := r.repository.Db.Exec("DELETE FROM urls WHERE id=?", id)
	if err != nil {
		return err
	}
	return nil
}
