package server

import (
	"encoding/json"
	"log"
	"net/http"
)

type SegmentName struct{
  Name string
}

// создание сегмента
func createSegmentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var req SegmentName
	err := decoder.Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	res, err := segmentRepo.Create(req.Name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error")
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}

}

// удаление сегмента
func deleteSegmentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var req SegmentName
	err := decoder.Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = segmentRepo.Delete(req.Name)
	if err != nil {
		log.Println("не нашел")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error! Not found!")
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Success")
	}
}
