package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Repository struct {
	Db *sql.DB
}

func NewRepository() *Repository {
	databaseHost := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	//databaseName := os.Getenv("db_name")

	//Define DB connection string
	connStr := fmt.Sprintf("postgresql://%s:%s@%s/mydb?sslmode=disable", username, password, databaseHost)

	//connect to db URI
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// close db when not in use
	defer db.Close()

	fmt.Println("Successfully connected!", db)
	return &Repository{Db: db}
}
