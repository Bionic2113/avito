package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type UserSegmentsData struct {
	Id        int      `json:"id,omitempty"`
	Add       []string `json:"add,omitempty"`
	Del       []string `json:"del,omitempty"`
	Durations []uint64 `json:"durations,omitempty"`
}

// добавление и удаление сегментов у юзера
// тут же можно и добавлять в эксперименты
func updateSegmentsForUserHandler(w http.ResponseWriter, r *http.Request) {
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

	res, err := userSegmnetService.UpdateUserSegments(req.Id, req.Add, req.Del, req.Durations)
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
	if r.Method != http.MethodPost {
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

type Date struct {
	Year  int `json:"year"`
	Month int `json:"month"`
}

// получение истории с добавлением и удалением сегентов у пользователя
func historySegmentsForUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var req Date
	err := decoder.Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	filename := userSegmnetService.CreateCSV(req.Year, req.Month)
	if filename != "Not found elements" {
		json.NewEncoder(w).Encode(fmt.Sprintf("%s/%s", r.URL.String(), filename))
	} else {
		json.NewEncoder(w).Encode("Not found elements")
	}
}

type SegmPercent struct {
	SegmentName string `json:"segment_name"`
	Percent     int    `json:"percent"`
}

// добавление сегмента какому-то проценту пользователей
func addPersentUsersForSegment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var req SegmPercent
	err := decoder.Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	res, err := userSegmnetService.AddPercent(req.SegmentName, req.Percent)
	if err != nil {
		log.Println("Error in updateSegmentsForUserHandler: ", err)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}

}
