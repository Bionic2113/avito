package server

import (
	"encoding/json"
	"log"
	"net/http"
)

type Message struct {
	Id int
}

// удаление пользователя
func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
}

// создание пользователя
func createUserHandler(w http.ResponseWriter, r *http.Request) {
}

// изменение данных пользователя
func updateUserHandler(w http.ResponseWriter, r *http.Request) {
}

// найти юзера по id
func findUserByIdHandler(w http.ResponseWriter, r *http.Request) {
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

	res, err := userRepo.FindById(req.Id)
	if err != nil {
		log.Println("не нашел")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Not found")
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}
}

// список всех юзеров
func findAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	res, err := userRepo.FindAll()
	if err != nil {
		log.Println("Что то не так с бд")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Not found")
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}
}
