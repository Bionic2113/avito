package service

import "github.com/Bionic2113/avito/internal/models"

// сначала найди пользователя и чекни есть ли какие-то
// нужные записи для удаления, чтобы не получить ошибки,
// если нет таких данных
// также нужно проверить нет ли уже каких-то, что надо
// добавить
// можно даже отчет в конце намутить какие лишние были
// в списках на удаление и добавление
func UpdateUserSegments(id int, add []string, del []string) ([]models.UserSegment, error){
 return nil, nil
} 
