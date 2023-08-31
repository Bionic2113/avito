package server

import (
	"encoding/json"
	"log"
	"net/http"
)

type UserSegmentsData struct {
	Id  int      `json:"id"`
	Add []string `json:"add,omitempty"`
	Del []string `json:"del,omitempty"`
}

// добавление и удаление сегментов у юзера
// тут же можно и добавлять в эксперименты
func updateSegmentsForUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("In handler")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var req UserSegmentsData
	err := decoder.Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	res, err := userSegmnetService.UpdateUserSegments(req.Id, req.Add, req.Del)
	if err != nil {
		log.Println("Error in updateSegmentsForUserHandler: ", err)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}
}

// получения активных сегментов пользователя
func activeUserSegmentsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var req Message
	err := decoder.Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	res, err := usRepo.FindAllById(req.Id, true)
	if err != nil {
		log.Println("Error in activeUserSegmentsHandler: ", err)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}

}

// получение истории с добавлением и удалением сегентов у пользователя
func historySegmentsForUserHandler(w http.ResponseWriter, r *http.Request) {
}

// добавление сегмента какому-то проценту пользователей
func addPersentUsersForSegment(w http.ResponseWriter, r *http.Request) {
}
