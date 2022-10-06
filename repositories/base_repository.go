package repositories

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type Repository struct {
	Db *sql.DB
}

func NewRepository() *Repository {
	databaseHost := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	connStr := fmt.Sprintf("postgresql://%s:%s@%s/mydb?sslmode=disable", username, password, databaseHost)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err.Error())
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "mydb", driver)
	if err != nil {
		panic(err.Error())
	}

	err = m.Up()
	if err != nil {
		if err != migrate.ErrNoChange {
			panic(err.Error())
		}
	}

	fmt.Println("Successfully connected!", db)
	return &Repository{Db: db}
}
