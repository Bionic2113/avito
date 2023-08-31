package main

import (
	"database/sql"
	"log"

	"github.com/Bionic2113/avito/internal/repository"
	"github.com/Bionic2113/avito/cmd/server"
	_ "github.com/lib/pq"
  )

var (
	userRepo        repository.UserRepository
	segmentRepo     repository.SegmentRepository
	userSegmentRepo repository.UserSegmentRepository
)

func init(){
  db, err := sql.Open("postgres", "user=postgres password=123 dbname=avito_test sslmode=disable") 
  if err != nil{
	log.Fatal(err)
  }
  userRepo = &repository.UserRepoDB{DB:db}
  segmentRepo = &repository.SegmentRepositoryDB{DB:db}
  userSegmentRepo = &repository.UserSegmentDB{DB:db}
}

func main() {
	server.StartServer(userRepo, segmentRepo, userSegmentRepo)
	// u, err := userSegmentRepo.FindAllById(1, true)
	// if err != nil{
	//   log.Fatal(err)
	// }
	// for _,v := range u{
	//   log.Println(v)
	// }
}
