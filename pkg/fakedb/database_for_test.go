package fakedb

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func SetupDB() (*sql.DB, error) {
	user, pass, port, db_name := "postgres", "123", 5432, "test_db"
	db, err := sql.Open("postgres", fmt.Sprintf("postgresql://%s:%s@localhost:%d/?sslmode=disable", user, pass, port))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", db_name))
	if err != nil {
		return nil, fmt.Errorf("failed to create database: %v", err)
	}

	db, err = sql.Open("postgres", fmt.Sprintf("postgresql://%s:%s@localhost:%d/%s?sslmode=disable", user, pass, port, db_name))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to test database: %v", err)
	}
	file, err := os.ReadFile("test_db_start.sql")
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(string(file))
	if err != nil {
		return nil, err
	}

	return db, nil
}

func DropDB(db *sql.DB) error {
	db_name := "test_db"
	_, err := db.Exec(fmt.Sprintf("DROP DATABASE %s", db_name))
	return err
}
