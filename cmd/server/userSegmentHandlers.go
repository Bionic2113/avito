package server

import "net/http"

// добавление и удаление сегментов у юзера
// тут же можно и добавлять в эксперименты
func updateSegmentsForUserHandler(w http.ResponseWriter, r *http.Request) {
}

// получения активных сегментов пользователя
func activeUserSegmentsHandler(w http.ResponseWriter, r *http.Request) {
}

// получение истории с добавлением и удалением сегентов у пользователя
func historySegmentsForUserHandler(w http.ResponseWriter, r *http.Request) {
}

// добавление сегмента какому-то проценту пользователей
func addPersentUsersForSegment(w http.ResponseWriter, r *http.Request) {
}
