package repository

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
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
	user := os.Getenv("name")           //пользователь Postgres
	password := os.Getenv("dbpassword") //Пароль Postgres
	dbname := os.Getenv("dbname")       //Название базы данных
	host := os.Getenv("dbhost")         //Хост базы данных
	port := os.Getenv("dbport")         //Прт базы данных

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

	return db

}

func (r *Repository) CreateUser() {

}
