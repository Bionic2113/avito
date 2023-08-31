package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Bionic2113/avito/cmd/server"
	"github.com/Bionic2113/avito/internal/repository"
	_ "github.com/lib/pq"
)

var (
	userRepo        repository.UserRepository
	segmentRepo     repository.SegmentRepository
	userSegmentRepo repository.UserSegmentRepository
)

func init() {
	user, pass, port, db_name := "postgres", "123", 5432, "avito_test"
	db, err := sql.Open(
		"postgres",
		fmt.Sprintf("postgresql://%s:%s@localhost:%d/%s?sslmode=disable", user, pass, port, db_name),
	)
	if err != nil {
		log.Fatalf("failed to connect to test database: %v", err)
	}

	userRepo = &repository.UserRepoDB{DB: db}
	segmentRepo = &repository.SegmentRepositoryDB{DB: db}
	userSegmentRepo = &repository.UserSegmentDB{DB: db}
}

func main() {
	server.StartServer(userRepo, segmentRepo, userSegmentRepo)
}
