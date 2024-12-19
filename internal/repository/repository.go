package repository

import (
	"WebSocket/internal/requests"
	"database/sql"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type Repository struct {
	DB *sql.DB
}

func New(DB *sql.DB) *Repository {
	return &Repository{
		DB: DB,
	}
}

func CreateTable() *sql.DB {
	user := os.Getenv("name")           //user Postgres
	password := os.Getenv("dbpassword") //password Postgres
	dbname := os.Getenv("dbname")       //name of the database
	host := os.Getenv("dbhost")         //host
	port := os.Getenv("dbport")         //database`s port

	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname))
	if err != nil {
		logrus.Fatalf("Failed to connect to database: %v", err)
		return nil
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		logrus.Fatalf("Failed to create migration driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"postgres", driver)
	if err != nil {
		logrus.Fatalf("Failed to create migrate instance: %v", err)
	}

	//Поднять миграции
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		logrus.Fatalf("Failed to apply migrations: %v", err)
	}
	logrus.Info("Succes migrations")

	//todo add down
	//

	return db

}

func (r *Repository) CreateUser(user requests.UserRegRequest) {

	query := "SELECT * FROM users WHERE email=$1"
	row := r.DB.QueryRow(query, user.Email)

	if row != nil {
		fmt.Println("Table has user")
	}
	// query = `INSERT INTO users (name, age) VALUES ($1, $2)`
	// _, err := r.DB.Exec(query, "John Doe", 30)
	// if err != nil {
	// 	logrus.Fatal(err)
	// }

	logrus.Info("Insert succes")
}
