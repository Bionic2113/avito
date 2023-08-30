package server

import (
	"log"
	"net/http"

	"github.com/Bionic2113/avito/internal/repository"
)

var (
  userRepo repository.UserRepository
  segmentRepo repository.SegmentRepository
  usRepo repository.UserSegmentRepository
)

// запуск сервера
func StartServer(u repository.UserRepository, s repository.SegmentRepository, us repository.UserSegmentRepository) {
	userRepo, segmentRepo, usRepo = u, s, us
	http.HandleFunc("/user/findByID", findUserByIdHandler)
	http.HandleFunc("/user/findAll", findAllUsersHandler)
	http.HandleFunc("/segment/create", createSegmentHandler)
	http.HandleFunc("/segment/delete", deleteSegmentHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

// проверка на удаление сегмента по истечению срока
func checkOnDelete() {}
