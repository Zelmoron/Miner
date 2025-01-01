package repository

import (
	"WebSocket/internal/requests"
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var ErrorSelectUser = errors.New("Error Select")

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

	// if err := m.Down(); err != nil && err != migrate.ErrNoChange {
	// 	logrus.Fatalf("Failed to delete migrations: %v", err)
	// }

	return db

}

func (r *Repository) CreateUser(user requests.UserRegRequest, hash string) error {

	query := `INSERT INTO users (name,email,password) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := r.DB.QueryRow(query, user.Name, user.Email, hash).Scan(&id)
	if err != nil {
		logrus.Info(err)
		return err
	}

	query = `INSERT INTO refresh_token (id) VALUES ($1)`
	_, err = r.DB.Exec(query, id)
	if err != nil {
		logrus.Info(err)
		return err
	}
	logrus.Info("Insert succes")
	return nil
}

func (r *Repository) GetUser(user string) error {
	query := "SELECT email FROM users WHERE email = $1"

	var s string
	row := r.DB.QueryRow(query, user).Scan(&s)

	if row == nil {
		logrus.Println("Table has user")
		return nil
	}
	return ErrorSelectUser
}

func (r *Repository) GetUserLogin(user string) (string, int, error) {

	query := "SELECT id,password FROM users WHERE email = $1"
	var id int
	var pass string

	row := r.DB.QueryRow(query, user).Scan(&id, &pass)

	if row != nil {
		logrus.Info("User not found")
		return "", 0, row
	}

	return pass, id, nil
}
