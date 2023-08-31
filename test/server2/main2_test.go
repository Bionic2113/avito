package server2_test

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/Bionic2113/avito/cmd/server"
	"github.com/Bionic2113/avito/internal/repository"
	_ "github.com/lib/pq"
)

var (
	userRepo        repository.UserRepository
	segmentRepo     repository.SegmentRepository
	userSegmentRepo repository.UserSegmentRepository
)

func TestMain(m *testing.M) {
	db, err := SetupDB()
	if err != nil {
		log.Fatal("Не создалась бд", err)
	}

	userRepo = &repository.UserRepoDB{DB: db}
	segmentRepo = &repository.SegmentRepositoryDB{DB: db}
	userSegmentRepo = &repository.UserSegmentDB{DB: db}
	server.Port = 8081
	go server.StartServer(userRepo, segmentRepo, userSegmentRepo)

	m.Run()
	if err = DropDB(db); err != nil {
		log.Fatal("Не удалилась бд", err)
	}
}
func SetupDB() (*sql.DB, error) {
	user, pass, port, db_name := "postgres", "123", 5432, "test_db"
	log.Println("Первое подключение начинается")
	db, err := sql.Open("postgres", fmt.Sprintf("postgresql://%s:%s@localhost:%d/?sslmode=disable", user, pass, port))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}
	
	log.Println("Первое подключение выполнено")
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", db_name))
	if err != nil {
		return nil, fmt.Errorf("failed to create database: %v", err)
	}

	log.Println("База создана")

	db, err = sql.Open("postgres", fmt.Sprintf("postgresql://%s:%s@localhost:%d/%s?sslmode=disable", user, pass, port, db_name))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to test database: %v", err)
	}
	log.Println("Открываю sql")
	file, err := os.ReadFile("test_db_start.sql")
	if err != nil {
		return nil, err
	}
	log.Println("Прочитал sql")
	log.Println(string(file))
	_, err = db.Exec(string(file))
	if err != nil {
		return nil, err
	}

	return db, nil
}

func DropDB(db *sql.DB) error {
	db.Close()
	user, pass, port, db_name := "postgres", "123", 5432, "test_db"
	log.Println("Подключение для удаления начинается")
	db, err := sql.Open("postgres", fmt.Sprintf("postgresql://%s:%s@localhost:%d/?sslmode=disable", user, pass, port))
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}
	
	log.Println("Подключение выполнено")
	_, err = db.Exec(fmt.Sprintf("DROP DATABASE %s", db_name))
	if err != nil {
		return fmt.Errorf("failed to create database: %v", err)
	}

	log.Println("База удалена")
	db.Close()
	return nil
}
