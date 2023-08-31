package server_test

import (
	"testing"

	"github.com/Bionic2113/avito/cmd/server"
	"github.com/Bionic2113/avito/internal/repository"
	database_test "github.com/Bionic2113/avito/test/database"

	gomock "go.uber.org/mock/gomock"
)

type Message struct {
	Id int `json:"id"`
}

var (
	userMockDB    *database_test.MockUserRepository
	segmentMockDB *database_test.MockSegmentRepository
	usMockDB      *database_test.MockUserSegmentRepository

	userRepo        repository.UserRepository
	segmentRepo     repository.SegmentRepository
	userSegmentRepo repository.UserSegmentRepository
)

func TestMain(m *testing.M) {
	t := &testing.T{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userMockDB = database_test.NewMockUserRepository(ctrl)
	segmentMockDB = database_test.NewMockSegmentRepository(ctrl)
	usMockDB = database_test.NewMockUserSegmentRepository(ctrl)
	go server.StartServer(userMockDB, segmentMockDB, usMockDB)

	m.Run()
	// db, err := SetupDB()
	// if err != nil {
	// 	t.Fatal("Не создалась бд", err)
	// }
	//
	// userRepo = &repository.UserRepoDB{DB: db}
	// segmentRepo = &repository.SegmentRepositoryDB{DB: db}
	// userSegmentRepo = &repository.UserSegmentDB{DB: db}
	// server.Port = 8081
	// go server.StartServer(userRepo, segmentRepo, userSegmentRepo)
	//
	// go server.StartServer(userMockDB, segmentMockDB, usMockDB)
	// m.Run()
	//
	// if err = DropDB(db); err != nil {
	// 	t.Fatal("Не удалилась бд", err)
	// }
}

// func SetupDB() (*sql.DB, error) {
// 	user, pass, port, db_name := "postgres", "123", 5432, "test_db"
// 	log.Println("Первое подключение начинается")
// 	db, err := sql.Open("postgres", fmt.Sprintf("postgresql://%s:%s@localhost:%d/?sslmode=disable", user, pass, port))
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to connect to database: %v", err)
// 	}
//
// 	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXIST %s", db_name))
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to create database: %v", err)
// 	}
//
// 	log.Println("Первое подключение выполнено")
//
// 	db, err = sql.Open("postgres", fmt.Sprintf("postgresql://%s:%s@localhost:%d/%s?sslmode=disable", user, pass, port, db_name))
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to connect to test database: %v", err)
// 	}
// 	log.Println("Открываю sql")
// 	file, err := os.ReadFile("test_db_start.sql")
// 	if err != nil {
// 		return nil, err
// 	}
// 	log.Println("Прочитал sql")
// 	log.Println(string(file))
// 	_, err = db.Exec(string(file))
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return db, nil
// }
//
// func DropDB(db *sql.DB) error {
// 	db_name := "test_db"
// 	_, err := db.Exec(fmt.Sprintf("DROP DATABASE %s", db_name))
// 	return err
// }
