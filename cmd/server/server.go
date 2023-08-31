package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Bionic2113/avito/internal/repository"
	"github.com/Bionic2113/avito/pkg/service"
)

var (
  userRepo repository.UserRepository
  segmentRepo repository.SegmentRepository
  usRepo repository.UserSegmentRepository
  userSegmnetService service.Service
  Port int = 8080 //для тестирования нужно
)

// запуск сервера
func StartServer(u repository.UserRepository, s repository.SegmentRepository, us repository.UserSegmentRepository) {
	userRepo, segmentRepo, usRepo = u, s, us
	userSegmnetService = *userSegmnetService.New(u, us)
	http.HandleFunc("/user/findByID", findUserByIdHandler)
	http.HandleFunc("/user/findAll", findAllUsersHandler)
	http.HandleFunc("/segment/create", createSegmentHandler)
	http.HandleFunc("/segment/delete", deleteSegmentHandler)
	http.HandleFunc("/user_segment/update", updateSegmentsForUserHandler)
	http.HandleFunc("/user_segment/active", activeUserSegmentsHandler)
	http.HandleFunc("/user_segment/history", historySegmentsForUserHandler)
	http.HandleFunc("/user_segment/add_percent", addPersentUsersForSegment)
	go trackingForDeletion()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", Port), nil))

}

// проверка на удаление сегмента по истечению срока
func trackingForDeletion() {
   userSegmnetService.CleaningOld()
}

