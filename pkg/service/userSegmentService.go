package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/Bionic2113/avito/internal/models"
	"github.com/Bionic2113/avito/internal/repository"
)

// сначала найди пользователя и чекни есть ли какие-то
// нужные записи для удаления, чтобы не получить ошибки,
// если нет таких данных
// также нужно проверить нет ли уже каких-то, что надо
// добавить
// можно даже отчет в конце намутить какие лишние были
// в списках на удаление и добавление
type Service struct {
	userRepo        repository.UserRepository
	userSegmentRepo repository.UserSegmentRepository
}

type elem struct {
	name   string
	status bool
}

func (service *Service) New(ur repository.UserRepository, usr repository.UserSegmentRepository) *Service{
  return &Service{userRepo: ur, userSegmentRepo: usr}
}

// Вопросы:
// 1) Если сегмент удален, то надо ли удалять у пользователя сегмент?
// Под удален имею ввиду, что active = false
// Решил, что в этом нет необходимости, тк может даже удаление
// сегмента вообще не нужно, ведь в задании об этом не упоминалось.
func (service *Service) UpdateUserSegments(id int, add []string, del []string) ([]models.UserSegment, error) {
	result, deleted, err := service.checkData(id, add, del)
	if err != nil {
		return nil, err
	}
	create_chan := make(chan []models.UserSegment)
	go func() {
		create_chan <- service.create(id, add)
	}()
	result = service.delete(result, del)
	new_data := <-create_chan
	result = append(result, deleted...)
	result = append(result, new_data...)
	return result, nil
}

// для создания же по сути тебе надо только айди чела и название сегмента
// поэтому вроде проблем в данных нет
func (service *Service) create(id int, add []string) []models.UserSegment {
  create_array := make([]models.UserSegment,len(add))
	for i, v := range add {
		userSegment := &models.UserSegment{
			User_id:      uint64(id),
			Segment_name: v,
		}
		userSegment, err := service.userSegmentRepo.Create(userSegment)
				if err != nil{
					userSegment = &models.UserSegment{Segment_name: "Error in the creation method"}
				}
		create_array[i] = *userSegment
	}
	return create_array
}

// но тут тебе надо список тех, у кого надо бы изменить данные
// получается что сюда надо передать массив из service.userSegmentRepo.FindAllById(id, true)
// затем ты его тут отредачишь, потом вернешь и к нему добавишь массив из create
func (service *Service) delete(data []models.UserSegment, del []string) []models.UserSegment {
	for i, userSegment := range data {
		for _, str := range del {
			if userSegment.Segment_name == str {
				userSegment.Active = false
				userSegment.DeletionTime = uint64(time.Now().Unix())
				userSegment, err := service.userSegmentRepo.Update(&userSegment)
				if err != nil{
					userSegment = &models.UserSegment{Segment_name: "Error in the deletion method"}
				}
				data[i] = *userSegment
			} 
		}
	}
	return data
}

func (service *Service) checkData(id int, add []string, del []string) ([]models.UserSegment, []models.UserSegment, error) {
	if user, err := service.userRepo.FindById(id); err != nil || !user.Active {
		return nil, nil, fmt.Errorf("user is not found or have status deleted. user status is %v", user.Active)
	}

	active, err := service.userSegmentRepo.FindAllById(id, true)
	if err != nil {
		return nil, nil, errors.New("error when searching for active")
	}
	activeChan := make(chan elem)
	go func() {
		var e elem
		for _, s := range add {
			e = contains(s, active)
			if e.status {
				break
			}
		}
		activeChan <- e
	}()
	deleted, err := service.userSegmentRepo.FindAllById(id, false)
	if err != nil {
		return nil, nil, errors.New("error when searching for deleted")
	}
	for _, s := range del {
		if deleted_elem := contains(s, deleted); deleted_elem.status {
			return nil, nil, fmt.Errorf("user already have %s in deleted array", deleted_elem.name)
		}
	}
	active_elem := <-activeChan
	if active_elem.status {
		return nil, nil, fmt.Errorf("user already have %s in active array", active_elem.name)
	}
	return active, deleted, nil
}

func contains(str string, arr []models.UserSegment) elem {
	for _, s := range arr {
		if str == s.Segment_name {
			return elem{s.Segment_name, true}
		}
	}
	return elem{status: false}
}
